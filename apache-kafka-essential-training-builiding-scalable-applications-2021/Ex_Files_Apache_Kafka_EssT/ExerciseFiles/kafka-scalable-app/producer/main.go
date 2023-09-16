package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

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

	i := 0

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Addr, broker2Addr, broker3Addr},
		Topic:   topic,
	})

	for i <= 20000 {
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte(fmt.Sprintf("this is a message %d", i)),
		},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 1)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+1)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 2)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+2)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 3)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+3)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 4)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+4)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 5)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+5)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 7)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+7)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 8)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+8)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 9)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+9)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 10)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+10)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 11)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+11)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 12)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+12)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 13)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+13)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 14)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+14)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 15)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+15)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 16)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+16)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 17)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+17)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 18)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+18)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 19)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+19)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 20)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+20)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 21)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+21)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 22)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+22)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 23)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+23)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 24)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+24)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 25)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+25)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 26)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+26)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 27)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+27)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 28)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+28)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 29)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+29)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 30)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+30)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 31)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+31)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 32)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+32)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 33)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+33)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 34)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+34)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 35)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+35)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 36)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+36)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 37)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+37)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 38)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+38)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 39)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+39)),
			},
			kafka.Message{
				Key:   []byte(strconv.Itoa(i + 40)),
				Value: []byte(fmt.Sprintf("this is a message %d", i+40)),
			},
		)
		if err != nil {
			fmt.Printf("could not publish message to kafka, err: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("published message %d\n", i)
		i += 41

		// sleep for a second
		// time.Sleep(time.Second * 2)
	}
}
