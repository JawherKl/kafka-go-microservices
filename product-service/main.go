package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"example.com/kafka-topics/product-service/dto"
)

const (
	KafkaServer  = "localhost:9092"
	KafkaTopic   = "orders-v1-topic"
	KafkaGroupId = "product-service"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": KafkaServer,
		"group.id":          KafkaGroupId,
		"auto.offset.reset": kafka.OffsetBeginning,
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	topic := KafkaTopic
	c.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			var order dto.Order
			err := json.Unmarshal(msg.Value, &order)
			if err != nil {
				fmt.Printf("Error decoding message: %v\n", err)
				continue
			}

			fmt.Printf("Received Order: %+v\n", order)
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
