package infrastructure

import (
	"context"
	"time"
	"log"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) interfaces.RedisCache {
	return &RedisCache{
		client: client,
	}
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", redis.Nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	prefixedKey := "blogs:" + key
	err := r.client.Set(ctx, prefixedKey, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	prefixedKey := "blogs:" + key
	err := r.client.Del(ctx, prefixedKey).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisCache) InvalidateAllBlogCaches(ctx context.Context) error {
	var cursor uint64

	// Log the start of the cache invalidation process
	log.Println("Starting cache invalidation for blog caches")

	for {
		// Scan for keys with the "blogs:" prefix
		keys, nextCursor, err := r.client.Scan(ctx, cursor, "blogs:*", 100).Result()
		if err != nil {
			// Log and return any errors encountered during scanning
			log.Printf("Error scanning Redis: %v", err)
			return err
		}

		// If keys are found, delete them
		if len(keys) > 0 {
			_, err = r.client.Del(ctx, keys...).Result()
			if err != nil {
				// Log and return any errors encountered during deletion
				log.Printf("Error deleting keys from Redis: %v", err)
				return err
			}
			
			// Log the number of keys deleted
			log.Printf("Deleted %d keys from Redis", len(keys))
		}

		// Move to the next cursor to continue scanning
		cursor = nextCursor
		if cursor == 0 {
			// Exit the loop when cursor is 0, indicating the end of iteration
			break
		}
	}

	// Log the completion of the cache invalidation process
	log.Println("Completed cache invalidation for blog caches")

	return nil
}
