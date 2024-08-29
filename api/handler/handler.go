package handler

import (
	"auth/service"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type MainHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	ResetPassword(c *gin.Context)

	ChangePassword(c *gin.Context)
	GetUsers(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetByIdUsers(c *gin.Context)
	Subscribe(c *gin.Context)
	GetFollowers(c *gin.Context)
	DeleteFollower(c *gin.Context)
	GetByIdFollower(c *gin.Context)
	GetFollowing(c *gin.Context)
	DeleteFollowing(c *gin.Context)
	GetByIdFollowing(c *gin.Context)
}

type handlerImpl struct {
	Service service.AuthenticateService
	Logger  *slog.Logger
	UserService service.MainService

}

func NewMainHandler(service service.AuthenticateService, logger *slog.Logger, user service.MainService) MainHandler {
	return &handlerImpl{
		Service: service,
		Logger:  logger,
		UserService: user,
	}
}
