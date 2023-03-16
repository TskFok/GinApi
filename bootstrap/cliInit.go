package bootstrap

import (
	"github.com/TskFok/GinApi/app/global"
	"github.com/TskFok/GinApi/utils/cache"
	"github.com/TskFok/GinApi/utils/conf"
	"github.com/TskFok/GinApi/utils/database"
	"github.com/TskFok/GinApi/utils/elasticsearch"
	"github.com/TskFok/GinApi/utils/logger"
)

func CliInit() {
	// 引入配置
	conf.InitConfig()
	// logger
	global.LoggerClient = logger.InitLogger()
	// redis
	global.RedisClient = cache.InitRedis()
	// mysql
	global.MysqlClient = database.InitMysql()

	if global.Env != "debug" {
		// elasticsearch
		global.ElasticsearchClient = elasticsearch.InitElasticsearch()
	}

}
