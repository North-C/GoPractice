package main

import (
	"os"

	"github.com/urfave/cli"
	log "github.com/sirupsen/logrus"
)

const usage = `mydocker is a simple container runtime implementation.
				Enjoy it, just for fun.`


func main(){
	app := cli.NewApp()		// 创建一个新的APP
	app.Name = "mydocker"
	app.Usage = usage

	app.Commands = []cli.Command{ initCommand, runCommand}
// 初始化logrus的日志配置
	app.Before = func(context *cli.Context) error{
		// Log as JSON instead of the default ASCII formatter
		log.SetFormatter(&log.JSONFormatter{})

		log.SetOutput(os.Stdout)
		return nil
	}
	if err := app.Run(os.Args); err!= nil{
		log.Fatal(err)
	}
}

