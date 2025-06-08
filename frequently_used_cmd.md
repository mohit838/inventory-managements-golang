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
