package main

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	defer admin.Close()

	// Define the topic configuration
	topic := kafka.TopicSpecification{
		Topic:             "my-topic",
		NumPartitions:     3,
		ReplicationFactor: 2,
	}

	// Create the topic
	results, err := admin.CreateTopics(context.Background(), []kafka.TopicSpecification{topic})
	if err != nil {
		fmt.Printf("Failed to create topics: %v\n", err)
	} else {
		for _, result := range results {
			fmt.Printf("Topic creation result: %v\n", result)
		}
	}
}
