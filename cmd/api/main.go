package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mohit838/inventory-managements-golang/logging"
	"github.com/mohit838/inventory-managements-golang/pkg/config"
	"github.com/mohit838/inventory-managements-golang/pkg/container"
	"github.com/mohit838/inventory-managements-golang/router"
)

func main() {
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
	r := router.Setup(router.Deps{TestService: c.TestService})

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
