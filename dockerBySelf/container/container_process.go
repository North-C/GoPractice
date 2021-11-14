package container

import (
	"os"
	"os/exec"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// 父进程，即当前进程执行的内容
// 第一种方法: /proc/self 即当前进程运行的环境, 在其中利用`exe init agrs`来执行fork 创建新的子进程， 设置CLONE的flag进行隔离
// 第二种： 使用匿名管道来实现父子通信
func NewParentProcess(tty bool) (*exec.Cmd, *os.File) {
	readPipe, writePipe, err := NewPipe()
	if err != nil {
		log.Errorf("New pipe error %v", err)
		return nil, nil
	}

	cmd := exec.Command("/proc/self/exe", "init") // 调用当前进程进行运行
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWIPC | syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWPID,
	}
	if tty { // 用户指定了`-ti`参数，则把当前进程的in和out导入到标准的流当中
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	// 传入管道文件读取端的句柄, ExtraFiles会外带这个文件句柄去创建子进程，它不包含标准的三个
	// 是除他们之外的第四个， /proc/self/fd可以看到
	cmd.ExtraFiles = []*os.File{readPipe}
	mntURL := "/root/mnt/" // Overlay下的挂载点应该是 /root/merged/
	rootURL := "/root/"
	NewWorkSpace(rootURL, mntURL)
	cmd.Dir = mntURL // 利用mnt来作为root目录，与宿主机隔离开
	return cmd, writePipe
}

// 新建管道，读写两端
func NewPipe() (*os.File, *os.File, error) {
	read, write, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	return read, write, nil
}

// 创建容器文件系统，分三部分： lower upper merged work
func NewWorkSpace(rootURL string, mntURL string) {
	CreateLowerDir(rootURL)
	CreateUpperDir(rootURL)
	// CreateMerged(rootURL)
	CreateWork(rootURL)

	CreateMountPoint(rootURL, mntURL)
}

// 删除创建出来的workspace
func DeleteWorkSpace(rootURL string, mntURL string) {

	DeleteMountPoint(rootURL, mntURL)
	DeleteWork(rootURL)
	DeleteUpper(rootURL)

}

// 将busybox.tar解压,busybox作为不变的lower
func CreateLowerDir(rootURL string) {
	busyboxURL := rootURL + "busybox/"
	busyboxTarURL := rootURL + "busybox.tar"
	exist, err := PathExists(busyboxURL)

	if err != nil {
		log.Infof("Fail to judge whether dir %s exists. %v", busyboxURL, err)
	}
	// 不存在目录
	if exist == false {
		if err := os.Mkdir(busyboxURL, 0777); err != nil {
			log.Errorf("Mkdir %s error %v", busyboxURL, err)
		}

		// 执行解压命令,输出stdout和stderr内容
		if _, err := exec.Command("tar", "-xvf", busyboxTarURL, "-C", busyboxURL).CombinedOutput(); err != nil {
			log.Errorf("Untar dir %s error %v", busyboxURL, err)
		}

	}
}

// upper 是读写层
func CreateUpperDir(rootURL string) {
	writeURL := rootURL + "upper/"
	exist, _ := PathExists(writeURL)

	if exist == false {
		if err := os.Mkdir(writeURL, 0777); err != nil {
			log.Errorf("Mkdir dir %s error. %v", writeURL, err)
		}
	}
}

// merged 是mountpoint,也是lower和upper的联合视图
func CreateMerged(rootURL string) {
	writeURL := rootURL + "merged/"
	if err := os.Mkdir(writeURL, 0777); err != nil {
		log.Errorf("Mkdir dir %s error. %v", writeURL, err)
	}
}

func CreateWork(rootURL string) {
	writeURL := rootURL + "work/"
	exist, _ := PathExists(writeURL)

	if exist == false {
		if err := os.Mkdir(writeURL, 0777); err != nil {
			log.Errorf("Mkdir dir %s error. %v", writeURL, err)
		}
	}
}

func CreateMountPoint(rootURL string, mntURL string) {
	exist, _ := PathExists(mntURL)

	if exist == false {
		if err := os.Mkdir(mntURL, 0777); err != nil {
			log.Errorf("Mkdir dir %s error. %v", mntURL, err)
		}
	}
	// 挂载为overlay文件系统，挂载到mnt下面，将其作为merged
	dirs := "lowerdir=" + rootURL + "busybox," + "upperdir=" + rootURL + "upper," + "workdir=" + rootURL + "work"

	cmd := exec.Command("mount", "-t", "overlay", "overlay", "-o", dirs, mntURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Errorf("%v", err)
	}
}

//删除镜像
func DeleteMountPoint(rootURL string, mntURL string) {
	cmd := exec.Command("umount", mntURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Errorf("%v", err)
	}
	if err := os.RemoveAll(mntURL); err != nil {
		log.Errorf("Remove dir %s error %v", mntURL, err)
	}
}

func DeleteUpper(rootURL string) {
	writeURL := rootURL + "upper/"
	if err := os.RemoveAll(writeURL); err != nil {
		log.Errorf("Remove dir %s error %v", writeURL, err)
	}
}

func DeleteWork(rootURL string) {
	writeURL := rootURL + "work/"
	if err := os.RemoveAll(writeURL); err != nil {
		log.Errorf("Remove dir %s error %v", writeURL, err)
	}
}

// 判断是否存在
func PathExists(url string) (bool, error) {
	_, err := os.Stat(url)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
