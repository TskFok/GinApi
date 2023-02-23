package conf

import (
	"fmt"
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
)

func InitConf() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("app/utils/conf/app.yaml")
	err := viper.ReadInConfig()

	if nil != err {
		fmt.Println(err)
	}

	RedisHost = viper.Get("redis.host").(string)
	RedisPassword = viper.Get("redis.password").(string)
	MysqlDsn = viper.Get("mysql.dsn").(string)
	MysqlPrefix = viper.Get("mysql.prefix").(string)
	AppReadTimeOut = viper.Get("app.read_time_out").(int)
	AppWriteTimeOut = viper.Get("app.write_time_out").(int)
	AppRunMode = viper.Get("app.run_mode").(string)
	AppHttpPort = viper.Get("app.http_port").(int)

}
