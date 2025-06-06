# Inventory Management Golang

A full-featured Inventory Management system built with Go for learning and portfolio purposes.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Project Installation](#project-installation)
3. [Configuration](#configuration)

   - [YAML Config File](#yaml-config-file)
   - [Environment Variables](#environment-variables)

4. [Database Migrations](#database-migrations)
5. [Running with Docker](#running-with-docker)
6. [Running Locally (Without Docker)](#running-locally-without-docker)
7. [Redis Initialization](#redis-initialization)
8. [Project Structure Overview](#project-structure-overview)
9. [API Endpoints (Sample)](#api-endpoints-sample)
10. [Logging](#logging)
11. [Testing](#testing)
12. [Contributing](#contributing)
13. [License](#license)

---

## Prerequisites

Before you begin, ensure you have:

- **Go 1.24 (or later)** installed and `GOPATH` set.
- **MySQL** (≥ 5.7 for JSON support) accessible (either locally or via Docker).
- **Redis** (≥ 6.0) accessible if you plan to use caching or pub/sub.
- **Docker & Docker Compose** (for containerized setup).
- **golang-migrate CLI** (if you intend to run migrations manually):

  ```bash
  go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  ```

- (Optional) **make** if you’d like a Makefile for common tasks.

---

## Project Installation

```bash
# Clone the repository
git clone https://github.com/your-username/inventory-managements-golang.git
cd inventory-managements-golang

# Download Go module dependencies
go mod download

# Clean up unused dependencies
go mod tidy
```

> **Note:** Adjust the module path in `go.mod` and imports if you fork or rename the repository.

---

## Configuration

### YAML Config File

The application reads configuration from a single YAML file by default (`config/local.yml`). You can override this by:

- Setting an environment variable:

  ```bash
  export CONFIG_PATH=/path/to/your/config.yml
  ```

- Passing a flag when running:

  ```bash
  go run main.go -config=/path/to/your/config.yml
  ```

#### Example `config/local.yml` (placeholders)

```yaml
env: "dev"

app:
  port: <APP_PORT>
  name: "Inventory Management"
  version: "1.0.0"
  description: "Manage inventory items, categories, and stock levels."

database:
  driver: "mysql"
  host: "<DB_HOST>"
  port: <DB_PORT>
  db_name: "<DB_NAME>"
  username: "<DB_USER>"
  password: "<DB_PASSWORD>"
  charset: "utf8mb4"
  parse_time: true
  loc: "Local"
  max_idle_conns: 10
  max_open_conns: 50

redis:
  host: "<REDIS_HOST>"
  port: <REDIS_PORT>
  password: "<REDIS_PASSWORD>"
  db: 0

jwt:
  secret: "<JWT_SECRET>"
  expires_in: "24h"

refresh_token:
  secret: "<REFRESH_SECRET>"
  expires_in: "7d"
```

- **`env`**: Application environment (`dev`, `staging`, `prod`, etc.).
- **`app.port`**: Port on which the HTTP server listens.
- **`database.*`**: MySQL connection details.
- **`redis.*`**: Redis host, port, password, and database index.
- **`jwt.*`** / **`refresh_token.*`**: Secrets and expiration durations.

---

### Environment Variables

If you prefer environment variables over YAML, you can override:

```bash
export CONFIG_PATH=./config/production.yml
export CONFIG_ENV=prod
```

> **Tip:** You could also use a `.env` file for local development (with [joho/godotenv](https://github.com/joho/godotenv)), but the code prioritizes YAML first. If you use a `.env`, ensure it doesn’t conflict with the YAML path.

---

## Database Migrations

We use [golang-migrate](https://github.com/golang-migrate/migrate) to version and apply schema changes. Migration files live under `cmd/migrate/migrations/`.

### Creating a New Migration

From the project root, run:

```bash
migrate create -ext sql -dir ./cmd/migrate/migrations -seq create_items_table
```

This generates:

```
cmd/migrate/migrations/
    1_create_items_table.up.sql
    1_create_items_table.down.sql
```

- **Edit** `1_create_items_table.up.sql` to add your `CREATE TABLE items (…)` DDL.
- **Edit** `1_create_items_table.down.sql` to add `DROP TABLE IF EXISTS items;`.

You can name files as `00000001_items_table.up.sql` if you prefer zero-padding. Just keep the version prefix monotonic (sequential or timestamp).

### Applying Migrations

To run all pending “up” migrations:

```bash
go run cmd/migrate/main.go up
```

This does:

1. Load `config/local.yml` (or path from `CONFIG_PATH`).
2. Build a MySQL DSN from `cfg.Database`.
3. Open an `*sqlx.DB`, extract the underlying `*sql.DB`.
4. Use `migrate.NewWithDatabaseInstance` to point at migrations folder.
5. Run all pending `*.up.sql` in version order.

To roll back one step (or run `*.down.sql`):

```bash
go run cmd/migrate/main.go down
```

If you build a binary (e.g. `inventory-migrate`), you can simply run:

```bash
./inventory-migrate up
./inventory-migrate down
```

> **Tip:** If your migrations folder is elsewhere, adjust `"file://cmd/migrate/migrations"` in `main.go`.

---

## Running with Docker

We provide a `docker-compose.yml` to spin up MySQL, Redis, and the Go application. By default, it expects:

- **MySQL** at `mysql:3306`
- **Redis** at `redis:6379`
- **Go app** listening on port `3179` (mapped to host)

### Example `docker-compose.yml`

```yaml
version: "3.8"

services:
  mysql:
    image: mysql:8.0
    container_name: ims_mysql
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: <DB_PASSWORD>
      MYSQL_DATABASE: <DB_NAME>
      MYSQL_USER: <DB_USER>
      MYSQL_PASSWORD: <DB_PASSWORD>
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

  redis:
    image: redis:7.0
    container_name: ims_redis
    restart: unless-stopped
    command: redis-server --requirepass "<REDIS_PASSWORD>"
    ports:
      - "6379:6379"

  app:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: ims_app
    depends_on:
      - mysql
      - redis
    environment:
      - CONFIG_PATH=/app/config/local.yml
    ports:
      - "3179:3179"
    volumes:
      - .:/app
    command: ["sh", "-c", "go run main.go"]

volumes:
  mysql_data:
```

#### Build and Run

```bash
# Build images and start containers in background
docker compose up -d --build

# View real-time logs for Go app
docker logs -f ims_app

# Open shell inside Go app container
docker exec -it ims_app sh
```

- **CONFIG_PATH**: We set `CONFIG_PATH=/app/config/local.yml` so that inside Docker, Go loads your YAML.
- **Live Reload**: The volume `.:/app` lets you edit code locally and restart the container manually (or use a hot-reload tool like `air`).

---

## Running Locally (Without Docker)

1. **Start MySQL & Redis locally** (defaults):

   ```bash
   # MySQL
   docker run -d --name ims-mysql \
     -e MYSQL_ROOT_PASSWORD=<DB_PASSWORD> \
     -e MYSQL_DATABASE=<DB_NAME> \
     -p 3306:3306 \
     mysql:8.0

   # Redis (password optional for dev)
   docker run -d --name ims-redis \
     -p 6379:6379 \
     redis:7.0
   ```

2. **Ensure `config/local.yml` matches local setup** (e.g., host: `"localhost"`, empty Redis password if none).

3. **Initialize Redis client** (in `main.go`):

   ```go
   if err := redis.RedisInitialized(cfg.Redis); err != nil {
       slog.Error("Redis initialization failed", "err", err)
       os.Exit(1)
   }
   slog.Info("Connected to Redis successfully")
   ```

   This ensures `redis.Client` is non-nil before handlers use it.

4. **Run migrations**:

   ```bash
   go run cmd/migrate/main.go up
   ```

5. **Start the app**:

   ```bash
   go run main.go
   ```

   You should see:

   ```
   {"time":"…","level":"INFO","msg":"Connected to Redis successfully"}
   {"time":"…","level":"INFO","msg":"Server starting","addr":":3179","env":"dev"}
   ```

6. **Hit endpoints** (e.g., `GET http://localhost:3179/health`, `GET http://localhost:3179/rts`, etc.).

---

## Redis Initialization

In `pkg/redis/redis.go`:

```go
var Client *redis.Client

func RedisInitialized(cfg config.Redis) error {
    addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

    Client = redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: cfg.Password,
        DB:       cfg.Db,
    })

    // Test connection
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := Client.Ping(ctx).Err(); err != nil {
        return fmt.Errorf("redis ping failed: %w", err)
    }
    return nil
}
```

- Call this once at startup (e.g., in `main()`) so that `redis.Client` is populated.
- If you skip this, `redis.Client` remains `nil`, and any `Set`/`Get` will panic.
- In handlers, prefer using `c.Request.Context()` instead of `context.Background()` so Redis calls respect HTTP cancellations.

---

## Project Structure Overview

```
├── cmd/
│   └── migrate/
│       ├── migrations/     # SQL files: *.up.sql / *.down.sql
│       └── main.go         # Migration runner (golang-migrate)
├── config/
│   └── local.yml           # YAML config
├── pkg/
│   ├── config/             # YAML-loading logic (AppConfig)
│   ├── container/          # Dependency injection / service builders
│   ├── logging/            # slog initialization
│   ├── redis/              # Redis client initialization
│   └── …
├── router/
│   └── router.go           # Gin routes & server creation
├── Dockerfile              # How to build the Go app image
├── docker-compose.yml      # MySQL, Redis, and Go app services
├── go.mod
├── go.sum
└── main.go                 # Entry point: logging, config, DI, start server
```

- **`main.go`**

  - Initializes structured logging (`logging.Init()`).
  - Loads config (`config.AppConfig()`).
  - Calls `redis.RedisInitialized(cfg.Redis)`.
  - Constructs DI container (`container.PkgContainer(cfg)`).
  - Defers `c.DBClose()`.
  - Sets up Gin with `router.Setup()`.
  - Creates and starts the HTTP server with graceful-shutdown logic.

- **`router/router.go`**

  - `Setup()` registers routes (e.g., `/health`, `/rts`, and inventory CRUD).
  - `CreateServer()` returns an `*http.Server` preconfigured with timeouts.

- **`pkg/config`**

  - Defines `Config`, `Database`, `Redis`, `Jwt`, `RefreshToken` structs.
  - `AppConfig()` reads YAML (or `-config` flag) and unmarshals into `*Config`.

- **`pkg/redis`**

  - Exports `var Client *redis.Client`.
  - `RedisInitialized(cfg config.Redis)` populates `Client` and pings.

- **`pkg/container`**

  - Builds repositories/services (e.g., `ItemService`, `CategoryService`) and injects `*sqlx.DB` + `*redis.Client` + others.

---

## API Endpoints (Sample)

Below are sample routes you might add for inventory management. Adjust URLs, request/response bodies, and methods as needed.

### Health Check

- **Endpoint:**

  ```
  GET /health
  ```

- **Response:**

  ```json
  {
    "status": "ok"
  }
  ```

### Redis Test

- **Endpoint:**

  ```
  GET /rts
  ```

- **Behavior:**

  1. Writes key `ping` → `pong` (TTL = 1 minute) into Redis.
  2. Reads it back.
  3. Returns `{ "ping": "pong" }` on success or 500-error JSON on failure.

### Inventory CRUD (Examples)

You’ll likely create routes under `/items` and `/categories`.

#### Create Item

```
POST /items
Content-Type: application/json

{
  "name": "Wireless Mouse",
  "category_id": 3,
  "quantity": 50,
  "price": 29.99,
  "tags": ["electronics", "accessories"]
}
```

**Response:**

```json
{
  "id": 1,
  "name": "Wireless Mouse",
  "category_id": 3,
  "quantity": 50,
  "price": 29.99,
  "tags": ["electronics", "accessories"],
  "created_at": "2025-06-06T16:00:00Z"
}
```

#### List Items

```
GET /items
```

**Response:**

```json
[
  {
    "id": 1,
    "name": "Wireless Mouse",
    "category_id": 3,
    "quantity": 50,
    "price": 29.99,
    "tags": ["electronics", "accessories"],
    "created_at": "2025-06-06T16:00:00Z"
  },
  {
    "id": 2,
    "name": "Notebook",
    "category_id": 5,
    "quantity": 200,
    "price": 2.49,
    "tags": ["stationery"],
    "created_at": "2025-06-05T10:15:00Z"
  }
]
```

#### Get Single Item

```
GET /items/:id
```

**Response:**

```json
{
  "id": 1,
  "name": "Wireless Mouse",
  "category_id": 3,
  "quantity": 50,
  "price": 29.99,
  "tags": ["electronics", "accessories"],
  "created_at": "2025-06-06T16:00:00Z"
}
```

#### Update Item

```
PUT /items/:id
Content-Type: application/json

{
  "quantity": 45,
  "price": 27.99
}
```

**Response:**

```json
{
  "id": 1,
  "name": "Wireless Mouse",
  "category_id": 3,
  "quantity": 45,
  "price": 27.99,
  "tags": ["electronics", "accessories"],
  "updated_at": "2025-06-06T17:00:00Z"
}
```

#### Delete Item (Soft-Delete)

```
DELETE /items/:id
```

**Response:**

```json
{
  "message": "Item soft-deleted successfully"
}
```

#### Create Category

```
POST /categories
Content-Type: application/json

{
  "name": "Electronics",
  "description": "All electronic gadgets"
}
```

**Response:**

```json
{
  "id": 3,
  "name": "Electronics",
  "description": "All electronic gadgets",
  "created_at": "2025-06-06T16:10:00Z"
}
```

#### List Categories

```
GET /categories
```

**Response:**

```json
[
  {
    "id": 1,
    "name": "Stationery",
    "description": "Office and school supplies",
    "created_at": "2025-06-05T09:00:00Z"
  },
  {
    "id": 3,
    "name": "Electronics",
    "description": "All electronic gadgets",
    "created_at": "2025-06-06T16:10:00Z"
  }
]
```

Adjust these examples based on your service layer and request/response DTOs.

---

## Logging

This project uses Go’s `log/slog` for structured, leveled logging. In `logging.Init()`, you can:

- Configure JSON vs. console format.
- Set log level (`INFO`, `DEBUG`, `ERROR`, etc.).
- Attach global fields (e.g., `app_name`, `version`, `env`).

Example usage in a handler:

```go
slog.Info("Creating new item", "user_id", userID, "item_name", payload.Name)
slog.Error("Failed to update item", "err", err)
```

In production, redirect logs to a file or an external log aggregator. In development, console output is fine.

---

## Testing

### Unit Tests

You can write unit tests for:

- Config loading (`config.AppConfig()` with a temporary YAML).
- Service layer (e.g., `ItemService`) by mocking the database or using an in-memory SQLite for speed.
- Redis operations by pointing to a local Redis instance or using a mock client interface.

Example:

```bash
go test ./pkg/config
go test ./pkg/redis
go test ./pkg/services
```

### Integration Tests

For end-to-end testing:

1. Spin up a MySQL container and Redis container (via Docker Compose).
2. Run migrations.
3. Start the Go app in test mode (with a separate YAML, e.g., `config/test.yml`, pointing to a test database).
4. Use [httpexpect](https://github.com/gavv/httpexpect) or plain `net/http` to send requests and assert responses.

You might add a Makefile target:

```makefile
test-integration:
    docker-compose -f docker-compose.test.yml up -d
    go run cmd/migrate/main.go up -config=config/test.yml
    go test ./integration
    docker-compose -f docker-compose.test.yml down
```

---

## Contributing

1. **Fork** the repository (top right on GitHub).
2. **Branch**:

   ```bash
   git checkout -b feature/awesome-feature
   ```

3. **Make changes**, ensure lint/tests pass.
4. **Commit** with clear messages:

   ```bash
   git commit -m "feat: add stock tracking field to items"
   ```

5. **Push** to your fork and open a Pull Request (PR) against `main`.
6. **Describe** your changes and any migration or breaking changes.
7. **Wait** for review and address comments.

Follow Go idioms (`go fmt`, `golint`) and add/update tests for new features.

---

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
