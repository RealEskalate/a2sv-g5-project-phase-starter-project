package config

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache() *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &RedisCache{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (r *RedisCache) GetCache(key string) (string, error) {

	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisCache) SetCache(key string, value string) error {

	err := r.client.Set(r.ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisCache) DeleteCache(key string) error {

	err := r.client.Del(r.ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
