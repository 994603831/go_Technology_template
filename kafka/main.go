package main

import (
	"fmt"
	"time"

	"github.com/shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	//生产者
	config.Producer.RequiredAcks = sarama.WaitForAll
	//分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	//创建生产者对象
	client, err := sarama.NewSyncProducer([]string{"169.254.166.74:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err: ", err)
		return
	}

	defer client.Close()
	for {
		msg := &sarama.ProducerMessage{}
		msg.Topic = "nginx_log"
		msg.Value = sarama.StringEncoder("this is a good test, my message is good")

		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed, err: ", err)
			return
		}

		fmt.Printf("pid: %v offset: %v\n", pid, offset)
		time.Sleep(10 * time.Millisecond)
	}

}
