package handler

import (
	"log/slog"
)

type MainHandler interface{
	
}

type handlerImpl struct{
	Logger *slog.Logger
}

func NewMainHandler(logger *slog.Logger)MainHandler{
	return &handlerImpl{
		Logger: logger,
	}
}

