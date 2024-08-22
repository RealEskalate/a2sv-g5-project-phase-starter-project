package bootstrap

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(env *Env) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     env.RedisHost + ":" + env.RedisPort,
		Password: env.RedisPassword,
		DB:       0,
	})

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	return client
}
