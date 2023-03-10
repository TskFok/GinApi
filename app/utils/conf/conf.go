package conf

import (
	"github.com/spf13/viper"
)

var (
	RedisHost       string
	RedisPassword   string
	MysqlDsn        string
	MysqlPrefix     string
	AppReadTimeOut  int
	AppWriteTimeOut int
	AppRunMode      string
	AppHttpPort     int
	JwtSecret       string
	JwtExpire       int
	LoggerFilePath  string
	KafkaHost       string
	KafkaTopic      string
)

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("app/utils/conf/app.yaml")
	err := viper.ReadInConfig()

	if nil != err {
		panic(err)
	}

	RedisHost = viper.Get("redis.host").(string)
	RedisPassword = viper.Get("redis.password").(string)
	MysqlDsn = viper.Get("mysql.dsn").(string)
	MysqlPrefix = viper.Get("mysql.prefix").(string)
	AppReadTimeOut = viper.Get("app.read_time_out").(int)
	AppWriteTimeOut = viper.Get("app.write_time_out").(int)
	AppRunMode = viper.Get("app.run_mode").(string)
	AppHttpPort = viper.Get("app.http_port").(int)
	JwtSecret = viper.Get("jwt.secret").(string)
	JwtExpire = viper.Get("jwt.expire").(int)
	LoggerFilePath = viper.Get("logger.file_path").(string)
	KafkaHost = viper.Get("kafka.host").(string)
	KafkaTopic = viper.Get("kafka.topic").(string)

}
