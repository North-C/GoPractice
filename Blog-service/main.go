package main

import (
	"log"
	"net/http"
	"time"

	"blog-service/global"
	"blog-service/internal/model"
	"blog-service/internal/routers"
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/natefinch/lumberjack.v2"				// 它的核心功能是将日志写入滚动文件中
)

// 初始化配置的读取
/*	init执行初始化操作，在main之前自动执行，执行顺序是：
	全局变量初始化 ---> init方法 ---> main方法
*/
func init(){
	// 配置初始化
	err := setupSetting()
	if err != nil{
		log.Fatalf("init.setupSetting err: %v", err)
	}
	// 数据库初始化
	err = setupDBEngine()
	if err != nil{
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	// 日志初始化
	err = setupLogger()
	if err!= nil{
		log.Fatalf("init.setupLogger err: %v", err)
	}
}


// 进行配置，将配置文件的内容映射到应用程序的配置结构当中
func setupSetting() error{
	setting, err := setting.NewSetting()
	if err != nil{
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil{
		return err
	}

	err = setting.ReadSection("App", &global.AppSettings)
	if err != nil{
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSettings)
	if err != nil{
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error{
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSettings)

	if err != nil{
		return err
	}
	return nil
}

func setupLogger()error{
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: global.AppSettings.LogSavePath + "/" + global.AppSettings.LogFileName +
		  			global.AppSettings.LogSavePath,	 	// 文件路径名
		MaxSize: 600,		// 最大占用空间600MB
		MaxAge: 10,			// 生存周期10天
		LocalTime: true, 		// 时间格式为本地时间
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func main() {
	// 将应用配置和gin的运行模式进行设置
	gin.SetMode(global.ServerSetting.RunMode)
	
	//global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycyj", "blog-service")
	
	router := routers.NewRouter()
	// 接入到服务器
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// 开始监听
	for{
		s.ListenAndServe()
	}
	
}



