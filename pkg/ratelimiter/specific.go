package ratelimit

import (
	"github.com/gin-gonic/gin"
	limiterPkg "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

func RefreshRateLimitMiddleware() gin.HandlerFunc {
	rate, err := limiterPkg.NewRateFromFormatted("3-M")
	if err != nil {
		panic(err)
	}

	store := memory.NewStore()
	limiterInstance := limiterPkg.New(store, rate)
	middleware := mgin.NewMiddleware(limiterInstance)

	return middleware
}
