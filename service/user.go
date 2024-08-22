package service

import (
	pb "auth/generated/users"
	"auth/storage"
	"log/slog"
)


type MainService interface{

}

type Service struct{
	pb.UnimplementedUserServiceServer
	Storage *storage.IStorage
	Logger *slog.Logger
}

func NewService(storage *storage.IStorage, logger *slog.Logger)*Service{
	return &Service{
		Storage: storage,
		Logger: logger,
	}
}