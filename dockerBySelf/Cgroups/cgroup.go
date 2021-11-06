package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

/*
	cgroup 采用层级进行管理，从上至下： Hierarchy(树状结构，方便继承) --> subsystem --> control group -->task
		分别可以管理CPU,Memory,存储，Network等，限制进程占用并实时地监控进程和统计信息。
		`cat /proc/cgroups`查看kernel支持的subsystem
	相互联系：1. 一个subsystem只能附加到一个hierarchy
			2. 一个hierarchy 可以附加多个 subsystem
			3. 一个进程可以作为多个cgroup的成员，但是cgroup必须在不同的hierarchy当中
			4. 在fork时，可以选择将子进程的cgroup进行移动
*/

// 挂载memory subsystem的 hierarchy
const cgroupmemoryHierarchyMount = "/sys/fs/cgroup/"

func main(){
	if os.Args[0] == "/proc/slf/exe"{
		// 容器进程
		fmt.Printf("")
	}

	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err:= cmd.Start(); err != nil{
		fmt.Println("Error when starting", err)
		os.Exit(1)
	}else{
		fmt.Printf("%v", cmd.Process.Pid)

		os.Mkdir(path.Join(cgroupmemoryHierarchyMount, "testmemorylimit"), 0755)

		ioutil.WriteFile(path.Join(cgroupmemoryHierarchyMount, "testmemorylimit", "cgroup.procs"),
			[]byte(strconv.Itoa(cmd.Process.Pid)), 0644)
		
		ioutil.WriteFile(path.Join(cgroupmemoryHierarchyMount, "testmemorylimit", 
			"memory.limit_in_bytes"), []byte("100m"), 0644)

	}
	cmd.Process.Wait()
}
