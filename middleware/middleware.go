package middleware

import (
	"time"

	"EndioriteAPI/config"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginmiddleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func CustomRateLimiterMiddleware() gin.HandlerFunc {
	store := memory.NewStore()

	ipLimits := config.GetIPLimits()
	defaultLimitStr := config.GetEnv("DEFAULT_LIMIT", "5-S")

	return func(c *gin.Context) {
		ip := c.ClientIP()

		rateString, ok := ipLimits[ip]
		if !ok {
			rateString = defaultLimitStr
		}

		rate, err := limiter.NewRateFromFormatted(rateString)
		if err != nil {
			rate = limiter.Rate{
				Period: time.Second,
				Limit:  5,
			}
		}

		instance := limiter.New(store, rate)
		middleware := ginmiddleware.NewMiddleware(instance)

		middleware(c)

		if c.IsAborted() {
			return
		}

		c.Next()
	}
}
