package infrastructures

import (
	"context"
	"time"

	"aait.backend.g10/domain"
	"github.com/go-redis/redis/v8"
)

type CacheRepo struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewCacheRepo(client *redis.Client, ctx context.Context) *CacheRepo {
	return &CacheRepo{
		Client: client,
		Ctx:    ctx,
	}
}

func (c *CacheRepo) Set(key string, value string, expiry_time time.Duration) *domain.CustomError {
	err := c.Client.Set(c.Ctx, key, value, expiry_time).Err()
	if err != nil {
		return domain.ErrCacheSetFailed
	}
	return nil
}

func (c *CacheRepo) Get(key string) (string, *domain.CustomError) {
	val, err := c.Client.Get(c.Ctx, key).Result()
	if err != nil {
		return "", domain.ErrCacheNotFound
	}
	return val, nil
}

func (c *CacheRepo) Delete(key string) *domain.CustomError {
	err := c.Client.Del(c.Ctx, key).Err()
	if err != nil {
		return domain.ErrCacheDeleteFailed
	}
	return nil
}

func (c *CacheRepo) Increment(key string) error {
	_, err := c.Client.Incr(c.Ctx, key).Result()
	if err != nil {
		return domain.ErrCacheIncrementFailed
	}
	return nil
}

func (c *CacheRepo) Decrement(key string) *domain.CustomError {
	_, err := c.Client.Decr(c.Ctx, key).Result()
	if err != nil {
		return domain.ErrCacheDecrementFailed
	}
	return nil
}
