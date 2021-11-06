package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

// 查看结果： 
// ipcs -q 查看消息队列  ipcmk -Q 创建新的消息队列
// 首先在宿主机shell当中新建message queue, 之后在另一个shell运行程序，继续查看消息队列，会发现为空
func main(){

	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | 
		syscall.CLONE_NEWIPC,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil{
		log.Fatal(err)
	}
}