package main

import (
	"log"
)

// Transport (HTTP, GRPC, Kafka) -> attach business logic to this transport

const consumeTopic = "obudata"

func main() {
	var svc CalculatorServicer
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)
	kafkaConsumer, err := NewKafkaConsumer(consumeTopic, svc)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
