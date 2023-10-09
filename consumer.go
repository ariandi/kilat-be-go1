package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Consumer configuration
	config := kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", // Replace with your Kafka broker(s)
		"group.id":          "my-consumer-group",
		"auto.offset.reset": "earliest",
	}

	// Create a Kafka consumer instance
	consumer, err := kafka.NewConsumer(&config)
	if err != nil {
		fmt.Printf("Error creating Kafka consumer: %v\n", err)
		os.Exit(1)
	}

	// Subscribe to the Kafka topic
	topic := "purchases"
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Printf("Error subscribing to topic: %v\n", err)
		os.Exit(1)
	}

	// Handle Kafka messages
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	run := true

	for run {
		fmt.Printf("test \n")
		fmt.Printf("consumer.Events() %v \n", consumer.Events())
		select {
		case ev := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", ev)
			run = false
		default:
			ev, err := consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
		}
	}

	fmt.Println("Closing consumer...")
	consumer.Close()
}
