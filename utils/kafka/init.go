package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/TskFok/GinApi/app/global"
)

func InitKafka() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //写到随机分区中，我们默认设置32个分区
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{global.KafkaHost}, config)
	if err != nil {
		panic("Producer closed, err:" + err.Error())

		return nil
	}

	return client
}
