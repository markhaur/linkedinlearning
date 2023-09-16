package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic       = "kafka.learning.orders"
	broker1Addr = "localhost:9092"
	broker2Addr = "localhost:9093"
	broker3Addr = "localhost:9094"
	groupID     = ""
)

func main() {
	ctx := context.Background()
	kafkaReader := NewKafkaReader(topic, groupID, broker1Addr, broker2Addr, broker3Addr)
	msgC := kafkaReader.Process(ctx)

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			processMessage(id, msgC)
		}(i)
	}
	wg.Wait()
}

func processMessage(processID int, msgC <-chan kafka.Message) {
	shouldBreak := false
	for {
		select {
		case msg := <-msgC:
			fmt.Printf("received: processID: %d, key: %s, value: %s\n", processID, msg.Key, msg.Value)
		case <-time.After(10 * time.Second):
			fmt.Printf("processID: %d, closing consumer after waiting for 5s\n", processID)
			shouldBreak = true
		}
		if shouldBreak {
			break
		}
	}
}

type KafkaReader struct {
	reader *kafka.Reader
}

func NewKafkaReader(topic, groupID string, brokers ...string) *KafkaReader {
	return &KafkaReader{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			Topic:    topic,
			GroupID:  groupID,
			MinBytes: 10,
			MaxBytes: 1024,
			MaxWait:  100 * time.Millisecond,
		}),
	}
}

func (k *KafkaReader) Process(ctx context.Context) <-chan kafka.Message {
	messageC := make(chan kafka.Message)

	go func() {
		defer close(messageC)
		for i := 0; i <= 20000; i++ {
			m, err := k.reader.ReadMessage(ctx)
			if err != nil {
				log.Fatalf("error while reading message: %v\n", err)
			}
			messageC <- m
		}
	}()

	return messageC
}
