package redis_service

import "github.com/go-redis/redis"

func ConnectStore(connectionString string) (*redis.Client, error) {
	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	return client, nil
}
