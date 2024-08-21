package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	icache "github.com/group13/blog/usecase/common/i_cache"
)

type redisCache struct {
	host      string
	port      string
	expiryDay time.Duration
	client    *redis.Client
}

var _ icache.ICache = &redisCache{}

func NewRedisCache(host string, port string, expiryDay time.Duration, db int) icache.ICache {
	return &redisCache{
		host:      host,
		port:      port,
		expiryDay: expiryDay,
		client: redis.NewClient(&redis.Options{
			Addr:     host + ":" + port,
			Password: "",
			DB:       db,
		}),
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
