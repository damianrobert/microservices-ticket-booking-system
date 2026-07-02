# Architecture: Ticket Booking System

## Overview
This repository contains a microservices-based Ticket Booking application designed to demonstrate cloud-native deployment, containerization, and orchestration patterns. The application simulates an event ticketing platform where users can browse events and purchase tickets. 

The architecture is deliberately split into distinct services to showcase inter-service communication (both synchronous and asynchronous), independent scalability, and polyglot persistence (using different database types for different use cases).

## System Components

| Component | Responsibility | Tech Stack | State |
| :--- | :--- | :--- | :--- |
| **Frontend UI** | Serves the web interface for users to browse and buy tickets. | TypeScript + React (Vite) | Stateless |
| **Catalog API** | Serves event details and availability. Read-heavy. | Go 1.22+ (chi router) | Stateless |
| **Catalog DB** | Stores event metadata (Name, Date, Venue). | MongoDB (NoSQL) | Stateful |
| **Order API** | Processes purchases and handles transactional logic. | Python 3.12+ (FastAPI) | Stateless |
| **Order DB** | Stores transaction records and user orders. | PostgreSQL 16 (SQL) | Stateful |
| **Message Broker** | Decouples the order process from background tasks. | RabbitMQ 3.x | Stateful |
| **Notification Worker** | Listens for new orders and simulates sending emails. | Python 3.12+ (aio-pika) | Stateless |
| **API Gateway** | Routes external traffic to the correct internal service. | Traefik (Ingress) | Stateless |

## Architecture Diagram

The system follows a standard API-driven microservices pattern with an event-driven background processor.

```text
                                +-------------------+
                                |                   |
                                |    Frontend UI    |
                                |                   |
                                +---------+---------+
                                          |
                                    (HTTP/REST)
                                          |
                                 +--------v--------+
                                 |                 |
                                 |   API Gateway   |
                                 |    (Ingress)    |
                                 |                 |
                                 +----+-------+----+
                                      |       |
                 +--------------------+       +--------------------+
                 |                                                 |
        +--------v--------+                               +--------v--------+
        |                 |                               |                 |
        |  Catalog API    |                               |    Order API    |
        |  (Reads data)   |                               | (Writes orders) |
        |                 |                               |                 |
        +--------+--------+                               +--------+--------+
                 |                                                 |
                 |                                        (Publishes Event)
        +--------v--------+                                        |
        |                 |                               +--------v--------+
        |   MongoDB       |                               |                 |
        | (Event Data)    |                               | Message Broker  |
        |                 |                               |   (RabbitMQ)    |
        +-----------------+                               |                 |
                                                          +--------+--------+
                                                                   |
                                                          (Consumes Event)
                                                                   |
        +-----------------+                               +--------v--------+
        |                 |                               |                 |
        |  PostgreSQL     | <-----------------------------+  Notification   |
        | (Order Data)    |        (Updates Status)       |     Worker      |
        |                 |                               |                 |
        +-----------------+                               +-----------------+
```

## Technology Stack

### Languages
| Language | Used By | Rationale |
| :--- | :--- | :--- |
| **TypeScript** | Frontend UI | Type-safe consumption of multiple REST APIs; industry-standard for React |
| **Go 1.22+** | Catalog API | Optimal for read-heavy, high-concurrency workloads; tiny Docker images; fast cold-start |
| **Python 3.12+** | Order API, Notification Worker | Async-first with FastAPI; rich ecosystem for business logic, ORM, and message queuing |

### Key Libraries & Frameworks

#### Frontend UI
| Package | Purpose |
| :--- | :--- |
| React 18 | UI component framework |
| Vite | Build tool and dev server |
| Axios | HTTP client for API calls |

#### Catalog API (Go)
| Package | Purpose |
| :--- | :--- |
| `chi` | Lightweight HTTP router |
| `go.mongodb.org/mongo-driver` | Official MongoDB driver |

#### Order API (Python)
| Package | Purpose |
| :--- | :--- |
| FastAPI | Async web framework with auto OpenAPI docs |
| SQLAlchemy (async) | ORM for PostgreSQL |
| `asyncpg` | High-performance async PostgreSQL driver |
| `aio-pika` | Async RabbitMQ client for publishing order events |
| Pydantic | Request/response validation and serialisation |

#### Notification Worker (Python)
| Package | Purpose |
| :--- | :--- |
| `aio-pika` | Async RabbitMQ consumer |
| `asyncpg` | Updates order status in PostgreSQL after processing |

### Infrastructure & Tooling
| Tool | Purpose |
| :--- | :--- |
| **Docker** | Containerises every service |
| **Docker Compose** | Local development orchestration |
| **Traefik** | API Gateway / reverse proxy with path-based routing |
| **RabbitMQ 3.x** | Message broker (`direct` exchange for order events) |
| **MongoDB** | Document store for event catalog data |
| **PostgreSQL 16** | Relational store for transactional order records |

### Design Decisions
- **Polyglot but minimal** — two backend languages (Go + Python). Go is chosen where raw throughput matters (reads); Python where expressive domain logic and async I/O matter (writes, workers).
- **Stateless services** — all APIs and the worker carry no in-memory state, making horizontal scaling in Kubernetes trivial.
- **Event-driven decoupling** — the Order API publishes to RabbitMQ and returns immediately; the Notification Worker processes asynchronously, preventing slow email logic from blocking the purchase flow.
