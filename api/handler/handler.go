package handler

import (
	"auth/service"
	"log/slog"
)

type MainHandler interface{
	
}

type handlerImpl struct{
	Logger *slog.Logger
	Auth service.AuthRepo
}

func NewMainHandler(logger *slog.Logger, auth service.AuthRepo)MainHandler{
	return &handlerImpl{
		Logger: logger,
		Auth: auth,
	}
}

