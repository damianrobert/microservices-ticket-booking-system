# Order API — Claude Instructions

## Role in the System

Transactional service that processes ticket purchases. It is the **second service to build**.
It introduces PostgreSQL writes, async SQLAlchemy, Pydantic validation, and RabbitMQ publishing
— more moving parts than the Catalog API, so take extra care to explain each layer separately.

---

## Tech Stack

| Concern           | Choice                              |
| :---              | :---                                |
| Language          | Python 3.12+                        |
| Framework         | FastAPI                             |
| Validation        | Pydantic v2                         |
| ORM               | SQLAlchemy 2.x (async)              |
| DB Driver         | `asyncpg`                           |
| Database          | PostgreSQL 16                       |
| Message Publisher | `aio-pika` (RabbitMQ async client)  |
| Config            | `pydantic-settings` / env vars      |

---

## Responsibility

- `POST /orders` — validate the request, persist the order in PostgreSQL, publish an
  `order.created` event to RabbitMQ, return the created order
- `GET /orders/{id}` — return a single order by ID

This service **does not** talk to the Catalog API directly. It publishes to RabbitMQ
and leaves the rest to the Notification Worker.

---

## Key Concepts to Cover When Teaching

Introduce these in order:

1. **FastAPI project layout** (`main.py`, `routers/`, `models/`, `db/`, `schemas/`)
2. **Pydantic schemas** — request body validation, response serialisation
3. **SQLAlchemy async engine** — `create_async_engine`, `AsyncSession`, `async_sessionmaker`
4. **Defining ORM models** — `Base`, `mapped_column`, `Mapped` (SQLAlchemy 2.x style)
5. **Writing the POST /orders handler** — dependency injection for the DB session
6. **Database migrations** — introduce `Alembic` for schema versioning
7. **Publishing to RabbitMQ with aio-pika** — connection, channel, exchange, message
8. **Environment variables** — `DATABASE_URL`, `RABBITMQ_URL`, `PORT`
9. **Dockerfile** — Python slim base image, non-root user

---

## Data Model

```
orders
------
id          UUID        primary key
event_id    TEXT        ID of the event (from Catalog)
user_email  TEXT        buyer's email address
quantity    INTEGER     number of tickets
status      TEXT        'pending' | 'confirmed' | 'failed'
created_at  TIMESTAMP   set at insert time
```

---

## RabbitMQ Event Published

Exchange: `orders` (direct)
Routing key: `order.created`

Payload (JSON):
```json
{
  "order_id": "<uuid>",
  "event_id": "<string>",
  "user_email": "<string>",
  "quantity": 2
}
```

---

## Environment Variables

| Variable        | Description                        | Example                                        |
| :---            | :---                               | :---                                           |
| `PORT`          | Port the HTTP server listens on    | `8081`                                         |
| `DATABASE_URL`  | PostgreSQL async connection string | `postgresql+asyncpg://user:pass@postgres/orders` |
| `RABBITMQ_URL`  | RabbitMQ connection string         | `amqp://guest:guest@rabbitmq/`                 |

---

## Status

- [ ] Project scaffolded (`pyproject.toml` / `requirements.txt`)
- [ ] PostgreSQL connection established (async SQLAlchemy)
- [ ] Alembic configured, initial migration written
- [ ] `POST /orders` implemented
- [ ] `GET /orders/{id}` implemented
- [ ] RabbitMQ publishing wired up
- [ ] Dockerfile written
