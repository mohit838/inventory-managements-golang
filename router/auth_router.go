package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mohit838/inventory-managements-golang/internal/handler"
	"github.com/mohit838/inventory-managements-golang/internal/service"
	ratelimit "github.com/mohit838/inventory-managements-golang/pkg/ratelimiter"
)

func AuthRouters(r *gin.RouterGroup, authService service.AuthService) {
	authHandler := handler.NewAuthHandler(authService)

	refreshLimiter := ratelimit.RefreshRateLimitMiddleware()

	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)
	r.POST("/auth/refresh", refreshLimiter, authHandler.RefreshToken)
}
