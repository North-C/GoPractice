package container

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// 在容器内部执行，容器内执行的第一个进程
// 使用匿名管道通信之后，需要进行修改
func RunContainerInitProcess(command string, args []string) error{
	cmdArray := readUserCommand()		// 获取用户命令
	if cmdArray == nil || len(cmdArray) == 0{
		return fmt.Errorf("Run container get user command error, cmdAraay is nil")
	}

	// systemd 加入linux之后， mount namespace就会变成 shared by default，所以需要声明新的mount namespace 独立
	syscall.Mount("", "/", "", syscall.MS_PRIVATE | syscall.MS_REC,"")
	
	// mount 系统调用, NOEXEC不允许其他程序运行，NOSUID不运行set-uid或者set-groupID，NODEV 默认设置的参数
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	// 挂载 proc文件系统到容器当中，之后可以查看系统进程的资源情况
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")

	// 去PATH中寻找命令的绝对路径
	path, err := exec.LookPath(cmdArray[0])
	if err!= nil{
		log.Errorf("exec loop path error %v", err)
		return err
	}

	log.Infof("Find path %s", path)	
	// 在子进程当中执行
	if err:= syscall.Exec(path, cmdArray[0:], os.Environ()); err!= nil{
		log.Errorf(err.Error())
	}
	/* 旧版本
	argv := []string{command}
	// 调用系统的execve函数，执行command程序，并覆盖当前进程的镜像，数据和堆栈等信息
	// 用它替换掉 一开始的 init进程
	 
	if err := syscall.Exec(command, argv, os.Environ()); err != nil{
		logrus.Errorf(err.Error())
	}
	*/
	return nil
}

func readUserCommand() []string{
	// index 为3的文件描述符，就是父进程传递过来的管道的一端
	pipe := os.NewFile(uintptr(3), "pipe")
	msg, err := ioutil.ReadAll(pipe)		//等在这里等待
	if err != nil{
		log.Errorf("init read pipe error %v", err)
	}
	msgStr := string(msg)
	return strings.Split(msgStr, " ")

}
/*
func setUpMount(){}

func pivotRoot(root string) error{
	
}
*/