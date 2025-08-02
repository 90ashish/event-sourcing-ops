event-sourcing/
├── cmd/
│   ├── producer/
│   │   └── main.go          # CLI entrypoint for producing OrderCreated events
│   └── consumer/
│       └── main.go          # CLI entrypoint for consuming & printing events
├── internal/
│   ├── config/
│   │   └── config.go        # load Kafka addresses (env or flags)
│   ├── event/
│   │   └── order.go         # OrderCreated struct, JSON (de)serialization
│   └── kafka/
│       ├── producer.go      # wraps sarama.SyncProducer
│       └── consumer.go      # wraps sarama.ConsumerGroup
├── docker-compose.yml       # Kafka & Zookeeper for local dev
├── go.mod
└── README.md
