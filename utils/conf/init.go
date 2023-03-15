package conf

import (
	"github.com/TskFok/GinApi/app/global"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func InitConfig() {
	global.Env = strings.Replace(os.Args[1], "--env=", "", 1)

	viper.SetConfigType("yaml")
	viper.SetConfigFile("config/" + global.Env + ".yaml")
	err := viper.ReadInConfig()

	if nil != err {
		panic(err)
	}

	global.RedisHost = viper.Get("redis.host").(string)
	global.RedisPassword = viper.Get("redis.password").(string)
	global.MysqlDsn = viper.Get("mysql.dsn").(string)
	global.MysqlPrefix = viper.Get("mysql.prefix").(string)
	global.AppReadTimeOut = viper.Get("app.read_time_out").(int)
	global.AppWriteTimeOut = viper.Get("app.write_time_out").(int)
	global.AppHttpPort = viper.Get("app.http_port").(int)
	global.JwtSecret = viper.Get("jwt.secret").(string)
	global.JwtExpire = viper.Get("jwt.expire").(int)
	global.LoggerFilePath = viper.Get("logger.file_path").(string)
	global.KafkaHost = viper.Get("kafka.host").(string)
	global.KafkaTopic = viper.Get("kafka.topic").(string)
	global.ElasticHost = viper.Get("elastic.host").(string)
	global.ElasticUsername = viper.Get("elastic.username").(string)
	global.ElasticPassword = viper.Get("elastic.password").(string)
}
