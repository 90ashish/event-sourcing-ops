package config

import (
	"flag"
	"os"
	"strings"
)

// load Kafka addresses (env or flags)

// Config holds all runtime configuration options
type Config struct {
	KafkaBrokers []string
}

// Load parses flags and environment variables into a Config struct
func Load() *Config {
	// read a comma-separated list of brokers from flag or fallback to env var
	brokers := flag.String("brokers", os.Getenv("KAFKA_BROKERS"), "comma-separated Kafka brokers")
	flag.Parse()

	// split the single string into a slice of broker addresses
	return &Config{
		KafkaBrokers: strings.Split(*brokers, ","),
	}
}
