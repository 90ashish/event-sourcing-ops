package kafka

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

// wraps sarama.ConsumerGroup

// Consumer wraps a Sarama ConsumerGroup
type Consumer struct {
	group   sarama.ConsumerGroup        // underlying consumer group client
	topics  []string                    // topics to subscribe to
	handler sarama.ConsumerGroupHandler // handler to process messages
}

// NewConsumer initializes the Kafka consumer group
func NewConsumer(brokers []string, groupID string, topics []string, handler sarama.ConsumerGroupHandler) *Consumer {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.DefaultVersion
	cg, err := sarama.NewConsumerGroup(brokers, groupID, cfg)
	if err != nil {
		log.Fatalf("consumer group init error: %v", err)
	}
	return &Consumer{group: cg, topics: topics, handler: handler}
}

// Start launches the consumer loop until the context is canceled
func (c *Consumer) Start(ctx context.Context) {
	for {
		if err := c.group.Consume(ctx, c.topics, c.handler); err != nil {
			log.Printf("consume error: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
	}
}
