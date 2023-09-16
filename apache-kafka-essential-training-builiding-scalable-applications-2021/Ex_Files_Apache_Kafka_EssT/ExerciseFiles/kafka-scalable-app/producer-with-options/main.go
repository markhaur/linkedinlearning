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

	syncWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          []string{broker1Addr, broker2Addr, broker3Addr},
		Topic:            topic,
		RequiredAcks:     -1,
		CompressionCodec: kafka.Gzip.Codec(),
	})
	key := []byte("key")
	value := []byte("this is a message")
	kMessage := kafka.Message{
		Key:   key,
		Value: value,
	}

	err := syncWriter.WriteMessages(ctx, kMessage)
	if err != nil {
		fmt.Printf("error while publishing sync message: %v\n", err)
	}

	// --------------- async writer
	asyncWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          []string{broker1Addr, broker2Addr, broker3Addr},
		Topic:            topic,
		RequiredAcks:     -1,
		CompressionCodec: kafka.Gzip.Codec(),
		Async:            true,
	})

	key = []byte("async-key")
	value = []byte("this is a async message")
	kMessage = kafka.Message{
		Key:   key,
		Value: value,
	}

	err = asyncWriter.WriteMessages(ctx, kMessage)
	if err != nil {
		fmt.Printf("error while publishing async message: %v\n", err)
	}

	// adding time to wait for async message to be properly sent
	time.Sleep(time.Second * 2)

	// --------------- writer with callback
}
