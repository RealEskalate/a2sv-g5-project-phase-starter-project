package middlewares

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func RateLimitMiddleware(redisClient *redis.Client, ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "rate_limit:" + ip
		limit := 30
		window := time.Minute

		count, err := redisClient.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		if count >= limit {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}

		log.Println("IP:", ip, "Count:", count)

		redisClient.Incr(ctx, key)
		redisClient.Expire(ctx, key, window)
		c.Next()
	}
}
