# ---------- STAGE 1: Build the Go binary ----------
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/inv ./cmd/api

# ---------- STAGE 2: Create minimal runtime image ----------
FROM alpine AS runtime
RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/inv /app/inv

COPY config/local.yml /app/config/local.yml

RUN addgroup -S appgroup \
 && adduser -S -G appgroup appuser \
 && chown -R appuser:appgroup /app

EXPOSE 7412

USER appuser

ENTRYPOINT ["/app/inv", "-config", "/app/config/local.yml"]
