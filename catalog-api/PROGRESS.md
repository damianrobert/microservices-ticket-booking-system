# catalog-api — Progress

## Status: In progress

## Steps

- [x] Project scaffolded (`go mod init`)
- [x] MongoDB connection established
- [ ] `GET /events` implemented
- [ ] `GET /events/{id}` implemented
- [ ] Dockerfile written
- [ ] Seeded with sample event data

## Log

_Entries are appended by Claude after each confirmed implementation step._

<!-- log entries appear below this line -->

### [2026-07-02] Project scaffolded (go mod init)
**What was built:** `catalog-api` now has a `go.mod` file declaring `module catalog-api`, making it a valid Go module.
**Key concept:** The module path is the root namespace for internal imports — Go resolves `catalog-api/handlers` to `<module root>/handlers` on disk, no relative paths needed.
**Next step:** MongoDB connection established

### [2026-07-02] MongoDB connection established
**What was built:** `db/mongo.go` now has a `Connect()` function using `mongo-driver/v2` that builds a `*mongo.Client` and `Ping`s it with a 5-second timeout; `main.go` calls it at startup and defers `Disconnect`.
**Key concept:** `context.Context` carries cancellation/timeouts through driver calls — a `WithTimeout` context makes `Ping` fail fast instead of hanging when Mongo is unreachable, verified by testing both with and without the container running.
**Next step:** `GET /events` implemented
