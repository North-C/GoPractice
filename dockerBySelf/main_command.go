package main

import (
	"dockerBySelf/cgroups/subsystems"
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
		cli.BoolFlag{		// 开启虚拟终端
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.StringFlag{
			Name: "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name: "cpushare",
			Usage: "cpu limit",
		},
		cli.StringFlag{			
			Name: "cpuset",
			Usage: "cpuset limit",
		},
		cli.StringFlag{		// 挂载容器的功能
			Name: "v",
			Usage: "volume",
		},
		cli.BoolFlag{		// 后台运行容器
			Name: "d",
			Usage: "daemon",
		},
		// 提供run后面的-name指定容器名字参数,便于查看信息
		cli.StringFlag{
			Name: "name",
			Usage: "cpuset name",
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
		detach := context.Bool("d")
		if tty && detach {		// 前台和后台不能同时提供
			return fmt.Errorf("ti and daemon can't both provided")
		}

		resConf := &subsystems.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuSet: context.String("cpuset"),
			CpuShare: context.String("cpushare"),
		}
		log.Infof("creatTty %v", tty)

		containerName := context.String("name")

		Run(tty, cmdArray, resConf, containerName)

		/* 挂载存储卷
		volume := context.String("v")
		Run(tty, cmdArray, volume) */
		return nil
	},
}

var commitCommand = cli.Command{
	Name: "commit",
	Usage: "commit a container into image",
	Action: func(context *cli.Context) error{
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		imageName := context.Args().Get(0)		
		commitContainer(imageName)
		return nil
	},
}

// 这是内部执行的，不允许外部调用
var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container.",
	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		err := container.RunContainerInitProcess()
		return err
	},
}
