package infrastructure

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheService interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
	Increment(key string) error
	Decrement(key string) error
}

type cacheService struct {
	client *redis.Client
	ctx    context.Context
}

// NewcacheService creates a new cacheService instance.
func NewCacheService(redisAddr, redisPassword string, redisDB int) CacheService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	return &cacheService{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (cs *cacheService) Increment(key string) error {
    return cs.client.Incr(cs.ctx, key).Err()
}

func (cs *cacheService) Decrement(key string) error {
    return cs.client.Decr(cs.ctx, key).Err()
}


// Set sets a value in the cache with an optional expiration time.
func (cs *cacheService) Set(key string, value interface{}, expiration time.Duration) error {
	err := cs.client.Set(cs.ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves a value from the cache.
func (cs *cacheService) Get(key string) (string, error) {
	val, err := cs.client.Get(cs.ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Key does not exist
	} else if err != nil {
		return "", err
	}
	return val, nil
}

// Delete removes a value from the cache.
func (cs *cacheService) Delete(key string) error {
	err := cs.client.Del(cs.ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
