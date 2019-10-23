package server

import (
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var kafkaWg sync.WaitGroup

func (t TailInfo) SendToKafka() {
	if t.ExitSign {
		wg.Done()
		return
	}

	for {
		message, ok := <-MessageChan
		if !ok {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		msg := &sarama.ProducerMessage{}
		msg.Topic = message.Topic
		msg.Value = sarama.StringEncoder(message.LineLog)
		_, _, err = KafkaClient.SendMessage(msg)
		if err != nil {
			err = fmt.Errorf("send message failed,", err)
			logs.Warn(err)
		}
		fmt.Println("send message to kafka", message.Topic, message.LineLog)
	}

}
