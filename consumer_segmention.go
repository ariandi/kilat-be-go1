package main

import (
	"github.com/ariandi/gocom"
	"log"
	"strconv"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// brokers := []string{"localhost:9092", "kafka-broker2:9092"} // Replace with your Kafka broker addresses
	brokers := []string{"localhost:9092"} // Replace with your Kafka broker addresses

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Retry.Backoff = 1000 * time.Millisecond
	// config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v", err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln("Failed to close producer:", err)
		}
	}()

	topic := "test-kafka" // Replace with your Kafka topic

	for i := 0; i < 10; i++ {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder("Hello, Kafka!" + strconv.Itoa(i)),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("Failed to produce message: %v", err)
		} else {
			log.Printf("Produced message to topic %s, partition %d, offset %d\n", topic, partition, offset)
		}
	}

	// Wait for a signal to gracefully shut down the producer
	//sigchan := make(chan os.Signal, 1)
	//signal.Notify(sigchan, os.Interrupt)
	//<-sigchan
	//log.Println("Interrupt received. Shutting down...")
	gocom.Start()
}
