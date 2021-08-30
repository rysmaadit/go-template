package service

import (
	"github.com/rysmaadit/go-template/app"
	"github.com/rysmaadit/go-template/external/jwt_client"
	"github.com/rysmaadit/go-template/external/redis"
)

type Dependencies struct {
	AuthService  AuthServiceInterface
	CheckService CheckService
}

func InstantiateDependencies(application *app.Application) Dependencies {
	jwtWrapper := jwt_client.New()
	authService := NewAuthService(application.Config, jwtWrapper)
	redisClient := redis.NewRedisClient(application.Config.RedisAddress)
	checkService := NewCheckService(redisClient)

	return Dependencies{
		AuthService:  authService,
		CheckService: checkService,
	}
}
