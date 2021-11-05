package model

import (
	"blog-service/global"
	"blog-service/pkg/setting"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Model struct{
	ID uint32 `gorm:"primary_key" json:"id"`
	CreatedBy string `json: "created_by"`
	CreatedOn uint32 `json: "created_on"`
	ModifiedBy string `json: "modified_by"`
	ModifiedOn uint32 `json: "modified_on"`
	DeletedOn uint32 `json: "deleted_on"`
	IsDel uint8 `json: "is_del"`
}


func NewDBEngine(databaseSettings *setting.DatabaseSettings) (*gorm.DB, error){
	// 初始化
	db, err := gorm.Open(databaseSettings.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSettings.Username,
		databaseSettings.Password,
		databaseSettings.Host,
		databaseSettings.DBName,
		databaseSettings.Charset,
		databaseSettings.ParseTime,
	))
	if err != nil{
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug"{
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSettings.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSettings.MaxOpenConns)

	return db, nil
}


