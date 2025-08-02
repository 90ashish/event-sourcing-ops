package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

// wraps sarama.SyncProducer

// Producer wraps a Sarama SyncProducer
type Producer struct {
	syncProducer sarama.SyncProducer // synchronous producer
	topic        string              // topic to send messages to
}

// NewProducer creates a SyncProducer and sets it up to report successes
func NewProducer(brokers []string, topic string) *Producer {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	p, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		log.Fatalf("producer init error: %v", err)
	}
	return &Producer{syncProducer: p, topic: topic}
}

// Send publishes a message with the given key and value
func (p *Producer) Send(key, value []byte) {
	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(value),
	}

	// send the message and log partition & offset on success
	partition, offset, err := p.syncProducer.SendMessage(msg)
	if err != nil {
		log.Printf("send failed: %v", err)
	} else {
		log.Printf("sent to partition %d @ offset %d", partition, offset)
	}
}
