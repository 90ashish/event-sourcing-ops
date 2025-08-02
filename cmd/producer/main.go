package main

import (
	"encoding/json"
	"event-sourcing-ops/internal/config"
	"event-sourcing-ops/internal/event"
	"event-sourcing-ops/internal/kafka"
	"time"
)

// CLI entrypoint for producing OrderCreated events

func main() {
	// load brokers from flags or env
	cfg := config.Load()

	// create a new Kafka producer targeting the "orders" topic
	prod := kafka.NewProducer(cfg.KafkaBrokers, "orders")

	// construct a new OrderCreated event
	oc := event.OrderCreated{
		ID:        "order-123", // unique order identifier
		Status:    "created",   // initial status
		TimeStamp: time.Now(),  // current time
	}

	// marshal the struct to JSON bytes
	data, _ := json.Marshal(oc)

	// send the event to Kafka (keyed by order ID)
	prod.Send([]byte(oc.ID), data)
}
