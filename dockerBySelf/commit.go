package main

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func commitContainer(imageName string) {
	mntURL := "/root/mnt"
	imageTar := "/root/" + imageName + ".tar"
	fmt.Printf("%s", imageTar)
	// CombinedOutput 输出命令绑定的 output和error
	if _, err := exec.Command("tar", "-czf", imageTar, "-C", mntURL, ".").CombinedOutput(); err != nil {
		log.Errorf("Tar folder %s error %v", mntURL, err)
	}
}
