package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/TskFok/GinApi/app/global"
	"github.com/TskFok/GinApi/utils/logger"
)

// Send 发送消息到kafka
func Send(msg string, key string) {
	// 构造一个消息
	producerMsg := &sarama.ProducerMessage{}
	producerMsg.Topic = global.KafkaTopic
	producerMsg.Value = sarama.StringEncoder(msg)
	producerMsg.Key = sarama.StringEncoder(key)

	// 发送消息
	pid, offset, err := global.KafkaClient.SendMessage(producerMsg)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info(fmt.Sprintf("pid:%v offset:%v\n", pid, offset))
}
