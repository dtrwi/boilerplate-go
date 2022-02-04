package driver

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// RedisBOption struct
type RedisOption struct {
	URL      string
	Password string
}

// NewRedis return a client connection handle to a Redis-like server.
func NewRedis(option RedisOption) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     option.URL,
		Password: option.Password,
		DB:       0,
	})

	_, err = client.Ping(context.Background()).Result()

	return
}
