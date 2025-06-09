package ratelimit

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginmiddleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

func Middleware() gin.HandlerFunc {
	// Define rate: 100 requests per minute (for example)
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  100,
	}

	store := memory.NewStore()
	instance := limiter.New(store, rate)

	return ginmiddleware.NewMiddleware(instance)
}
