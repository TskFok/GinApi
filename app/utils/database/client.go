package database

import (
	"github.com/TskFok/GinApi/app/utils/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open(mysql.Open(conf.MysqlDsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.MysqlPrefix, //前缀
			SingularTable: true,             //复数表名
		},
	})

	if nil != err {
		panic("fail to open mysql connect ")
	}
}
