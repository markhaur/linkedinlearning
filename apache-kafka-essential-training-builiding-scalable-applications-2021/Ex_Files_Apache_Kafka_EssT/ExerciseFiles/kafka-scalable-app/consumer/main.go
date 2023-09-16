package main

import (
	"context"
	"fmt"

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
		Brokers: []string{broker1Addr, broker2Addr, broker3Addr},
		Topic:   topic,
		GroupID: "kafka-go-consumer",
	})

	for i := 0; i <= 200; i++ {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("error while reading message: %v\n", err)
			break
		}

		fmt.Printf("received: %s\n", string(msg.Value))
	}
}
