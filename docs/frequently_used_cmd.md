# Frequently used CMD and Notes of Projects

## Migrates

    migrate create -ext sql -dir ./cmd/migrate/migrations -seq create_items_table

## Tables we have

    | Domain        | Tables Created                                      |
    | ------------- | --------------------------------------------------- |
    | Auth & Access | `users`, `roles`, `permissions`, `role_permissions` |
    | Inventory     | `products`, `categories`                            |
    | Sales         | `sales`, `sale_items`, `customers`                  |
    | Alerts        | `notifications`                                     |

## Feature Development Workflow (Best Practice)

    | Step | Layer          | Purpose                                |
    | ---- | -------------- | -------------------------------------- |
    | 1️⃣  | **DB Table**   | Migrate schema via SQL files           |
    | 2️⃣  | **Model**      | Map DB to Go structs (`models/`)       |
    | 3️⃣  | **DTOs**       | Separate input/output from models      |
    | 4️⃣  | **Repository** | Direct DB access logic (`repository/`) |
    | 5️⃣  | **Service**    | Business logic layer (`service/`)      |
    | 6️⃣  | **Handler**    | HTTP routes via Gin (`handler/`)       |
    | 7️⃣  | **Middleware** | Reusable route guards (`middleware/`)  |
    | 8️⃣  | **Router**     | Register endpoints                     |

## Suggested Folder Convention

    /internal/
    handler/
        auth.go
        user.go
        product.go

    service/
        auth.go
        user.go
        product.go

    repository/
        user.go
        product.go

    middleware/
        auth.go
        permission.go

    /models/
    user.go
    product.go

    /dtos/
    user.go
    auth.go
    product.go

    /pkg/
    auth/          // jwt.go, token.go
    config/
    redis/
    db/

## For swagger

    export PATH=$PATH:$(go env GOPATH)/bin
    swag --version

    echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
    source ~/.bashrc

    swag init --generalInfo cmd/api/main.go --output docs
