package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic       = "kafka.learning.orders"
	broker1Addr = "192.168.27.129:9092"
	broker2Addr = "192.168.27.129:9093"
	broker3Addr = "192.168.27.129:9094"
)

func main() {
	ctx := context.Background()

	i := 0

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Addr, broker2Addr, broker3Addr},
		Topic:   topic,
	})

	for i <= 20 {
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte(fmt.Sprintf("this is a message %d", i)),
		})
		if err != nil {
			fmt.Printf("could not publish message to kafka, err: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("published message %d\n", i)
		i++

		// sleep for a second
		time.Sleep(time.Second)
	}
}
