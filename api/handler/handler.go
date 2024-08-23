package handler

import (
	"auth/service"
	"log/slog"
)

type MainHandler interface {
}

type handlerImpl struct {
	Service service.AuthenticateService
	Logger  *slog.Logger
}

func NewMainHandler(service service.AuthenticateService, logger *slog.Logger) MainHandler {
	return &handlerImpl{
		Service: service,
		Logger:  logger,
	}
}
