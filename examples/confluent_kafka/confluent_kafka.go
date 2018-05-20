package confluent_kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

// Fire up "hello world" for every sec seconds.
func WriteMessage(p *kafka.Producer, t string, sec int32) {
	for {
		fmt.Printf("Sending message: Hello World!..\n")

		c := make(chan kafka.Event)
		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &t,
				Partition: kafka.PartitionAny},
			Value: []byte("Hello World!")}, c)

		if err != nil {
			fmt.Printf("Produce failed: %v\n", err)
			break
		}

		e := <-c
		fmt.Printf("channel: %v\n", e)
		close(c)

		time.Sleep(time.Second * time.Duration(sec))
	}
}

// Listen.
func ListenMessage(c *kafka.Consumer) {
	for {
		mv := c.Poll(1000)
		if mv == nil {
			continue
		}
		switch m := mv.(type) {
		case *kafka.Message:
			fmt.Printf("%s: %s\n", m.TopicPartition, string(m.Value))
		default:
			fmt.Printf("Pass %v\n", m)
		}
	}
}
