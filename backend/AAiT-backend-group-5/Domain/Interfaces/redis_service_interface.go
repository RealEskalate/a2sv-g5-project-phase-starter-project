package interfaces

import (
	"context"
	"time"
)

type RedisCache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
	InvalidateAllBlogCaches(ctx context.Context) error
}
