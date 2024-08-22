package domain




type Cache interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
	Keys(pattern string) ([]string, error)
}