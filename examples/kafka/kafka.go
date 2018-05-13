package kafka

import (
	"context"
	"fmt"
	kg "github.com/segmentio/kafka-go"
	"time"
)

// Fire up "hello world" for every sec seconds.
func WriteMessage(ctx context.Context, w *kg.Writer, sec int32) {
	for {
		fmt.Printf("Sending message: Hello World!..\n")
		err := w.WriteMessages(ctx, kg.Message{Value: []byte("Hello World!")})
		if err != nil {
			fmt.Printf("Write failed: %v", err)
			break
		}
		time.Sleep(time.Second * time.Duration(sec))
	}
}

// Listen.
func ListenMessage(ctx context.Context, r *kg.Reader) {
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
