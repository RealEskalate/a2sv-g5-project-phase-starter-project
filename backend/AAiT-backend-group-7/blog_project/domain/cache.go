// domain/cache.go

package domain

import (
    "context"
    "time"
)

type Cache interface {
    Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
    Get(ctx context.Context, key string, dest interface{}) error
    Del(ctx context.Context, key string) error
    DelByPattern(ctx context.Context, pattern string) error
}