package redis_service

import "github.com/go-redis/redis"

// ConnectStore connects to a redis store and returns a client
func ConnectStore(connectionString string) (*redis.Client, error) {
	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	return client, nil
}

// DisconnectStore disconnects from a redis store
func DisconnectStore(client *redis.Client) error {
	return client.Close()
}
