package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/TskFok/GinApi/app/utils/conf"
	"github.com/TskFok/GinApi/app/utils/logger"
)

// Send 发送消息到kafka
func Send(msg string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //写到随机分区中，我们默认设置32个分区
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	producerMsg := &sarama.ProducerMessage{}
	producerMsg.Topic = conf.KafkaTopic
	producerMsg.Value = sarama.StringEncoder(msg)

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{conf.KafkaHost}, config)
	if err != nil {
		fmt.Println("Producer closed, err:", err)
		return
	}
	defer client.Close()

	// 发送消息
	pid, offset, err := client.SendMessage(producerMsg)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info(fmt.Sprintf("pid:%v offset:%v\n", pid, offset))
}
