package container

import (
	"log"

	"github.com/mohit838/inventory-managements-golang/pkg/config"
	"github.com/mohit838/inventory-managements-golang/pkg/db"
	"github.com/mohit838/inventory-managements-golang/pkg/redis"
)

type Container struct {
	DBClose func() error
}

func PkgContainer(cfg *config.Config) (*Container, error) {

	// Database Initialized
	//---------------------------------------------
	db, err := db.DbInitialized(cfg.Database)
	if err != nil {
		log.Fatalf("DB init failed: %v", err)
	}
	defer db.Close()

	// Redis Initialized
	//---------------------------------------------
	redis.RedisInitialized(cfg.Redis)

	// Initialized or declared services

	return &Container{
		// Pass services
		DBClose: db.Close,
	}, nil

}
