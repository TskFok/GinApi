package elasticsearch

import (
	"context"
	"fmt"
	"github.com/TskFok/GinApi/app/global"
	"github.com/TskFok/GinApi/utils/logger"
	"github.com/olivere/elastic/v7"
)

// Add 写入
func Add(log interface{}, index string) {
	ctx := context.Background()

	res, err := global.ElasticsearchClient.Index().Index(index).BodyJson(log).Do(ctx)

	if nil != err {
		logger.Error(err.Error())
	}

	fmt.Println(res)
}

// Batch 批量写入
func Batch(logs []interface{}) {
	ctx := context.TODO()

	bulk := global.ElasticsearchClient.Bulk().Index("log")

	for _, v := range logs {
		doc := elastic.NewBulkIndexRequest().Doc(v)

		bulk.Add(doc)
	}

	res, err := bulk.Do(ctx)

	if nil != err {
		logger.Error(err.Error())
	}

	fmt.Println(res)
}
