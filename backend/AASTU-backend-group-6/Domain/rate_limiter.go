package domain

import "github.com/gin-gonic/gin"

type RateLimiter interface {
	RateLimitMiddleware() gin.HandlerFunc
	FlushAfterSuccess(ip string) error
}