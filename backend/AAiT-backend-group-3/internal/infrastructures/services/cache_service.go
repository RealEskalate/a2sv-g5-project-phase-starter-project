package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type ICacheService interface {
	SetBlog(key string, value interface{}, expiration time.Duration) error
	GetBlog(key string, result interface{}) error
	Delete(key string) error
	BlacklistTkn(token string, expiration time.Duration) error
	IsTknBlacklisted(token string) (bool, error)
}

type cacheService struct {
	client *redis.Client
	ctx    context.Context
}

func NewCacheService(redisAddr, redisPassword string, redisDB int) ICacheService {
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

func (cs *cacheService) SetBlog(key string, value interface{}, expiration time.Duration) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = cs.client.Set(cs.ctx, key, jsonValue, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cs *cacheService) GetBlog(key string, result interface{}) error {
	val, err := cs.client.Get(cs.ctx, key).Result()
	if err == redis.Nil {
		return nil 
	} else if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), result)
	if err != nil {
		return err
	}
	return nil
}

func (cs *cacheService) BlacklistTkn(token string, expiration time.Duration) error {
    return cs.client.Set(cs.ctx, token, "blacklisted", expiration).Err()
}

func (cs *cacheService) IsTknBlacklisted(token string) (bool, error) {
    result, err := cs.client.Get(cs.ctx, token).Result()
    if err == redis.Nil {
        return false, nil 
    } else if err != nil {
        return false, err
    }
    return result == "blacklisted", nil
}

func (cs *cacheService) Delete(key string) error {
	err := cs.client.Del(cs.ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
