package main

import (
	"context"
	"fmt"
	"log"
	"upwork-proposal/config"

	"github.com/segmentio/kafka-go"
)

func main() {

	// TODO: Use group consumer
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{config.KafkaEndpoint},
		Topic:     config.KafkaTopic,
		Partition: 0,
		// MinBytes:  10e3, // 10KB
		// MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(0)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		// TODO: Call external API with goroutine
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
