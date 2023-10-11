package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func produceMessage(topic string, key string, value string, priority int, producer *kafka.Producer) error {
	// Determine the partition based on the priority.
	partition := kafka.PartitionAny
	if priority == 1 {
		partition = 0 // High-priority partition
	} else if priority == 2 {
		partition = 1 // Medium-priority partition
	} else {
		partition = 2 // Low-priority partition
	}

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: partition},
		Key:            []byte(key),
		Value:          []byte(value),
	}

	return producer.Produce(message, nil)
}

func main() {
	// Initialize Kafka producer configuration.
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", // Replace with your Kafka broker(s) address.
	}

	// Create a Kafka producer.
	producer, err := kafka.NewProducer(config)
	if err != nil {
		fmt.Printf("Error creating Kafka producer: %v\n", err)
		return
	}
	defer producer.Close()

	topic := "my-priority-topic"

	// Produce high-priority message.
	err = produceMessage(topic, "key-1", "High-priority message", 1, producer)
	if err != nil {
		fmt.Printf("Error producing high-priority message: %v\n", err)
	}

	// Produce medium-priority message.
	err = produceMessage(topic, "key-2", "Medium-priority message", 2, producer)
	if err != nil {
		fmt.Printf("Error producing medium-priority message: %v\n", err)
	}

	// Produce low-priority message.
	err = produceMessage(topic, "key-3", "Low-priority message", 3, producer)
	if err != nil {
		fmt.Printf("Error producing low-priority message: %v\n", err)
	}
}
