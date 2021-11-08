package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

// 查看UTS的效果，对hostname进行了隔离
// readlink /proc/[pid]/ns/uts 返回对应的uts namespace
func main(){
	// Command返回一个Cmd结构体，但只会设置Path和Args
	// fork一个新的进程
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// 添加一个clone的flag
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// cmd执行了Run方法之后就无法再次使用
	if err := cmd.Run(); err != nil{
		log.Fatal(err)
	}
}