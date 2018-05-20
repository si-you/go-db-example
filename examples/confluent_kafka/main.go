package main

import (
	"flag"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	ck "github.com/si-you/go-db-example/examples/confluent_kafka"
)

var (
	brokerAddr = flag.String("broker_addr", "localhost:9092", "Kafka broker addr.")
	topic      = flag.String("topic", "greetings", "Kafka topic name.")
	groupId    = flag.String("group_id", "siyou", "kafka consumer group id.")
)

func main() {
	flag.Parse()

	fmt.Printf("Creating a new producer..\n")
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": *brokerAddr})
	if err != nil {
		fmt.Printf("Failed to create a new producer: %v\n", err)
	}

	fmt.Printf("Creating a new consumer..\n")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  *brokerAddr,
		"group.id":           *groupId,
		"session.timeout.ms": 6000})
	if err != nil {
		fmt.Printf("Failed to create a new consumer: %v\n", err)
	}

	// Write every 10 seconds.
	fmt.Printf("Launching the writing thread..\n")
	go ck.WriteMessage(p, *topic, 10)

	// Listen is blocking.
	fmt.Printf("Start lisetning the topic...\n")
	err = c.Subscribe(*topic, nil)
	if err != nil {
		fmt.Printf("Failed to subscribe topic: %v\n", err)
	}
	ck.ListenMessage(c)
}
