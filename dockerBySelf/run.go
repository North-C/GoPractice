package main

import (
	// "dockerBySelf/cgroups"
	// "dockerBySelf/cgroups/subsystems"
	"dockerBySelf/cgroups"
	"dockerBySelf/cgroups/subsystems"
	"dockerBySelf/container"
	"os"

	"strings"

	log "github.com/sirupsen/logrus"
)

/* Tag 3.1
// 运行父进程
func Run(tty bool, command string) {

	/* Tag 3.1
	// 运行 父进程
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait() // 等待正常结束
	os.Exit(-1)
	parent, writePipe := container.NewParentProcess(tty)
}
*/

/* Tag 4.3
func Run(tty bool, comArray []string, volume string) {

	parent, writePipe := container.NewParentProcess(tty, volume)
	if parent == nil {
		log.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	//cgroupManager := cgroups.NewCgroupManager("donkey-cgroup")
	//defer cgroupManager.Destroy()
	//cgroupManager.Set(res)
	//cgroupManager.Apply(parent.Process.Pid)

	// 发送用户命令
	sendInitCommand(comArray, writePipe)
	parent.Wait()
	mntURL := "/root/mnt"
	rootURL := "/root"
	container.DeleteWorkSpace(rootURL, mntURL, volume)
	os.Exit(0)
}
*/
// Tag chapter 5
func Run(tty bool, cmdArray []string, resConf *subsystems.ResourceConfig, containerName string) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		log.Errorf("new parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		log.Error(err)
	}

	// 加入资源管理器
	cgroupsManager := cgroups.NewCgroupManager("mydocker-group")
	defer cgroupsManager.Destroy()
	cgroupsManager.Set(resConf)
	cgroupsManager.Apply(parent.Process.Pid)

	sendInitCommand(cmdArray, writePipe)
	if tty { // 等待子进程结束，如果没有tty则不用等
		parent.Wait()
	}
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
