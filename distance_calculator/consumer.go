package main

import (
	"encoding/json"

	"github.com/Admiral-Simo/toll-calculator/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer    *kafka.Consumer
	isRunning   bool
	calcService CalculatorServicer
}

func NewKafkaConsumer(topic string, svc CalculatorServicer) (*KafkaConsumer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "distance_calculator",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}

	err = consumer.Subscribe("obudata", nil)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumer:    consumer,
		isRunning:   false,
		calcService: svc,
	}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info("kafka transport started")
	c.isRunning = true
	c.ReadMessageLoop()
}

func (c *KafkaConsumer) ReadMessageLoop() {
	for c.isRunning {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			logrus.Errorf("kafka consume error: %s", err)
			continue
		}
		var obu *types.OBUData
		if err = json.Unmarshal(msg.Value, &obu); err != nil {
			logrus.Errorf("JSON serialization error: %s", err)
			continue
		}
		distance, err := c.calcService.CalculateDistance(obu)
		if err != nil {
			logrus.Errorf("calculation error: %s", err)
			continue
		}
		_ = distance
	}
}

func (c *KafkaConsumer) Stop() {
	logrus.Info("Stopping Kafka consumer")
	c.isRunning = false
	c.consumer.Close()
}
