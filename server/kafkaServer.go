package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/TskFok/GinApi/app/global"
	"github.com/TskFok/GinApi/bootstrap"
	"github.com/TskFok/GinApi/utils/elasticsearch"
	"sync"
	"time"
)

/*
*
kafka服务端
*/
func main() {
	bootstrap.Init()

	var wg sync.WaitGroup

	//配置
	host := global.KafkaHost
	topic := global.KafkaTopic
	consumer, err := sarama.NewConsumer([]string{host}, nil)
	if err != nil {
		panic(fmt.Sprintf("Failed to start consumer: %s", err))
	}
	partitionList, err := consumer.Partitions("task-status-data") // 通过topic获取到所有的分区
	if err != nil {
		panic("Failed to get the list of partition: " + err.Error())
	}

	for partition := range partitionList { // 遍历所有的分区
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest) // 针对每个分区创建一个分区消费者
		if err != nil {
			panic(fmt.Sprintf("Failed to start consumer for partition %d: %s\n", partition, err))
		}
		wg.Add(1)
		go func(sarama.PartitionConsumer) { // 为每个分区开一个go协程取值
			for msg := range pc.Messages() { // 阻塞直到有值发送过来，然后再继续等待
				switch string(msg.Key) {
				case "log":
					res := make(map[string]interface{})
					json.Unmarshal([]byte(msg.Value), &res)

					request := res["request"]
					response := res["response"]
					header := res["header"]
					query := res["query"]

					ret, jsonErr := json.Marshal(request)

					if nil != jsonErr {
						fmt.Println(jsonErr.Error())
					}

					code := response.(map[string]interface{})["code"]
					rep, jsonErr := json.Marshal(response)

					if nil != jsonErr {
						fmt.Println(jsonErr.Error())
					}

					hea, jsonErr := json.Marshal(header)

					if nil != jsonErr {
						fmt.Println(jsonErr.Error())
					}

					re := ApiLogStruct{
						Code:     int(code.(float64)),
						Request:  string(ret),
						Response: string(rep),
						Header:   string(hea),
						Query:    query.(string),
						Created:  time.Now().Format("2006-01-02 15:04:05"),
					}

					elasticsearch.Add(re, "api_log")
				}

			}
			defer pc.AsyncClose()
			wg.Done()
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}

type ApiLogStruct struct {
	Created  string `json:"created,omitempty"`
	Request  string `json:"request,omitempty"`
	Response string `json:"response,omitempty"`
	Header   string `json:"header,omitempty"`
	Query    string `json:"query,omitempty"`
	Code     int    `json:"code,omitempty"`
}
