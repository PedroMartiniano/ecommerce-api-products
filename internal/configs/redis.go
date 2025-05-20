package configs

import (
	"github.com/redis/go-redis/v9"
)

func initRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     GetEnv("REDIS_URL"),
		Password: GetEnv("REDIS_PASSWORD"),
		DB:       0,
	})

	return client
}
