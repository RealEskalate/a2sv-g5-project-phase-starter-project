package infrastructure

import (
	"Blog_Starter/utils"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type cacheServic struct {
	client *redis.Client
	ctx    context.Context
}

func NewcacheServic(addr string, password string, db int) utils.Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &cacheServic{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *cacheServic) Get(key string) (interface{}, bool) {
	result, err := r.client.Get(key).Result()
	if err == redis.Nil {
		return nil, false // Cache miss
	} else if err != nil {
		fmt.Println("some error happend in the caching database: ", err)
		return nil, false
	}
	return result, true // Cache hit
}

func (r *cacheServic) Set(key string, value interface{}, expiration time.Duration) {
	r.client.Set(key, value, expiration).Err()

}

func (r *cacheServic) Delete(key string) {

	err := r.client.Del(key).Err()
	if err != nil {
		fmt.Println("some error happend in the caching database: ", err)
	}

}
