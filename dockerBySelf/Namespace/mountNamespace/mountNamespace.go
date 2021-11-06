package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

// chroot 将一个子目录变成根节点，Mount Namespace更强大
// flag 是 CLONE_NEWNS

// 运行程序后，在内部进行mount： `mount -t proc proc /proc`
// mount之后就可以查看进程 `ps -ef`
func main(){
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
		syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID |
		syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}