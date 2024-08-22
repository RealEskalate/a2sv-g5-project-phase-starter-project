package utils

import "time"

type Cache interface {
	Get(key string) (value interface{}, found bool)
	Set(key string, value interface{}, expiration time.Duration)
	Delete(key string)
}
