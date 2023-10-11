package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"sync"
)

func main() {
	// Define your Kafka broker(s) address.
	broker := "localhost:9092" // Replace with your Kafka broker(s) address.

	// Create a Kafka consumer configuration.
	consumerConfig := &kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "my-consumer-group",
		"auto.offset.reset": "earliest",
	}

	// Number of partitions to consume from.
	numPartitions := 3 // Replace with the number of partitions you want to consume from.

	// Create a WaitGroup to wait for all consumers to finish.
	var wg sync.WaitGroup

	// Create and start consumers for each partition.
	for partition := 0; partition < numPartitions; partition++ {
		wg.Add(1)
		go func(partition int) {
			defer wg.Done()

			// Create a Kafka consumer.
			consumer, err := kafka.NewConsumer(consumerConfig)
			if err != nil {
				fmt.Printf("Error creating Kafka consumer for partition %d: %v\n", partition, err)
				return
			}
			defer consumer.Close()

			// Define the topic you want to consume from.
			topic := "webhook"

			// Assign the consumer to a specific partition.
			partitionAssignment := []kafka.TopicPartition{
				{
					Topic:     &topic,
					Partition: int32(partition),
				},
			}

			err = consumer.Assign(partitionAssignment)
			if err != nil {
				fmt.Printf("Error assigning partition %d: %v\n", partition, err)
				return
			}

			// Consume messages from the assigned partition.
			for {
				msg, err := consumer.ReadMessage(-1) // -1 means no timeout, blocks until a message is received.
				if err != nil {
					fmt.Printf("Error consuming message from partition %d: %v\n", partition, err)
					continue
				}

				fmt.Printf("Received message from partition %d, offset %d: %s\n",
					partition, msg.TopicPartition.Offset, string(msg.Value))
			}
		}(partition)
	}

	// Wait for all consumers to finish.
	wg.Wait()
}
