package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	client    *redis.Client
	expiryDay time.Duration
}

var _ ICache = &redisCache{}

func NewRedisCache(client *redis.Client, expiryDay time.Duration) ICache {
	return &redisCache{
		client: client,
		expiryDay: time.Second * expiryDay,
	}
}


func (r *redisCache) Get(key string) (interface{}, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil 
		}
		return nil, err 
	}

	var result interface{}
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, err 
	}

	return result, nil
}


func (r *redisCache) Set(key string, value interface{}) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err 
	}

	err = r.client.Set(context.Background(), key, jsonData, r.expiryDay).Err()
	if err != nil {
		return err 
	}

	return nil
}
