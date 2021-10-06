package app

import (
	"github.com/rysmaadit/go-template/external/jwt_client"
	"github.com/rysmaadit/go-template/service"
)

type Dependencies struct {
	AuthService service.AuthServiceInterface
}

func InstantiateDependencies(application *Application) Dependencies {
	jwtWrapper := jwt_client.New()
	authService := service.NewAuthService(application.Config, jwtWrapper)

	return Dependencies{
		AuthService: authService,
	}
}
