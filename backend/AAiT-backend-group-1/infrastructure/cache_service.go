package infrastructure

import (
    "context"
    "time"

    "github.com/go-redis/redis/v8"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
)

type cacheService struct {
    client *redis.Client
    ctx    context.Context
}

// NewCacheService creates a new CacheService instance.
func NewCacheService(redisAddr string, redisPassword string, redisDB int) domain.CacheService {
    rdb := redis.NewClient(&redis.Options{
        Addr:     redisAddr,
        Password: redisPassword, // no password set
        DB:       redisDB,       // use default DB
    })

    return &cacheService{
        client: rdb,
        ctx:    context.Background(),
    }
}

// Set sets a value in the cache with an optional expiration time.
func (cs *cacheService) Set(key string, value string, expiration time.Duration) error {
    return cs.client.Set(cs.ctx, key, value, expiration).Err()
}

// Get retrieves a value from the cache.
func (cs *cacheService) Get(key string) (string, error) {
    val, err := cs.client.Get(cs.ctx, key).Result()
    if err == redis.Nil {
        return "", nil // Key does not exist
    } else if err != nil {
        return "", err
    }
    return val, nil
}

// Delete removes a value from the cache.
func (cs *cacheService) Delete(key string) error {
    return cs.client.Del(cs.ctx, key).Err()
}
