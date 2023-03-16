package bootstrap

import (
	"github.com/TskFok/GinApi/app/global"
	"github.com/TskFok/GinApi/utils/cache"
	"github.com/TskFok/GinApi/utils/conf"
	"github.com/TskFok/GinApi/utils/database"
	"github.com/TskFok/GinApi/utils/elasticsearch"
	"github.com/TskFok/GinApi/utils/kafka"
	"github.com/TskFok/GinApi/utils/logger"
)

func Init() {
	// 引入配置
	conf.InitConfig()
	// logger
	global.LoggerClient = logger.InitLogger()
	// redis
	global.RedisClient = cache.InitRedis()
	// mysql
	global.MysqlClient = database.InitMysql()

	//debug环境下不需要开启elasticsearch和kafka,使用文件记录日志,
	//正式环境需要开启elasticsearch和kafka服务,使用elasticsearch记录日志
	if global.Env != "debug" {
		// elasticsearch
		global.ElasticsearchClient = elasticsearch.InitElasticsearch()
		// kafka
		global.KafkaClient = kafka.InitKafka()
	}

}
