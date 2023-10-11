package main

import (
	"fmt"
	util "github.com/ariandi/kilat-be-go1/utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	// Define your Kafka broker(s) address.
	broker := "localhost:9092" // Replace with your Kafka broker(s) address.
	configMap := &kafka.ConfigMap{"bootstrap.servers": broker}
	topic := "chat_bot"

	// Create a Kafka admin client
	adminClient, err := kafka.NewAdminClient(configMap)
	if err != nil {
		fmt.Printf("Error creating admin client: %v\n", err)
		return
	}
	defer adminClient.Close()

	// List all existing topics
	topics, err := adminClient.GetMetadata(nil, true, 5000)
	if err != nil {
		fmt.Printf("Error getting topic metadata: %v\n", err)
		return
	}

	// Check if the topic exists
	if _, exists := topics.Topics[topic]; exists {
		fmt.Printf("Topic '%s' exists.\n", topic)
	} else {
		fmt.Printf("Topic '%s' does not exist.\n", topic)
		return
	}

	// Create a Kafka producer.
	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		fmt.Printf("Error creating Kafka producer: %v\n", err)
		return
	}
	defer producer.Close()

	// Produce messages with partition names.
	go func() {
		for i := 0; i < 1000; i++ {
			message := &kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Key:            []byte("6281219836581"), // Use the partition name as the key.
				Value:          []byte("Message for partition 6281219836581 just for testing" + util.Int64ToString(int64(i))),
			}

			// Send the message to Kafka.
			err = producer.Produce(message, nil)
			if err != nil {
				fmt.Printf("Error producing message: %v\n", err)
				return
			}

			fmt.Printf("Message sent to partition: %s\n", string(message.Key))
		}
	}()

	// Produce messages with partition names.
	message2 := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte("628569834394"), // Use the partition name as the key.
		Value:          []byte("Message for partition 628569834394"),
	}

	// Send the message to Kafka.
	err = producer.Produce(message2, nil)
	if err != nil {
		fmt.Printf("Error producing message: %v\n", err)
		return
	}

	producer.Flush(10000) // Adjust the timeout as needed.

	fmt.Printf("Message sent to partition: %s\n", string(message2.Key))
}
