package interfaces

import "aait.backend.g10/domain"

type CacheRepoInterface interface {
	Set(key string, value string) *domain.CustomError
	Get(key string) (string, *domain.CustomError)
	Delete(key string) *domain.CustomError
	Increment(key string) *domain.CustomError
	Decrement(key string) *domain.CustomError
}
