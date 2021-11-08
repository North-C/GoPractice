package main

import (
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

func Run(tty bool, comArray []string){
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil{
		log.Errorf("New parent process error")
		return 
	}
	if err := parent.Start(); err!= nil{
		log.Error(err)
	}
	// 发送用户命令
	sendInitCommand(comArray, writePipe)
	parent.Wait()
	os.Exit(0)	
}

func sendInitCommand(comArray []string, writePipe *os.File){
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
