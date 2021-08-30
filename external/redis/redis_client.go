package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type Client interface {
	Get(key string) ([]byte, error)
	Ping() error
}

type client struct {
	redisClient *redis.Client
}

func (c *client) Get(key string) ([]byte, error) {
	ctx := context.Background()
	respByte, err := c.redisClient.Get(ctx, key).Bytes()
	if err != nil && err != redis.Nil {
		log.Warning(err)
	}
	return respByte, err
}

func (c *client) Ping() error {
	ctx := context.Background()
	return c.redisClient.Ping(ctx).Err()
}

func NewRedisClient(address string) *client {
	redisClient := redis.NewClient(
		&redis.Options{
			Addr: address,
			DB:   0,
		},
	)
	return &client{
		redisClient: redisClient,
	}
}
