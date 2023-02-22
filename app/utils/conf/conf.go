package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func GetConf(key string) any {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("app/utils/conf/app.yaml")
	err := viper.ReadInConfig()

	if nil != err {
		fmt.Println(err)
	}

	return viper.Get(key)
}
