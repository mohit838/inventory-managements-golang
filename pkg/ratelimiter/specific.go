package ratelimit

import (
	"github.com/gin-gonic/gin"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

func RefreshRateLimitMiddleware() gin.HandlerFunc {
	rateStr := "3-M" // 3 requests per minute per IP
	rate, err := limiter.NewRateFromFormatted(rateStr)
	if err != nil {
		panic(err)
	}

	store := memory.NewStore()
	middleware := mgin.NewMiddleware(limiter.New(store, rate))

	return middleware
}
