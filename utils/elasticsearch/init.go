package elasticsearch

import (
	"fmt"
	"github.com/TskFok/GinApi/app/global"
	"github.com/olivere/elastic/v7"
)

func InitElasticsearch() *elastic.Client {
	// 创建client
	client, err := elastic.NewClient(
		elastic.SetURL(global.ElasticHost),
		elastic.SetBasicAuth(global.ElasticUsername, global.ElasticPassword))
	if err != nil {
		panic(fmt.Sprintf("连接失败: %v\n", err))
	}
	isRunning := client.IsRunning()

	if !isRunning {
		panic("es连接失败")
	}

	return client
}
