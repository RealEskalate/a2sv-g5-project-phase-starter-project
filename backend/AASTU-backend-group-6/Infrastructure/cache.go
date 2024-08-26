package infrastructure

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

// CacheService is a struct that holds the redis client
type CacheService struct {
	client *redis.Client
}

// NewCacheService is a function that initializes a Redis client and returns a new CacheService
func NewCacheService() *CacheService {
	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address on localhost
		Password: "",                // No password set for local Redis
		DB:       0,                 // Use default DB
	})

	// Test the connection with a ping
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	return &CacheService{
		client: client,
	}
}
