package api

import (
	"auth/config"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	StartRouter(cfg config.Config) error
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

func (r *controllerImpl) SetUpRouter() {

}
