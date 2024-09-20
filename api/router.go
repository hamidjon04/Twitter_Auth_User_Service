package api

import (
	"auth/api/handler"
	"auth/config"
	"auth/service"
	"log/slog"
	_ "auth/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

// @title AuthService API
// @version 1.0
// @description Auth service
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http
func (r *controllerImpl) StartRouter(cfg config.Config) error {
	return r.Router.Run(cfg.USER_ROUTER)
}

func (r *controllerImpl) SetUpRouter(authS service.AuthenticateService, users service.Service) {
	h := handler.NewMainHandler(authS, r.Logger, &users)

	r.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	auth := r.Router.Group("/auth")
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)
	auth.POST("/logout", h.Logout)
	auth.POST("/reset-password", h.ResetPassword)
}
