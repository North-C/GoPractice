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
	cmd.Dir = "/root/busybox" // 利用busybox来作为root目录
	return cmd, writePipe
}

func NewPipe() (*os.File, *os.File, error) {
	read, write, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	return read, write, nil
}
