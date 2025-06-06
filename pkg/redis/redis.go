package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/mohit838/inventory-managements-golang/pkg/config"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func RedisInitialized(cfg config.Redis) error {

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.Db,
	})

	// Test connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("redis ping failed: %w", err)
	}

	return nil

}
