package conf

import (
	"bytes"
	_ "embed"
	"github.com/TskFok/GinApi/app/global"
	"github.com/spf13/viper"
	"os"
	"strings"
)

//go:embed conf.yaml
var conf []byte

func InitConfig() {
	global.Env = strings.Replace(os.Args[1], "--env=", "", 1)
	viper.SetConfigType("yaml")

	err := viper.ReadConfig(bytes.NewReader(conf))

	if nil != err {
		panic(err)
	}

	dir, err := os.UserHomeDir()

	if nil != err {
		panic(err.Error())
	}

	runtimePath := viper.Get("logger.file_path").(string)
	global.LoggerFilePath = runtimePath

	_, fErr := os.Stat(runtimePath)

	//当前是否存在runtime目录,不存在会在系统用户目录中生成runtime目录
	if nil != fErr {
		afterAppend := append([]byte(dir), "/"...)
		afterAppend = append(afterAppend, runtimePath...)
		runtimePath = string(afterAppend)

		_, fiErr := os.Stat(runtimePath)

		if nil != fiErr {
			mErr := os.Mkdir(runtimePath, 0755)

			if nil != mErr {
				panic(mErr.Error())
			}
		}

		global.LoggerFilePath = runtimePath
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

	global.KafkaHost = viper.Get("kafka.host").(string)
	global.KafkaTopic = viper.Get("kafka.topic").(string)
	global.ElasticHost = viper.Get("elastic.host").(string)
	global.ElasticUsername = viper.Get("elastic.username").(string)
	global.ElasticPassword = viper.Get("elastic.password").(string)
}
