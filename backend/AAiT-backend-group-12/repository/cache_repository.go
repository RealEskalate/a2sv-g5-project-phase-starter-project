package repository

import (
	"blog_api/domain"
	"time"

	"github.com/go-redis/redis"
)

type CacheRepository struct {
	redisClient *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) *CacheRepository {
	return &CacheRepository{redisClient}
}

func (r *CacheRepository) CacheData(key string, value string, expiration time.Duration) domain.CodedError {
	status := r.redisClient.Set(key, value, expiration)
	if status.Err() != nil {
		return domain.NewError(status.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

func (r *CacheRepository) InCache(key string) bool {
	status := r.redisClient.Exists(key)
	if status.Err() != nil {
		return false
	}

	return status.Val() == 1
}
