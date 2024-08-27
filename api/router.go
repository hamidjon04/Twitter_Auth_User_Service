package api

import (
	"auth/api/handler"
	"auth/config"
	"auth/service"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	StartRouter(cfg config.Config) error
	SetUpRouter(auth service.AuthenticateService, users service.Service)
}

type controllerImpl struct {
	Router *gin.Engine
	Logger *slog.Logger
}

func NewController(router *gin.Engine, logger *slog.Logger) Controller {
	return &controllerImpl{
		Router: router,
		Logger: logger,
	}
}

func (r *controllerImpl) StartRouter(cfg config.Config) error {
	return r.Router.Run(cfg.USER_ROUTER)
}

func (r *controllerImpl) SetUpRouter(auth service.AuthenticateService, users service.Service) {
	h := handler.NewMainHandler(auth, r.Logger, &users)

	user := r.Router.Group("/user")
	user.POST("/register", h.Register)
	user.POST("/login", h.Login)
	user.POST("/logout", h.Logout)
}
