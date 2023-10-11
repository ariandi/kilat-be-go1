package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Define your Kafka broker(s) address.
	broker := "localhost:9092" // Replace with your Kafka broker(s) address.

	// Create a Kafka producer configuration.
	producerConfig := &kafka.ConfigMap{
		"bootstrap.servers": broker,
	}

	// Create a Kafka producer.
	producer, err := kafka.NewProducer(producerConfig)
	if err != nil {
		fmt.Printf("Error creating Kafka producer: %v\n", err)
		return
	}
	defer producer.Close()

	// Define the topic and message you want to send.
	topic := "test-kafka2"
	messageValue := "Hello, Kafka!"
	partition := 2 // Specify the desired partition number.

	// Create a Kafka message with the desired partition.
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(partition)},
		Value:          []byte(messageValue),
	}

	// Send the message to Kafka.
	err = producer.Produce(message, nil)
	if err != nil {
		fmt.Printf("Error sending message to Kafka: %v\n", err)
		return
	}

	fmt.Printf("Message sent to partition %d: %s\n", partition, messageValue)

	// Wait for the message to be delivered (optional).
	deliveryReport := <-producer.Events()
	m := deliveryReport.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Message delivered to partition %d at offset %v\n",
			m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
}
