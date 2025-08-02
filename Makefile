SHELL := /bin/bash

# default brokers if none provided
BROKERS ?= localhost:9092

.PHONY: help deps build compose-up compose-down run-producer run-consumer clean

help:
	@echo "Usage:"
	@echo "  make compose-up       # Start Zookeeper + Kafka"
	@echo "  make compose-down     # Stop ZK + Kafka"
	@echo "  make deps             # Tidy & download Go modules"
	@echo "  make build            # Compile producer & consumer binaries"
	@echo "  make run-producer     # Run producer (BROKERS=$(BROKERS))"
	@echo "  make run-consumer     # Run consumer (BROKERS=$(BROKERS))"
	@echo "  make clean            # Remove binaries & docker-compose down"

deps:
	go mod tidy

build:
	mkdir -p bin
	go build -o bin/producer cmd/producer/main.go
	go build -o bin/consumer cmd/consumer/main.go

compose-up:
	docker-compose up -d

compose-down:
	docker-compose down

run-producer:
	go run cmd/producer/main.go --brokers=$(BROKERS)

run-consumer:
	go run cmd/consumer/main.go --brokers=$(BROKERS)

clean:
	rm -rf bin/
	docker-compose down

## cmd to run in mac :=
# docker context ls
# unset DOCKER_HOST
# brew install colima
# colima start
# docker context use colima //
# docker ps
# make up
# If needed, set Colima context: docker context use colima