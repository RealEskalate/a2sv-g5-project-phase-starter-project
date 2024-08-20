package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(env Env, ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     env.REDIS_DB_ADDRESS,
		Password: env.REDIS_DB_PASSWORD,
		DB:       env.REDIS_DB,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return client
}
