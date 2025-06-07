package container

import (
	"log"

	"github.com/mohit838/inventory-managements-golang/internal/repository"
	"github.com/mohit838/inventory-managements-golang/internal/service"
	"github.com/mohit838/inventory-managements-golang/pkg/config"
	"github.com/mohit838/inventory-managements-golang/pkg/db"
	"github.com/mohit838/inventory-managements-golang/pkg/redis"
)

type Container struct {
	TestService service.TestService
	DBClose     func() error
}

func PkgContainer(cfg *config.Config) (*Container, error) {

	// Database Initialized
	//---------------------------------------------
	db, err := db.DbInitialized(cfg.Database)
	if err != nil {
		log.Fatalf("DB init failed: %v", err)
	}
	

	// Redis Initialized
	//---------------------------------------------
	redis.RedisInitialized(cfg.Redis)

	// Initialized services
	testRepo := repository.NewTestRepository(db)
    testService := service.NewTestService(testRepo)

	return &Container{
		TestService: testService,
		DBClose:     db.Close,
	}, nil

}
