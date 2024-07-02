package consumers

import (
	_ "context"
	"encoding/json"
	"fmt"
	"log"
	"ms-go/app/models"
	"ms-go/app/services/products"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartKafkaConsumer() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "product-group",
		"auto.offset.reset": "earliest",
	}

	c, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}

	topic := "rails-to-go"
	c.SubscribeTopics([]string{topic}, nil)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true

	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("Message on %s: %s\n", e.TopicPartition, string(e.Value))

				var product models.Product
				err := json.Unmarshal(e.Value, &product)
				if err != nil {
					fmt.Printf("Failed to unmarshal product: %v\n", err)
					continue
				}

				_, err = products.Create(product, false)
				if err != nil {
					fmt.Printf("Failed to create product: %v\n", err)
					continue
				}

				fmt.Println("Product created successfully")
			case kafka.Error:
				fmt.Printf("Error: %v\n", e)
				run = false
			}
		}
	}

	fmt.Println("Closing consumer")
	c.Close()
}
