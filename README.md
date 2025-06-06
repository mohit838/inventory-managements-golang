# Inventory Managements Golang

Inventory management system for portfolio project with golang.

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

- Go 1.24 (or later) installed and `GOPATH` set.
- MySQL (Community or Docker) accessible (version ≥ 5.7 for JSON support).
- Redis (version ≥ 6.0) accessible (if you plan to use caching or pub/sub).
- Docker & Docker Compose (for containerized setup).
- `golang-migrate` CLI (if you intend to run migrations manually).

  ```bash
  go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  ```

- (Optional) `make` if you want to add Makefile targets for common tasks.

---

## Project Installation

```bash
# Clone the repo (if you haven’t already)
git clone https://github.com/your-username/inventory-managements-golang.git
cd inventory-managements-golang

# Download dependencies
go mod download

# Clean up unused dependencies
go mod tidy
```

> **Note:** The module path should match your repository. Adjust imports if you fork or rename.

---

## Configuration

### YAML Config File

This project uses a single YAML file for all configuration. By default, the code looks for `config/local.yml` unless overridden by:

- An environment variable:

  ```bash
  export CONFIG_PATH=/path/to/your/config.yml
  ```

- A command‐line flag:

  ```bash
  go run main.go -config=/path/to/your/config.yml
  ```

#### Example `config/local.yml`

```yaml
env: "dev"

app:
  port: 1234
  name: "Project Name"
  version: "1.0.0"
  description: "Project description."

database:
  driver: "mysql"
  host: "localhost"
  port: 3306
  db_name: "db_name"
  username: "root"
  password: "123456789"
  charset: "utf8mb4"
  parse_time: true
  loc: "Local"
  max_idle_conns: 10
  max_open_conns: 100

redis:
  host: "localhost"
  port: 6379
  password: "123456789"
  db: 0

jwt:
  secret: "tHISiSsERECT"
  expires_in: "24h"

refresh_token:
  secret: "tHISiSsERECTrEFRESH"
  expires_in: "7d"
```

- **`env`**: Sets the application environment (`dev`, `staging`, `prod`, etc.).
- **`app.port`**: Port on which the HTTP server listens.
- **`database.*`**: All MySQL connection details.
- **`redis.*`**: Redis host, port, password, and database index.
- **`jwt.*`** / **`refresh_token.*`**: Secrets and expiration durations.

---

### Environment Variables

If you prefer environment variables over YAML, you can override the following:

```bash
export CONFIG_PATH=./config/production.yml
export CONFIG_ENV=prod
```

