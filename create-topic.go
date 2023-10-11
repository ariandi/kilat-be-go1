package main

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Create a Kafka admin client
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		fmt.Printf("Error creating admin client: %v\n", err)
		return
	}
	defer adminClient.Close()

	// Specify the topic configuration
	topicConfig := &kafka.TopicSpecification{
		Topic:             "chat_bot",
		NumPartitions:     1, // You can adjust the number of partitions as needed.
		ReplicationFactor: 1, // You can adjust the replication factor as needed.
	}

	ctx := context.Background()

	// Create the topic
	topics := []kafka.TopicSpecification{*topicConfig}
	_, err = adminClient.CreateTopics(ctx, topics, kafka.SetAdminOperationTimeout(5000))
	if err != nil {
		fmt.Printf("Error creating topic: %v\n", err)
	} else {
		fmt.Println("Topic created successfully!")
	}
}
