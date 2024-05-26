package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Admiral-Simo/toll-calculator/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("failed to create consumer: %s", err)
	}

	topic := "obudata"
	partitions := []kafka.TopicPartition{
		{Topic: &topic, Partition: 0, Offset: kafka.Offset(0)},
	}

	err = consumer.Assign(partitions)
	if err != nil {
		log.Fatalf("failed to subscribe to topic: %s", err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}
		var obu *types.OBUData
		err = json.Unmarshal(msg.Value, &obu)
		if err != nil {
			fmt.Println("Error unmarsheling", err)
			continue
		}
		fmt.Println(obu)
	}
}
