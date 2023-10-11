package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Define your Kafka broker(s) address.
	broker := "localhost:9092" // Replace with your Kafka broker(s) address.
	topic := "chat_bot"

	// Create a Kafka consumer configuration.
	consumerConfig := &kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "my-consumer-group",
		"auto.offset.reset": "earliest",
	}

	// Create a Kafka consumer.
	consumer, err := kafka.NewConsumer(consumerConfig)
	if err != nil {
		fmt.Printf("Error creating Kafka consumer: %v\n", err)
		return
	}
	defer consumer.Close()

	// Subscribe to the Kafka topic.
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Printf("Error subscribing to topic: %v\n", err)
		return
	}

	// Consume messages.
	for {
		msg, err := consumer.ReadMessage(-1) // -1 means no timeout, blocks until a message is received.
		if err != nil {
			fmt.Printf("Error consuming message: %v\n", err)
			continue
		}

		partitionName := string(msg.Key) // Extract the partition name from the key.
		message := string(msg.Value)

		fmt.Printf("Received message on partition %s: %s\n", partitionName, message)

		// Filter messages based on the key.
		if partitionName == "6281219836581" {
			fmt.Printf("Received message with key %s: %s\n", partitionName, message)
		}

		if partitionName == "628569834394" {
			fmt.Printf("Received message with key %s: %s\n", partitionName, message)
		}
	}
}
