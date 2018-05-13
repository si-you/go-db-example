package main

import (
	"context"
	"flag"
	"fmt"
	"time"
	kg "github.com/segmentio/kafka-go"
	"github.com/si-you/go-db-example/examples/kafka"
)

var (
	brokerAddr = flag.String("broker_addr", "localhost:9092", "Kafka broker addr.")
	topic = flag.String("topic", "greetings", "Kafka topic name.")
	groupId = flag.String("group_id", "siyou", "kafka consumer group id.")
)

func main() {
	flag.Parse()

	fmt.Printf("Creating a new writer..\n")
	w := kg.NewWriter(kg.WriterConfig{
		Brokers: []string{*brokerAddr},
		Topic: *topic,
		Balancer: &kg.LeastBytes{},
	})

	fmt.Printf("Creating a new reader..\n")
	r := kg.NewReader(kg.ReaderConfig{
		Brokers: []string{*brokerAddr},
		GroupID: *groupId,
		Topic:  *topic,
		MinBytes: 1,
		MaxBytes: 10e6,
		CommitInterval: time.Second,
	})

	// Write every 10 seconds.
	fmt.Printf("Launching the writing thread..\n")
	go kafka.WriteMessage(context.Background(), w, 10)

	// Listen is blocking.
	fmt.Printf("Start lisetning the topic...\n")
	kafka.ListenMessage(context.Background(), r)
}
