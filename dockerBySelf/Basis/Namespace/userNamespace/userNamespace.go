package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

// 隔离用户的用户组ID, 常用的是 在非root用户创建一个User namespace, 在里面映射成 root用户
// `id` 命令查看id 号
func main() {

	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
		/*
			以下两种情况会导致UidMapping和GidMapping设置不同的相关数值：
			1. HostID非本进程所有（与Getuid()和Getgid()不等）
			2. Size大于1 （则肯定包含非当前进程的UID和GID）
			需要Host使用Root权限才能执行该权限
		*/
		UidMappings: []syscall.SysProcIDMap{
			// 设置成 root 权限
			{
				ContainerID: 0,
				HostID:      0,
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      0,
				Size:        1,
			},
		},
		// GidMappingsEnableSetgroups: true,
		// Credential:                 &syscall.Credential{Uid: uint32(1), Gid: uint32(1)},
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(-1)
}
