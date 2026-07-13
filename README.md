# E-Commerce Microservices in Go

A production-inspired E-Commerce backend built using **Golang**, **Microservices**, **gRPC**, **Kafka**, **Docker**, **Prometheus**, and **Grafana**.

This project demonstrates how modern backend systems are designed using service-oriented architecture, asynchronous communication, centralized API routing, and observability.

---

## Architecture

```
                    +------------------+
                    |   API Gateway    |
                    |      :8080       |
                    +--------+---------+
                             |
      -------------------------------------------------
      |                    |                         |
      |                    |                         |
+-------------+     +--------------+        +----------------+
| Auth Service|     | Order Service| -----> | Payment Service|
|   REST API  |     | REST + gRPC  |        | gRPC Server    |
+-------------+     +--------------+        +----------------+
                                                  |
                                                  |
                                              Kafka Event
                                                  |
                                                  ▼
                                         +----------------------+
                                         | Notification Service |
                                         | Kafka Consumer       |
                                         +----------------------+

                     +----------------------+
                     |     Prometheus       |
                     +----------+-----------+
                                |
                                ▼
                        +---------------+
                        |    Grafana    |
                        +---------------+
```

---

# Tech Stack

### Language

* Go 1.25

### API

* REST
* gRPC

### Messaging

* Apache Kafka

### Framework

* Echo

### Database

* (To be implemented)

### Monitoring

* Prometheus
* Grafana

### Infrastructure

* Docker
* Docker Compose

---

# Microservices

| Service              | Description                                                                   |
| -------------------- | ----------------------------------------------------------------------------- |
| API Gateway          | Single entry point using Reverse Proxy                                        |
| Auth Service         | JWT Authentication APIs                                                       |
| User Service         | User CRUD APIs                                                                |
| Order Service        | Creates orders and invokes Payment Service via gRPC                           |
| Payment Service      | Processes payments and publishes Kafka events                                 |
| Notification Service | Consumes Kafka events and simulates email notifications                       |
| Common               | Shared reusable library for middleware, configuration, logging and monitoring |

---

# Communication

## REST

* Client → API Gateway
* API Gateway → Services

---

## gRPC

```
Order Service
        |
        ▼
Payment Service
```

---

## Kafka

```
Payment Service
        |
        ▼
Kafka
        |
        ▼
Notification Service
```

---

# Monitoring

Implemented using Prometheus and Grafana.

Current metrics include:

* HTTP Request Count
* HTTP Request Duration
* Go Runtime Metrics
* Memory Usage
* Goroutines
* CPU Metrics
* Garbage Collection Metrics

---

# Folder Structure

```
ecommerce-microservices/

├── api-gateway/
├── auth-service/
├── common/
├── notification-service/
├── order-service/
├── payment-service/
├── product-service/
├── user-service/
│
├── monitoring/
│   ├── prometheus.yml
│   └── grafana/
│
├── docker-compose.yml
├── go.work
└── README.md
```

---

# Running the Project

## Clone

```bash
git clone https://github.com/rajabhishekmaurya/ecommerce-microservices.git

cd ecommerce-microservices
```

---

## Start Infrastructure

```bash
docker compose up -d
```

This starts:

* ZooKeeper
* Kafka
* Prometheus
* Grafana

---

## Start Services

Run each service in a separate terminal.

### API Gateway

```bash
cd api-gateway

go run cmd/main.go
```

---

### Auth Service

```bash
cd auth-service

go run cmd/main.go
```

---

### User Service

```bash
cd user-service

go run cmd/main.go
```

---

### Order Service

```bash
cd order-service

go run cmd/main.go
```

---

### Payment Service

```bash
cd payment-service

go run cmd/main.go
```

---

### Notification Service

```bash
cd notification-service

go run cmd/main.go
```

---

# Example Flow

1. Client sends request to API Gateway.
2. API Gateway forwards the request to Order Service.
3. Order Service creates an order.
4. Order Service invokes Payment Service via gRPC.
5. Payment Service processes payment.
6. Payment Service publishes a Kafka event.
7. Notification Service consumes the event.
8. Notification Service simulates sending an email notification.

---

# Observability

Prometheus

```
http://localhost:9090
```

Grafana

```
http://localhost:3000
```

---

# Features Implemented

* API Gateway
* JWT Authentication
* User CRUD APIs
* Order Service
* Payment Service (gRPC)
* Kafka Producer
* Kafka Consumer
* Reverse Proxy
* Docker Compose
* Prometheus Monitoring
* Grafana Dashboard
* Shared Common Library

---

# Planned Features

* Product Service
* Inventory Service
* Database Integration
* Redis Caching
* Distributed Tracing
* Request Correlation IDs
* Business Metrics
* Unit Tests
* Integration Tests
* GitHub Actions CI
* Kubernetes Deployment

---

# Learning Objectives

This project demonstrates:

* Microservice Architecture
* REST API Development
* gRPC Communication
* Event-Driven Architecture
* Kafka Messaging
* API Gateway Pattern
* Reverse Proxy
* Docker & Docker Compose
* Prometheus Monitoring
* Grafana Dashboards
* Production-oriented Project Structure

---

# Author

**Raj Abhishek Maurya**

Backend Engineer | Golang | Microservices | Distributed Systems | Kafka
