package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Producer configuration
	config := kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", // Replace with your Kafka broker(s)
	}

	// Create a Kafka producer instance
	producer, err := kafka.NewProducer(&config)
	if err != nil {
		fmt.Printf("Error creating Kafka producer: %v\n", err)
		return
	}

	// Produce a message to a Kafka topic
	topic := "topics-adl1_d"
	message := "Hello, Kafka!"

	// Asynchronous produce
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		fmt.Printf("Failed to produce message: %v\n", err)
	}

	// Wait for delivery report
	go func() {
		fmt.Printf("cek event %v \n", producer.Events())
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition.Error)
				} else {
					fmt.Printf("Message delivered to topic %s [%d] at offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	// Produce delivery reports in the background
	producer.Flush(15 * 1000) // 15 seconds

	fmt.Println("Producer finished")
}
