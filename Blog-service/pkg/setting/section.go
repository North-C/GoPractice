package setting

import (
	"time"
)
// 为配置文件声明对应的结构体
type ServerSettings struct{
	RunMode  string
	HttpPort string
	ReadTimeout 	time.Duration
	WriteTimeout 	time.Duration
}

type AppSettings struct{
	DefaultPageSize int
	MaxPageSize 	int
	LogSavePath 	string
	LogFileName 	string
  	LogFileExt 	string
}

type DatabaseSettings struct{
	DBType			string
  	Username		string
  	Password		string
  	Host			string
  	DBName 			string
  	TablePrefix 	string
  	Charset			string
  	ParseTime		bool
  	MaxIdleConns 	int
  	MaxOpenConns	int
}

// 读取区段配置的配置方法
func (s *Setting) ReadSection(k string, v interface{}) error{
	err := s.vp.UnmarshalKey(k, v)
	if err != nil{
		return err
	}

	return nil
}