package router

import (
	"context"
	"net/http"
	"os"

	"github.com/rysmaadit/go-template/constant"
	"github.com/sirupsen/logrus"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"

	"github.com/rysmaadit/go-template/app"
	"github.com/rysmaadit/go-template/handler"
	"github.com/rysmaadit/go-template/middleware"
	"github.com/rysmaadit/go-template/service"
)

func NewRouter(dependencies app.Dependencies) http.Handler {
	r := mux.NewRouter()

	setAuthRouter(r, dependencies.AuthService)
	ctx := context.Background()

	//loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	r.Use(middleware.LoggingMiddleware(logger, ctx))

	logrus.Info(ctx.Value(constant.RequestID))
	return r
}

func setAuthRouter(router *mux.Router, dependencies service.AuthServiceInterface) {
	router.Methods(http.MethodGet).Path("/auth/token").Handler(handler.GetToken(dependencies))
	router.Methods(http.MethodPost).Path("/auth/token/validate").Handler(handler.ValidateToken(dependencies))
}
