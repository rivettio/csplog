package initall

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/Shopify/sarama"
)

func InitKafka() (client sarama.SyncProducer, err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(LogConfAll.KafkaConf.KafkaAddr, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	logs.Error("init kafka success")
	return
}

func InitKafkaConsumer() (consumer sarama.Consumer, err error) {
	consumer, err = sarama.NewConsumer(LogConfAll.KafkaConf.KafkaAddr, nil)
	if err != nil {
		err = fmt.Errorf("Failed to start consumer: %s", err)
		return
	}
	logs.Error("init kafka Consumer success")
	return
}