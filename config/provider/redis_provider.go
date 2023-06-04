package provider

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisProvider(env *EnvProvider) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     env.redisUrl,
		Password: env.redisPassword,
		DB:       0, // use the default DB
	})

	return rdb
}
