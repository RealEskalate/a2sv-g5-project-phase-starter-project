package infrastracture

import (
    "net/http"
	"group3-blogApi/config"
    "time"

    "github.com/gin-gonic/gin"
)

type RateLimitterMiddleware struct {
    MaxRequest int
    Expiration time.Duration
    Storage    map[string]*Request
}

type Request struct {
    Count int
    Last  time.Time
}

func NewRateLimitterMiddleware() *RateLimitterMiddleware {
    maxRequest := config.EnvConfigs.RateLimitMaxRequest
    expiration := config.EnvConfigs.RateLimitExpiration
	
    return &RateLimitterMiddleware{
        MaxRequest: maxRequest,
        Expiration: time.Minute * time.Duration(expiration),
        Storage:    make(map[string]*Request),
    }
}

func (r *RateLimitterMiddleware) RateLimitter() gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := c.ClientIP()
        request, ok := r.Storage[ip]
        if !ok {
            request = &Request{
                Count: 0,
                Last:  time.Now(),
            }
            r.Storage[ip] = request
        }

        if request.Count >= r.MaxRequest {
            if time.Since(request.Last) < r.Expiration {
                c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
                return
            }
            request.Count = 0
        }

        request.Count++
        request.Last = time.Now()
        c.Next()
    }
}