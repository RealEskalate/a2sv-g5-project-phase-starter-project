package infrastructure

import (
	"context"
	"time"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) interfaces.RedisCache {
	return &RedisCache{
		client: client,
	}
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", redis.Nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	prefixedKey := "blogs:" + key
	err := r.client.Set(ctx, prefixedKey, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	prefixedKey := "blogs:" + key
	err := r.client.Del(ctx, prefixedKey).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisCache) InvalidateAllBlogCaches(ctx context.Context) error {
	var cursor uint64

	for {
		keys, nextCursor, err := r.client.Scan(ctx, cursor, "blogs:*", 100).Result()
		if err != nil {
			return err
		}

		if len(keys) > 0 {
			_, err = r.client.Del(ctx, keys...).Result()
			if err != nil {
				return err
			}

		}

		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}

	return nil
}
