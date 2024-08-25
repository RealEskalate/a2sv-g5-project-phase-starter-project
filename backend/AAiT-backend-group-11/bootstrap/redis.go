
package bootstrap

import (
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(env *Env) *redis.Client{
	redisClient := redis.NewClient(&redis.Options{
		Addr: env.RedisAddress, 
	})

	return redisClient

}

func CloseRedisConnection(redisClient *redis.Client){
	err := redisClient.Close()
	if err != nil {
		panic(err)
	}
}