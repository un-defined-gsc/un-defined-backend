package store

import (
	"strconv"

	"github.com/gofiber/storage/redis/v3"
)

func NewFiberStore(host, password, port string) *redis.Storage {
	p, err := strconv.Atoi(port)
	if err != nil {
		p = 6379
	}
	return redis.New(redis.Config{
		Host:     host,
		Port:     p,
		Password: password,
	})

}
