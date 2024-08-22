package domain

// CacheRepository is an interface for cache repository

type Cache interface {
	// GetCache is a function to get cache
	GetCache(key string) (string, error)
	// SetCache is a function to set cache
	SetCache(key string, value string) error
	// DeleteCache is a function to delete cache
	DeleteCache(key string) error
}
