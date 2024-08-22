package interfaces

import (
	"time"

	"aait.backend.g10/domain"
)

type CacheRepoInterface interface {
	Set(key string, value string, expiry_time time.Duration) *domain.CustomError
	Get(key string) (string, *domain.CustomError)
	Delete(key string) *domain.CustomError
	Increment(key string) *domain.CustomError
	Decrement(key string) *domain.CustomError
}
