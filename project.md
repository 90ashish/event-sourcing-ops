# 📦 Project: “Event-Sourced Order Processing System”

## 1. Project Overview

Build a microservice-based order processing platform where every state change (`OrderCreated`, `PaymentProcessed`, `InventoryReserved`, `OrderShipped`, etc.) is modeled as an **immutable event**.

- **Event store:** Kafka topics are your source-of-truth.  
- **Read models:** Consumers project events into queryable views (e.g. PostgreSQL).  
- **Monitoring:** Instrument every service with Prometheus metrics and visualize in Grafana.

---

## 2. Tech Stack

- **Language:** Go (use the `Sarama` client)
- **Messaging:** Apache Kafka (run locally via Docker Compose)
- **Storage:**
  - Event store: Kafka topics
  - Read model: PostgreSQL or MongoDB
- **Metrics:** `prometheus/client_golang` library
- **Dashboards:** Grafana
- **Containerization:** Docker & Docker Compose
- **(Optional Phase 3):** 
  - Schema registry (Avro/Protobuf)  
  - Distributed tracing (OpenTelemetry + Jaeger)

---

## 3. System Architecture

```

 +-------------+      +--------------+      +--------------+
 | Order API   | ---> | Kafka Topic: | ---> | Order Read   |
 | (Go/Sarama) |      | orders       |      | Model (DB)   |
 +-------------+      +--------------+      +--------------+
        |                     |
        v                     v
 +-------------+      +--------------+
 | Payment     |      | Inventory    |
 | Service     | <--- | Service      |
 +-------------+      +--------------+
         \                /
          \              /
           v            v
           +------+ +------+
           |Metrics| |Logs  |
           +------+ +------+
               |        |
               v        v
          +----------------+
          | Prometheus &   |
          | Grafana        |
          +----------------+
```

---

## 4. Phased Roadmap

| Phase | Focus                              | Deliverables                                                                                 |
|-------|------------------------------------|----------------------------------------------------------------------------------------------|
| **1** | **MVP: Single-service event flow** | - Stand up Kafka locally (Docker Compose)                                                    |
|       |                                    | - Go producer: create `OrderCreated` events                                                  |
|       |                                    | - Go consumer: subscribe, print events                                                       |
| **2** | **Event sourcing & read model**    | - Persist events to Kafka topics                                                             |
|       |                                    | - Consumer writes to PostgreSQL “orders” table (view projection)                             |
|       |                                    | - Basic REST API to query orders by status                                                   |
| **3** | **Monitoring integration**         | - Add Prometheus instrumentation in each service (counters, histograms)                      |
|       |                                    | - Run Prometheus to scrape `/metrics` endpoints                                              |
|       |                                    | - Build Grafana dashboards: throughput, consumer lag, error rates                            |
| **4** | **Resilience & schema management** | - Implement retry/back-off in consumers                                                      |
|       |                                    | - Add a Schema Registry (Avro/Protobuf) for event contracts                                  |
| **5** | **Distributed tracing & alerting** | - Integrate OpenTelemetry (traces spanning API→Kafka→DB)                                     |
|       |                                    | - Configure Grafana alerts (e.g. high lag, error spikes)                                     |
| **6** | **Local container orchestration**  | - Full stack in Docker Compose (Kafka, Zookeeper, Prometheus, Grafana, PostgreSQL, services) |
| **7** | **Future extensions**              | - Add Notification service (email/SMS)                                                       |
|       |                                    | - Multi-region replication (Kafka Connect)                                                   |
|       |                                    | - High-availability & Kubernetes manifests                                                   |

## 4. Highlights
- Designed & implemented an event-sourced order processing platform in Go, leveraging Kafka as the immutable event store

- Built lightweight producer & consumer services (Sarama client) to handle OrderCreated, PaymentProcessed, InventoryReserved event flows

- Projected events into a PostgreSQL read model, enabling efficient querying of current order state

- Instrumented each microservice with Prometheus metrics (event throughput, consumer lag, error rates) and built Grafana dashboards for real-time monitoring

- Orchestrated the full stack locally via Docker Compose, ensuring reproducible developer environments

- Enhanced system resilience through retry/back-off strategies, schema versioning (Avro/Protobuf), and distributed tracing (OpenTelemetry + Jaeger)