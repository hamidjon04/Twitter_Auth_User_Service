package service

import (
	pb "auth/generated/users"
	"auth/storage"
	"context"
	"fmt"
	"log/slog"
)

type MainService interface {
	GetUsers(context.Context, *pb.GetUserReq) (*pb.GetUserRes, error)
	DeleteUsers(context.Context, *pb.Id) (*pb.Massage, error)
	GetByIdUsers(context.Context, *pb.Id) (*pb.User, error)
	AddFollower(context.Context, *pb.FollowerReq) (*pb.Massage, error)
	GetFollowers(context.Context, *pb.GetFollowersReq) (*pb.GetaFollowersRes, error)
	DeleteFollower(context.Context, *pb.DeleteFollowers) (*pb.Massage, error)
	GetByIdFollower(context.Context, *pb.Id) (*pb.Followers, error)
	AddFollowing(context.Context, *pb.FollowingReq) (*pb.Massage, error)
	GetFollowing(context.Context, *pb.GetFollowingReq) (*pb.GetaFollowingRes, error)
	DeleteFollowing(context.Context, *pb.DeleteFollowings) (*pb.Massage, error)
	GetByIdFollowing(context.Context, *pb.Id) (*pb.Following, error)
}

type Service struct {
	pb.UnimplementedUserServiceServer
	Storage storage.IStorage
	Logger  *slog.Logger
}

func NewService(storage storage.IStorage, logger *slog.Logger) *Service {
	return &Service{
		Storage: storage,
		Logger:  logger,
	}
}

func (s *Service) GetUsers(ctx context.Context, request *pb.GetUserReq) (*pb.GetUserRes, error) {
	resp, err := s.Storage.UserRepo().GetUser(request)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Error on get user: %v", err))
		return resp, nil
	}

	return resp, err
}
