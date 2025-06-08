package router

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mohit838/inventory-managements-golang/internal/middleware"
	"github.com/mohit838/inventory-managements-golang/pkg/config"
	"github.com/mohit838/inventory-managements-golang/pkg/redis"
)

func Setup(d Deps) *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Public routes
	//----------------------------------------------
	// Server health route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Redis test
	r.GET("/rts", func(c *gin.Context) {
		ctx := context.Background()
		err := redis.Client.Set(ctx, "ping", "pong", time.Minute).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis write failed"})
			return
		}
		val, err := redis.Client.Get(ctx, "ping").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis read failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"ping": val})
	})

	// All project router will add here
	//----------------------------------------------
	// Public routes
	authRoutes := r.Group("/api/v1")
	AuthRouters(authRoutes, d.AuthService)

	// Protected routes
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(d.JWTService))
	{
		protected.GET("/me", func(c *gin.Context) {
			userID := c.MustGet("userID").(int64)
			c.JSON(http.StatusOK, gin.H{"user_id": userID})
		})
	}

	return r
}

// Create server
// -------------------------------------------------
func CreateServer(cfg *config.Config, handler http.Handler) *http.Server {
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	return &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
