package db_adapters

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(host, port, password string) (redisClient *redis.Client, err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
	})
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return
}
