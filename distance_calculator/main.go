package main

import (
	"log"
)

// Transport (HTTP, GRPC, Kafka) -> attach business logic to this transport

const consumeTopic = "obudata"
const produceTopic = "distancedata"

func main() {
	var svc CalculatorServicer
	var prod DataProducer
	prod, err := NewKafkaProducer(produceTopic)
	if err != nil {
		log.Fatal("error creating kafka producer:", err)
	}
	svc = NewCalculatorService(prod)
	svc = NewLogMiddleware(svc)
	kafkaConsumer, err := NewKafkaConsumer(consumeTopic, svc)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
