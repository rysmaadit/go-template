package service

import (
	"fmt"
	"github.com/rysmaadit/go-template/external/redis"
	log "github.com/sirupsen/logrus"
)

type checkService struct {
	redisClient redis.Client
}

type CheckService interface {
	CheckRedis() ([]byte, error)
}

func NewCheckService(redisClient redis.Client) *checkService {
	return &checkService{
		redisClient: redisClient,
	}
}

func (c *checkService) CheckRedis() ([]byte, error) {
	err := c.redisClient.Ping()
	if err != nil {
		log.Warning(fmt.Errorf("redis ping failed: %v", err))
		return nil, err
	}
	return []byte("Success"), err
}
