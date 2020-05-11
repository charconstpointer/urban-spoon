package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	topic := "my-topic"
	partition := 0

	time.Sleep(20 * time.Second)
	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, partition)
	if err != nil {
		fmt.Println(err)
	}

	for {

		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		conn.WriteMessages(
			kafka.Message{Value: []byte("one!")},
			kafka.Message{Value: []byte("two!")},
			kafka.Message{Value: []byte("three!")},
		)

		time.Sleep(5000 * time.Millisecond)
		fmt.Println("ping")
	}
}
