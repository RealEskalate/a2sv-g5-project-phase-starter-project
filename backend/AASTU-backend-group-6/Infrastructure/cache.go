package infrastructure

import (
	"context"
	"log"
	"time"

	domain "blogs/Domain"
	
	"github.com/redis/go-redis/v9"
)

// CacheService is a struct that holds the redis client

// NewCacheService is a function that initializes a Redis client and returns a new CacheService
func NewCacheService(env Config) *domain.CacheService {
	// Create a new Redis client
	var redisUrl string
	if env.RedisURL != "" {
		redisUrl= env.RedisURL
	}else {
		redisUrl= "localhost:6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr:     redisUrl,         // Redis server address
		Password: "",               // No password set for local Redis
		DB:       0,                // Use default DB
	})

	// Test the connection with a ping
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	return &domain.CacheService{
		Client: client,
	}
}
