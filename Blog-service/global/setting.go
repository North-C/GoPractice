package global

import "blog-service/pkg/setting"


// 全局配置，将其和应用程序关联起来
var(
	ServerSetting  		*setting.ServerSettings
	AppSettings			*setting.AppSettings
	DatabaseSettings	*setting.DatabaseSettings
)

