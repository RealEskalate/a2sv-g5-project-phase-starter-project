package repository

import (
	"blog_api/domain"
	"time"

	"github.com/go-redis/redis"
)

/* A struct that takes a `*redis.Client` and implements the `domain.CacheRepositoryInterface` interface */
type CacheRepository struct {
	cacheClient *redis.Client
}

/* Creates and returns a `CacheRepository` instance */
func NewCacheRepository(redisClient *redis.Client) *CacheRepository {
	return &CacheRepository{redisClient}
}

/* Sets the key-value pair in the cache and sets the lifespan to `expiration` */
func (r *CacheRepository) CacheData(key string, value string, expiration time.Duration) domain.CodedError {
	status := r.cacheClient.Set(key, value, expiration)
	if status.Err() != nil {
		return domain.NewError(status.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

/* Checks if the key exists in the cache */
func (r *CacheRepository) IsCached(key string) bool {
	status := r.cacheClient.Exists(key)
	if status.Err() != nil {
		return false
	}

	return status.Val() == 1
}

/* Retrieves the cached data for a given key */
func (r *CacheRepository) GetCacheData(key string) (string, domain.CodedError) {
	status := r.cacheClient.Get(key)
	if status.Err() != nil {
		return "", domain.NewError(status.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	return status.Val(), nil
}