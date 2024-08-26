package domain

import (
	"github.com/redis/go-redis/v9"
)
// CacheService is a struct that holds the redis client
type CacheService struct {
	Client *redis.Client
}
