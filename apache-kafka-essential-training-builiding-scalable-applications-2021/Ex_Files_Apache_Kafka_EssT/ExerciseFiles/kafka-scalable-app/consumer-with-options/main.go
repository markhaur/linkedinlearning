package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic       = "kafka.learning.orders"
	broker1Addr = "localhost:9092"
	broker2Addr = "localhost:9093"
	broker3Addr = "localhost:9094"
)

func main() {
	ctx := context.Background()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{broker1Addr, broker2Addr, broker3Addr},
		Topic:          topic,
		GroupID:        "kafka-options-consumers",
		MinBytes:       10,
		MaxWait:        100 * time.Millisecond,
		MaxBytes:       1024,
		CommitInterval: 5,
	})
	defer reader.Close()

	for i := 0; i <= 200; i++ {
		message, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("error while reading the message %v\n", err)
		}
		fmt.Printf("received => key: %s, Value: %s\n", string(message.Key), string(message.Value))
	}
}
