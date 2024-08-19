package repository

import (
	"context"

	"github.com/go-redis/redis"
)

type CacheRepository struct {
	redisClient *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) *CacheRepository {
	return &CacheRepository{redisClient}
}

func (r *CacheRepository) CacheData(c context.Context) {

}
