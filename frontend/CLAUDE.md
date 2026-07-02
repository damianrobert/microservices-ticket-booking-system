# Frontend UI — Claude Instructions

## Role in the System

Browser-based UI that lets users browse events and submit ticket orders.
It is the **fourth service to build**, once the backend APIs are working and reachable
through the API Gateway. Building it last means the user already understands the data
shapes the APIs return.

---

## Tech Stack

| Concern       | Choice                  |
| :---          | :---                    |
| Language      | TypeScript              |
| Framework     | React 18                |
| Build Tool    | Vite                    |
| HTTP Client   | Axios                   |
| Styling       | Plain CSS (no framework) |

---

## Responsibility

- **Events page** — fetch and display the list of events from Catalog API (`GET /api/catalog/events`)
- **Event detail page** — show a single event and an order form (`GET /api/catalog/events/:id`)
- **Order form** — submit a purchase to Order API (`POST /api/orders`), show confirmation

All API calls go through the Traefik API Gateway:
- `/api/catalog/*` → Catalog API
- `/api/orders/*` → Order API

---

## Key Concepts to Cover When Teaching

Introduce these in order:

1. **Vite + React + TypeScript project scaffold** (`npm create vite@latest`)
2. **TypeScript interfaces** for API response shapes (Event, Order)
3. **React component structure** — pages vs components, props
4. **`useEffect` + `useState`** — fetching data on mount, loading and error states
5. **Axios** — making GET and POST requests, handling errors
6. **React Router** (if multi-page) — route definitions, `useParams` for event ID
7. **Forms in React** — controlled inputs, `onSubmit`, preventing default
8. **Vite proxy config** — proxying `/api` to the API Gateway in dev mode

---

## TypeScript Interfaces (Expected API Shapes)

```ts
interface Event {
  id: string;
  name: string;
  date: string;
  venue: string;
  availableTickets: number;
}

interface Order {
  id: string;
  eventId: string;
  userEmail: string;
  quantity: number;
  status: "pending" | "confirmed" | "failed";
  createdAt: string;
}
```

---

## Environment Variables

| Variable         | Description                    | Example                   |
| :---             | :---                           | :---                      |
| `VITE_API_BASE`  | Base URL for API Gateway       | `http://localhost:80`     |

In dev, Vite's `server.proxy` can forward `/api` to the gateway directly.

---

## Status

- [ ] Vite + React + TypeScript project initialised
- [ ] Axios configured with base URL
- [ ] Events list page implemented
- [ ] Event detail page implemented
- [ ] Order form implemented and wired to Order API
- [ ] Loading and error states handled
- [ ] Dockerfile written (Nginx to serve the built static files)
