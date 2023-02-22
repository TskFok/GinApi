package database

import (
	"fmt"
	"github.com/TskFok/GinApi/app/utils/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetClient() *gorm.DB {
	dsn := conf.GetConf("mysql.dsn")
	prefix := conf.GetConf("mysql.prefix")
	db, err := gorm.Open(mysql.Open(dsn.(string)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix.(string), //前缀
			SingularTable: true,            //复数表名
		},
	})

	if nil != err {
		fmt.Println(err)
		panic("fail to open mysql connect ")
	}

	return db
}
