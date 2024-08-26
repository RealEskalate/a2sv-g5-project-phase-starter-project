package infrastructure

import (
	domain "blogs/Domain"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RateLimiter struct {
    client *redis.Client
}

func NewRateLimiter(redisClient *redis.Client) domain.RateLimiter {
    return &RateLimiter{client: redisClient}
}

func (rl *RateLimiter) RateLimitMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := c.ClientIP()
        key := fmt.Sprintf("otp_verify_%s", ip)
        
        // Increment request count in Redis
        count, err := rl.client.Incr(ctx, key).Result()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
            c.Abort()
            return
        }

        // Set expiration time for the key if it is a new key
        if count == 1 {
            rl.client.Expire(ctx, key, time.Minute*5)
        }

        if count > 5 {
            c.JSON(http.StatusTooManyRequests, domain.ErrorResponse{Message: "Too many requests" , Status: 429})
            c.Abort()
            return
        }

        c.Next()
    }
}

func (rl *RateLimiter) FlushAfterSuccess(ip string) error {
    key := fmt.Sprintf("otp_verify_%s", ip)
    return rl.client.Del(ctx, key).Err()
}
