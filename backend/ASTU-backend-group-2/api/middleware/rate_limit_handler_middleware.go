package middleware

import (
	"net/http"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	custom_error "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/errors"
	"github.com/gin-gonic/gin"
)

type RateLimitMiddleware struct {
	MaxRequest     int
	Expiration     time.Duration
	RequestStorage map[string]*Request
}

type Request struct {
	Count int
	Last  time.Time
}

func NewRateLimitterMiddleware(env *bootstrap.Env) *RateLimitMiddleware {
	maxRequest := env.RateLimitMaxRequest
	expiration := env.RateLimitExpirationMin

	return &RateLimitMiddleware{
		MaxRequest:     maxRequest,
		Expiration:     time.Minute * time.Duration(expiration),
		RequestStorage: make(map[string]*Request),
	}
}

func (r *RateLimitMiddleware) RateLimitter(c *gin.Context) {
	ip := c.ClientIP()
	request, ok := r.RequestStorage[ip]
	if !ok {
		request = &Request{
			Count: 0,
			Last:  time.Now(),
		}
		r.RequestStorage[ip] = request
	}

	if request.Count >= r.MaxRequest {
		if time.Since(request.Last) < r.Expiration {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, custom_error.ErrMessage(custom_error.ErrRateLimitExceeded))
			return
		}
		request.Count = 0
	}

	request.Count++
	request.Last = time.Now()
	c.Next()

}
