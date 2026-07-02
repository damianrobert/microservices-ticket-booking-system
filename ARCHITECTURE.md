# Architecture: Ticket Booking System

## Overview
This repository contains a microservices-based Ticket Booking application designed to demonstrate cloud-native deployment, containerization, and orchestration patterns. The application simulates an event ticketing platform where users can browse events and purchase tickets. 

The architecture is deliberately split into distinct services to showcase inter-service communication (both synchronous and asynchronous), independent scalability, and polyglot persistence (using different database types for different use cases).

## System Components

| Component | Responsibility | Typical Tech Stack | State |
| :--- | :--- | :--- | :--- |
| **Frontend UI** | Serves the web interface for users to browse and buy tickets. | React / Vue / HTML+JS | Stateless |
| **Catalog API** | Serves event details and availability. Read-heavy. | [Your Language Choice] | Stateless |
| **Catalog DB** | Stores event metadata (Name, Date, Venue). | MongoDB (NoSQL) | Stateful |
| **Order API** | Processes purchases and handles transactional logic. | [Your Language Choice] | Stateless |
| **Order DB** | Stores transaction records and user orders. | PostgreSQL (SQL) | Stateful |
| **Message Broker**| Decouples the order process from background tasks. | RabbitMQ / Redis | Stateful |
| **Notification Worker** | Listens for new orders and simulates sending emails. | [Your Language Choice] | Stateless |

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
