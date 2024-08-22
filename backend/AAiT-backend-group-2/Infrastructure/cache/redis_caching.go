package cache


import (
	"context"
    "encoding/json"
    "github.com/go-redis/redis/v8"
    "AAiT-backend-group-2/Domain"
    "time"
)

type RedisCache struct {
	client *redis.Client
	ctx context.Context
}


func NewRedisCache(redisAddr string, redisPassword string,db int) domain.Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       db,
	})

	return &RedisCache{
		client: client,
		ctx: context.Background(),
	}
}



func (r *RedisCache) Set(key string, value interface{}) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(r.ctx, key, jsonValue, 1*time.Hour).Err()
}


func (r *RedisCache) Get(key string) (interface{}, error) {
	jsonValue, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var value interface{}
	err = json.Unmarshal([]byte(jsonValue), &value)
	if err != nil {
		return nil, err
	}
	return value, nil
}


func (r *RedisCache) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

func (r *RedisCache) Keys(pattern string) ([]string, error) {
	return r.client.Keys(r.ctx, pattern).Result()
}