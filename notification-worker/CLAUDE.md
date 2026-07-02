# Notification Worker — Claude Instructions

## Role in the System

Background worker that consumes `order.created` events from RabbitMQ, simulates sending
a confirmation email, then updates the order status in PostgreSQL to `confirmed`.

It is the **third service to build**. It has no HTTP server — it is a pure async consumer
process. This is a good place to discuss worker process design and graceful shutdown.

---

## Tech Stack

| Concern           | Choice                             |
| :---              | :---                               |
| Language          | Python 3.12+                       |
| Message Consumer  | `aio-pika` (RabbitMQ async client) |
| DB Driver         | `asyncpg`                          |
| Database          | PostgreSQL 16 (shared with Order API, read for order lookup / write for status update) |
| Config            | Environment variables              |

---

## Responsibility

1. Connect to RabbitMQ on startup
2. Declare queue bound to the `orders` exchange with routing key `order.created`
3. For each message received:
   - Parse the JSON payload
   - Simulate sending an email (log the action, no real email sent)
   - Update the matching order's `status` to `confirmed` in PostgreSQL
   - Acknowledge the message (so RabbitMQ removes it from the queue)
4. On error: reject and requeue the message (or dead-letter it)

---

## Key Concepts to Cover When Teaching

Introduce these in order:

1. **Why a separate worker?** — decoupling, back-pressure, retries without blocking the HTTP path
2. **aio-pika consumer pattern** — `connect_robust`, `Queue.consume`, message callbacks
3. **Message acknowledgement** — ack vs nack, what happens on crash before ack
4. **Direct asyncpg usage** — when a full ORM is overkill for a worker
5. **Graceful shutdown** — catching `SIGTERM`, waiting for in-flight messages to complete
6. **Dockerfile** — same Python slim base as Order API, no exposed port needed

---

## RabbitMQ Event Consumed

Exchange: `orders` (direct)
Queue: `notification.order.created`
Routing key: `order.created`

Expected payload (JSON):
```json
{
  "order_id": "<uuid>",
  "event_id": "<string>",
  "user_email": "<string>",
  "quantity": 2
}
```

---

## PostgreSQL Write

Updates a single row in the `orders` table (owned by Order API):

```sql
UPDATE orders SET status = 'confirmed' WHERE id = $1;
```

---

## Environment Variables

| Variable        | Description                        | Example                                          |
| :---            | :---                               | :---                                             |
| `RABBITMQ_URL`  | RabbitMQ connection string         | `amqp://guest:guest@rabbitmq/`                   |
| `DATABASE_URL`  | PostgreSQL async connection string | `postgresql://user:pass@postgres/orders`         |

---

## Status

- [ ] Project scaffolded (`pyproject.toml` / `requirements.txt`)
- [ ] RabbitMQ consumer connected and queue declared
- [ ] Message handler implemented (parse → log → update DB → ack)
- [ ] Graceful shutdown on SIGTERM implemented
- [ ] Dockerfile written
