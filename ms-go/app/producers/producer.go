package producers

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

func InitKafkaProducer() error {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}

	p, err := kafka.NewProducer(config)
	if err != nil {
		return err
	}

	producer = p

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	return nil
}

func ProduceMessage(topic string, message interface{}) error {
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}

func CloseKafkaProducer() {
	producer.Close()
}
