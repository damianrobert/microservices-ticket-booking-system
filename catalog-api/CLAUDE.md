# Catalog API — Claude Instructions

## Role in the System

Read-heavy service that serves event data to the Frontend UI via the API Gateway.
It is the **first service to build** in this project — it introduces Go HTTP patterns
and MongoDB without any async messaging complexity.

---

## Tech Stack

| Concern      | Choice                          |
| :---         | :---                            |
| Language     | Go 1.22+                        |
| Router       | `chi` (lightweight HTTP router) |
| Database     | MongoDB                         |
| DB Driver    | `go.mongodb.org/mongo-driver`   |
| Config       | Environment variables           |

---

## Responsibility

- `GET /events` — return a list of all events (name, date, venue, available tickets)
- `GET /events/{id}` — return a single event by ID

This service **only reads** from MongoDB. It never writes. It does not talk to the
Order API, RabbitMQ, or PostgreSQL.

---

## Key Concepts to Cover When Teaching

Introduce these in order — do not skip ahead:

1. **Go module setup** (`go mod init`, `go.mod`, `go.sum`)
2. **Project layout** (`main.go`, `handlers/`, `db/`, `models/`)
3. **chi router** — registering routes, handler function signature
4. **Connecting to MongoDB** — `mongo.Connect`, context usage, connection pooling
5. **Handler pattern** — reading path params, querying MongoDB, writing JSON response
6. **Environment variables** — reading `MONGO_URI`, `PORT` from the environment
7. **Dockerfile** — multi-stage build, distroless final image

---

## Environment Variables

| Variable    | Description                      | Example                              |
| :---        | :---                             | :---                                 |
| `PORT`      | Port the HTTP server listens on  | `8080`                               |
| `MONGO_URI` | MongoDB connection string        | `mongodb://mongo:27017`              |
| `MONGO_DB`  | Database name                    | `catalog`                            |

---

## Status

- [x] Project scaffolded (`go mod init`)
- [x] MongoDB connection established
- [ ] `GET /events` implemented
- [ ] `GET /events/{id}` implemented
- [ ] Dockerfile written
- [ ] Seeded with sample event data
