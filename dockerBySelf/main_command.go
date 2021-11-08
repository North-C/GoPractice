package main

import (
	"dockerBySelf/container"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name: "run",
	Usage: `Create a container with namespace and cgroups limit
			mydocker run -ti [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	// run 执行的函数
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray = append(cmdArray, arg)
		}

		// cmd := context.Args().Get(0) // 对应的是`run -ti /bin/sh`中最后的`/bin/sh`

		tty := context.Bool("ti")
		Run(tty, cmdArray)
		return nil
	},
}

// 这是内部执行的，不允许外部调用
var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container.",
	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		cmd := context.Args().Get(0)
		log.Infof("init command %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}
