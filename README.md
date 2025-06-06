# Inventory Managements Golang

Inventory management for portfolio

## Project Install

```
    go mod download
```

```
    go mod tidy
```

## Migration CMD (Table)

    migrate create -ext sql -dir ./cmd/migrate/migrations create_todo_table

## Docker CMD

```
    docker compose up -d --build
```

```
    docker logs -f <container_id>
```

```
    docker exec -it <container_id> sh
```
