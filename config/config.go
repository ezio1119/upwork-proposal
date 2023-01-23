package config

import "os"

const (
	defaultKafkaEndpoint = "kafka:9092"
	defaultKafkaTopic    = "example-topic"
)

var (
	KafkaEndpoint string
	KafkaTopic    string
)

func init() {
	if KafkaEndpoint = os.Getenv("KAFKA_ENDPOINT"); KafkaEndpoint == "" {
		KafkaEndpoint = defaultKafkaEndpoint
	}

	if KafkaTopic = os.Getenv("KAFKA_TOPIC"); KafkaTopic == "" {
		KafkaTopic = defaultKafkaTopic
	}
}
