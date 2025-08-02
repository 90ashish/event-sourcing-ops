package main

import (
	"context"
	"encoding/json"
	"event-sourcing-ops/internal/config"
	"event-sourcing-ops/internal/event"
	"event-sourcing-ops/internal/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

// CLI entrypoint for consuming & printing events

// handler implements sarama.ConsumerGroupHandler to process messages
type handler struct{}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (h *handler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (h *handler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim is called once per partition to read messages
func (h *handler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// iterate over incoming messages
	for msg := range claim.Messages() {
		var oc event.OrderCreated
		if err := json.Unmarshal(msg.Value, &oc); err != nil {
			log.Printf("unmarshal error: %v", err)
		} else {
			log.Printf("received event: %+v", oc)
		}
		// mark message as processed so Kafka can commit the offset
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	// load configuration (e.g. brokers list from flags or env)
	cfg := config.Load()

	// create a cancellable context so we can shut down gracefully
	ctx, cancel := context.WithCancel(context.Background())

	// trap SIGINT/SIGTERM signals to trigger the cancel() call
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		cancel()
	}()

	// create a new Kafka consumer group
	cons := kafka.NewConsumer(
		cfg.KafkaBrokers,   // list of broker addresses
		"order-group",      // consumer group ID
		[]string{"orders"}, // topics to subscribe to
		&handler{},         // our handler implementation
	)

	// start consuming; this blocks until ctx is canceled
	cons.Start(ctx)
}
