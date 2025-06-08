// @title Inventory Management API
// @version 1.0
// @description This is a sample server for Inventory Management.
// @termsOfService http://swagger.io/terms/

// @contact.name Mohit
// @contact.url http://www.mohitul-islam.com
// @contact.email connect@mohitul-islam.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3179
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mohit838/inventory-managements-golang/docs"
	"github.com/mohit838/inventory-managements-golang/logging"
	"github.com/mohit838/inventory-managements-golang/pkg/config"
	"github.com/mohit838/inventory-managements-golang/pkg/container"
	"github.com/mohit838/inventory-managements-golang/router"
)

func main() {
	docs.SwaggerInfo.Title = "Inventory Management API"
	docs.SwaggerInfo.Description = "This is a sample server for Inventory Management."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3179"
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Logging Init
	//---------------------------------------------
	logging.Init()

	// Config file loading
	//---------------------------------------------
	cfg, err := config.AppConfig()
	if err != nil {
		slog.Error("unable to load config:", "err", err)
	}

	// Setup full dependency container
	// ----------------------------------------
	c, err := container.PkgContainer(cfg)
	if err != nil {
		slog.Error("Failed to initialize container", "err", err)
		os.Exit(1)
	}
	defer c.DBClose()

	// Router setup (All routers in one file)
	//-----------------------------------------
	// Note:: Pass services from container
	// Note:: We are using Gin
	r := router.Setup(router.Deps{
		AuthService: c.AuthService,
		JWTService:  c.JWTService,
	})

	// Start the server
	//----------------------------------------
	srv := router.CreateServer(cfg, r)

	// Run server in background
	// ----------------------------------------
	go func() {
		slog.Info("Server listening", "addr", srv.Addr, "env", cfg.Env)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Unexpected server error", "err", err)
			os.Exit(1)
		}
	}()

	// Wait for OS shutdown signal
	// ----------------------------------------
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	slog.Warn("Shutdown signal received", "signal", sig)

	// Graceful shutdown with timeout
	// ----------------------------------------
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Graceful shutdown failed", "err", err)
		os.Exit(1)
	}

	slog.Info("Server exited cleanly")
}