> **Tip:** You could also use a `.env` file for local development (with tools like [joho/godotenv](https://github.com/joho/godotenv)), but our setup prioritizes YAML first. If you do use a `.env`, be sure it does not conflict with the YAML path.

---

## Database Migrations

We use [golang-migrate](https://github.com/golang-migrate/migrate) to version and apply schema changes. Migration files live in `cmd/migrate/migrations/`.

### Creating a New Migration

From the project root, run:

```bash
migrate create -ext sql -dir ./cmd/migrate/migrations -seq create_todo_table
```

This generates:

```
cmd/migrate/migrations/
    1_create_todo_table.up.sql
    1_create_todo_table.down.sql
```

- **Edit** `1_create_todo_table.up.sql` to add your `CREATE TABLE todos (…)` DDL.
- **Edit** `1_create_todo_table.down.sql` to add `DROP TABLE IF EXISTS todos;`.

You can name the files however you like (e.g. `00000001_todo_table.up.sql`), but make sure the version prefix is monotonic (sequential or timestamp).

### Applying Migrations

If you prefer running migrations manually:

```bash
go run cmd/migrate/main.go up
```

This does the following:

1. Loads `config/local.yml` (or the path specified by `CONFIG_PATH`).
2. Builds a MySQL DSN from `cfg.Database`.
3. Opens an `*sqlx.DB` and extracts the underlying `*sql.DB`.
4. Uses `migrate.NewWithDatabaseInstance` to point at your migration folder.
5. Runs all pending `*.up.sql` scripts in order.

To roll back one step (or apply the next `*.down.sql`):

```bash
go run cmd/migrate/main.go down
```

If you want to generate SQL without applying it, you can also build a binary and use:

```bash
./todo-migrate up
./todo-migrate down
```

> **Tip:** If your DDL folder is somewhere else, adjust the `"file://cmd/migrate/migrations"` path in `main.go` accordingly.

---

## Running with Docker

We provide a `docker-compose.yml` to spin up MySQL, Redis, and the Go application itself. By default, it expects:

- **MySQL** at `mysql:3306`
- **Redis** at `redis:6379`
- **Go app** listening on port `1234` (mapped to your host)

### Example `docker-compose.yml`

```yaml
version: "3.8"

services:
  mysql:
    image: mysql:8.0
    container_name: ims_mysql
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: 123456789
      MYSQL_DATABASE: db_name
      MYSQL_USER: root
      MYSQL_PASSWORD: 123456789
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

  redis:
    image: redis:7.0
    container_name: ims_redis
    restart: unless-stopped
    command: redis-server --requirepass "123456789"
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
# Build images and start containers in the background
docker compose up -d --build

# View real-time logs for your Go app
docker logs -f ims_app

# Open a shell inside the running Go app container
docker exec -it ims_app sh
```

- **Configuration Path**: We set `CONFIG_PATH=/app/config/local.yml` so that inside the container, the Go process loads `config/local.yml`.
- **Live Reload** (optional): The volume `.:/app` can be helpful during development. If you modify Go code, you can manually restart the container or use a tool like `air` for hot reloading.

---

## Running Locally (Without Docker)

1. **Start MySQL & Redis locally** (assuming defaults):

   ```bash
   # MySQL
   docker run -d --name ims-mysql -e MYSQL_ROOT_PASSWORD=123456789 -e MYSQL_DATABASE=db_name -p 3306:3306 mysql:8.0

   # Redis (no password for quick dev)
   docker run -d --name ims-redis -p 6379:6379 redis:7.0
   ```

2. **Ensure `config/local.yml` matches your local setup** (e.g., `host: "localhost"`, `password: ""` for Redis if no auth).

3. **Initialize Redis client** (inside code):

   In `main.go`, you should have something like:

   ```go
   if err := redis.RedisInitialized(cfg.Redis); err != nil {
       slog.Error("Redis initialization failed", "err", err)
       os.Exit(1)
   }
   slog.Info("Connected to Redis successfully")
   ```

   This guarantees `redis.Client` is non‐nil before any handler uses it.

4. **Run migrations**:

   ```bash
   go run cmd/migrate/main.go up
   ```

5. **Start the app**:

   ```bash
   go run main.go
   ```

   You should see logs like:

   ```
   {"time":"…","level":"INFO","msg":"Connected to Redis successfully"}
   {"time":"…","level":"INFO","msg":"Server starting","addr":":3179","env":"dev"}
   ```

6. **Hit your endpoints** (e.g., `GET http://localhost:3179/health`, `GET http://localhost:3179/rts`, etc.).

---

## Redis Initialization

In `pkg/redis/redis.go`, we define:

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

- Call this once at startup (e.g. in `main()`) so that `redis.Client` is initialized.
- If you skip this call, `redis.Client` remains `nil`, and any attempt to `Set`/`Get` will panic.
- It’s best practice to use `c.Request.Context()` in handlers so that Redis calls respect HTTP cancellations.

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
│   ├── config/             # YAML‐loading logic (AppConfig)
│   ├── container/          # Dependency injection / service builders
│   ├── logging/            # slog or zap initialization
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

  - Initializes structured logging (e.g. `slog.Init()`).
  - Loads config (`config.AppConfig()`).
  - Calls `redis.RedisInitialized(cfg.Redis)`.
  - Constructs the DI container (`container.PkgContainer(cfg)`).
  - Defers `c.DBClose()`.
  - Sets up Gin with `router.Setup()`.
  - Creates and starts the HTTP server with graceful‐shutdown logic.

- **`router/router.go`**

  - `Setup()` registers public routes (e.g. `/health`, `/rts`, plus you’ll likely add your CRUD endpoints here).
  - `CreateServer()` returns an `*http.Server` preconfigured with read/write timeouts.

- **`pkg/config`**

  - Defines `Config`, `Database`, `Redis`, `Jwt`, `RefreshToken` structs.
  - `AppConfig()` reads YAML or a `-config` flag and unmarshals into `*Config`.

- **`pkg/redis`**

  - Exports `var Client *redis.Client`.
  - `RedisInitialized(cfg config.Redis)` populates `Client` and tests a PING.

- **`pkg/container`**

  - Builds any repositories/services (e.g. `TodoService`, `UserService`) and injects `*sqlx.DB` + `*redis.Client` + any other dependencies.

---

## API Endpoints (Sample)

Below are sample routes you might find useful as you expand your REST API. Adjust URLs, request/response bodies, and HTTP methods as needed.

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
  3. Returns `{ "ping": "pong" }` on success or a 500‐error JSON on failure.

### Todos CRUD (Examples)

You’ll likely add a `TodoHandler` under `/todos` or a similar grouping:

- **Create Todo**

  ```
  POST /todos
  Content-Type: application/json

  {
    "title": "Buy groceries",
    "priority": "high",
    "tags": ["shopping", "errands"]
  }
  ```

  **Response:**

  ```json
  {
    "id": 1,
    "title": "Buy groceries",
    "completed": false,
    "is_active": true,
    "priority": "high",
    "tags": ["shopping", "errands"],
    "created_at": "2025-06-06T16:00:00Z"
  }
  ```

- **List Todos**

  ```
  GET /todos
  ```

  **Response:**

  ```json
  [
    {
      "id": 1,
      "title": "Buy groceries",
      "completed": false,
      "is_active": true,
      "priority": "high",
      "tags": ["shopping", "errands"],
      "created_at": "2025-06-06T16:00:00Z"
    },
    {
      "id": 2,
      "title": "Send emails",
      "completed": true,
      "is_active": true,
      "priority": "medium",
      "tags": ["work"],
      "created_at": "2025-06-05T10:15:00Z"
    }
  ]
  ```

- **Get Single Todo**

  ```
  GET /todos/:id
  ```

  **Response:**

  ```json
  {
    "id": 1,
    "title": "Buy groceries",
    "completed": false,
    "is_active": true,
    "priority": "high",
    "tags": ["shopping", "errands"],
    "created_at": "2025-06-06T16:00:00Z"
  }
  ```

- **Update Todo**

  ```
  PUT /todos/:id
  Content-Type: application/json

  {
    "title": "Buy groceries and supplies",
    "completed": true
  }
  ```

  **Response:**

  ```json
  {
    "id": 1,
    "title": "Buy groceries and supplies",
    "completed": true,
    "is_active": true,
    "priority": "high",
    "tags": ["shopping", "errands"],
    "updated_at": "2025-06-06T17:00:00Z"
  }
  ```

- **Delete Todo** (soft‐delete)

  ```
  DELETE /todos/:id
  ```

  **Response:**

  ```json
  {
    "message": "Todo soft-deleted successfully"
  }
  ```

Adjust the above based on your service layer and how you structure request/response DTOs.

---

## Logging

This project uses Go’s `log/slog` for structured, leveled logging. In `logging.Init()`, you can:

- Configure JSON vs. console format.
- Set log level (`INFO`, `DEBUG`, `ERROR`, etc.).
- Attach global fields (e.g. `app_name`, `version`, `env`).

Example usage in a handler:

```go
slog.Info("Creating new todo", "user_id", userID, "title", payload.Title)
slog.Error("Failed to save todo", "err", err)
```

In production, redirect logs to a file or an external log aggregator. In development, console output is fine.

---

## Testing

### Unit Tests

You can write unit tests for:

- Config loading (`config.AppConfig()` with a temporary YAML).
- Service layer (e.g. `TodoService`) by mocking the database or using an in-memory SQLite for speed.
- Redis operations by pointing to a local Redis instance or using a mock client interface.

Example:

```bash
go test ./pkg/config
go test ./pkg/redis
go test ./pkg/services
```

### Integration Tests

For end-to-end testing:

1. Spin up a MySQL container and Redis container (using Docker Compose).
2. Run migrations.
3. Start the Go app in test mode (maybe with a separate YAML `config/test.yml` that points to a test database).
4. Use a tool like [httpexpect](https://github.com/gavv/httpexpect) or plain `net/http` to send requests to your endpoints and assert responses.

You might add a `Makefile` target:

```makefile
test-integration:
    docker-compose -f docker-compose.test.yml up -d
    go run cmd/migrate/main.go up -config=config/test.yml
    go test ./integration
    docker-compose -f docker-compose.test.yml down
```

---

## Contributing

1. Fork the repository (top right on GitHub).
2. Create a topic branch:

   ```bash
   git checkout -b feature/awesome-feature
   ```

3. Make your changes. Ensure linting/tests pass.
4. Commit with clear messages:

   ```bash
   git commit -m "feat: add priority enum to todos"
   ```

5. Push to your fork and open a Pull Request (PR) against `main`.
6. Describe your changes, why they’re needed, and any migration or breaking changes.
7. Wait for review and address comments.

Please follow Go idioms (e.g., `go fmt`, `golint`) and add or update tests for new features.

---

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

---

With these additions, any new contributor or user should be able to:

- Clone and set up the project
- Understand where to configure environment details (YAML vs. `.env`)
- Run migrations and start the server (both with and without Docker)
- Verify Redis connectivity
- See sample endpoints that illustrate how to build CRUD routes
- Write and execute tests
- Contribute in a structured manner
