package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)


type Client interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context) error
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)
	HSet(ctx context.Context, key string, field string, value interface{}) error
	HGet(ctx context.Context, key string, field string) (interface{}, error)
	HGetAll(ctx context.Context, key string) (map[string]interface{}, error)
	HDel(ctx context.Context, key string, fields ...string) error
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Exists(ctx context.Context, key string) (bool, error)
	Del(ctx context.Context, keys ...string) error
	Keys(ctx context.Context, pattern string) ([]string, error)
	FlushAll(ctx context.Context) error
	Err() error
	Close() error
}

type RedisClient struct {
	cl *redis.Client
}

func NewClient(addr string) Client {
	return &RedisClient{
		cl: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

func (r *RedisClient) Connect(ctx context.Context) error {
	return r.cl.Ping(ctx).Err()
}

func (r *RedisClient) Disconnect(ctx context.Context) error {
	return r.cl.Close()
}

func (r *RedisClient) Ping(ctx context.Context) error {
	return r.cl.Ping(ctx).Err()
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.cl.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string) (interface{}, error) {
	return r.cl.Get(ctx, key).Result()
}

func (r *RedisClient) HSet(ctx context.Context, key string, field string, value interface{}) error {
	return r.cl.HSet(ctx, key, field, value).Err()
}

func (r *RedisClient) HGet(ctx context.Context, key string, field string) (interface{}, error) {
	return r.cl.HGet(ctx, key, field).Result()
}

func (r *RedisClient) HGetAll(ctx context.Context, key string) (map[string]interface{}, error) {
	result, err := r.cl.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return convertToStringInterfaceMap(result), nil
}

func convertToStringInterfaceMap(input map[string]string) map[string]interface{} {
	output := make(map[string]interface{}, len(input))
	for key, value := range input {
		output[key] = value
	}
	return output
}

func (r *RedisClient) HDel(ctx context.Context, key string, fields ...string) error {
	return r.cl.HDel(ctx, key, fields...).Err()
}

func (r *RedisClient) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.cl.Expire(ctx, key, expiration).Err()
}

func (r *RedisClient) Exists(ctx context.Context, key string) (bool, error) {
	exists, err := r.cl.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (r *RedisClient) Del(ctx context.Context, keys ...string) error {
	return r.cl.Del(ctx, keys...).Err()
}

func (r *RedisClient) Keys(ctx context.Context, pattern string) ([]string, error) {
	return r.cl.Keys(ctx, pattern).Result()
}

func (r *RedisClient) FlushAll(ctx context.Context) error {
	return r.cl.FlushAll(ctx).Err()
}

func (r *RedisClient) Close() error {
	return r.cl.Close()
}

func (r *RedisClient) Err() error {
	return r.cl.Context().Err()
}
