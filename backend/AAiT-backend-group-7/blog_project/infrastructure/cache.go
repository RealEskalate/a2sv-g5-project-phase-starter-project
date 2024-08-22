package infrastructure

import (
    "context"
    "encoding/json"
    "time"

    "github.com/go-redis/redis/v8"
    "blog_project/domain"
)

type RedisCache struct {
    client *redis.Client
}

func NewRedisCache(addr string) domain.Cache {
    return &RedisCache{
        client: redis.NewClient(&redis.Options{
            Addr: addr,
        }),
    }
}

func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
    json, err := json.Marshal(value)
    if err != nil {
        return err
    }
    return c.client.Set(ctx, key, json, expiration).Err()
}

func (c *RedisCache) Get(ctx context.Context, key string, dest interface{}) error {
    val, err := c.client.Get(ctx, key).Result()
    if err != nil {
        return err
    }
    return json.Unmarshal([]byte(val), dest)
}

func (c *RedisCache) Del(ctx context.Context, key string) error {
    return c.client.Del(ctx, key).Err()
}

func (c *RedisCache) DelByPattern(ctx context.Context, pattern string) error {
    keys, err := c.client.Keys(ctx, pattern).Result()
    if err != nil {
        return err
    }
    if len(keys) > 0 {
        return c.client.Del(ctx, keys...).Err()
    }
    return nil
}