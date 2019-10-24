package server

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/Shopify/sarama"
)

func ReadKafka(c ConsumerInfo) {

	partitionList, err := c.Consumer.Partitions(c.Topic)
	if err != nil {
		err = fmt.Errorf("Failed to get the list of partitions: ", err)
		logs.Warn(err)
		return
	}

	for partition := range partitionList {
		pc, err := c.Consumer.ConsumePartition(c.Topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			err = fmt.Errorf("Failed to start consumer for partition %d: %s\n", partition, err)
			logs.Warn(err)
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {

				c.MessageChan <- string(msg.Value)
				fmt.Println(string(msg.Value))

				if c.ExitSign && string(msg.Value) == "" {
					return
				}
			}

		}(pc)
	}
	wg.Wait()
}
