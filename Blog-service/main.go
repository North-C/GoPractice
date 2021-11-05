package main

import (
	"log"
	"net/http"
	"time"

	"blog-service/global"
	"blog-service/internal/model"
	"blog-service/internal/routers"
	"blog-service/pkg/setting"

	"github.com/gin-gonic/gin"
)

// 初始化配置的读取
/*	init执行初始化操作，在main之前自动执行，执行顺序是：
	全局变量初始化 ---> init方法 ---> main方法
*/
func init(){
	err := setupSetting()
	if err != nil{
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil{
		log.Fatalf("init.setupDBEngine err: %v", err)
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

func main() {
	// 将应用配置和gin的运行模式进行设置
	gin.SetMode(global.ServerSetting.RunMode)
	
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}



