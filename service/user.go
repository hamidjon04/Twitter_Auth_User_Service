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
	Subscribe(ctx context.Context, req *pb.FollowingReq) (*pb.Massage, error)
	GetFollowers(context.Context, *pb.GetFollowersReq) (*pb.GetaFollowersRes, error)
	DeleteFollower(context.Context, *pb.DeleteFollowerReq) (*pb.Massage, error)
	GetByIdFollower(context.Context, *pb.DeleteFollowerReq) (*pb.Follow, error)
	GetFollowing(context.Context, *pb.GetFollowingReq) (*pb.GetaFollowingRes, error)
	DeleteFollowing(context.Context, *pb.DeleteFollowerReq) (*pb.Massage, error)
	GetByIdFollowing(context.Context, *pb.DeleteFollowerReq) (*pb.Follow, error)
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
 
func (s *Service) DeleteUsers(ctx context.Context, req *pb.Id) (*pb.Massage, error) {
	resp, err := s.Storage.UserRepo().DeleteUsers(req)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Error on delete user: %v", err))
		return resp, err
	}
	return resp, err
}

func (s *Service) GetByIdUsers(ctx context.Context, req *pb.Id) (*pb.User, error) {
	resp, err := s.Storage.UserRepo().GetByIdUsers(req)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Error on get by id user: %v", err))
		return resp, err
	}

	return resp, nil
}

func (s *Service) GetFollowers(ctx context.Context, req *pb.GetFollowersReq) (*pb.GetaFollowersRes, error) {
	resp, err := s.Storage.FollowRepository().GetFollowers(req)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Error on get followers: %v", err))
		return resp, err
	}

	return resp, err
}

func (s *Service) DeleteFollower(ctx context.Context, req *pb.DeleteFollowerReq) (*pb.Massage, error) {
	resp, err := s.Storage.FollowRepository().DeleteFollower(req)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Error follower ochirishda: %v", err))
		return resp, err
	}

	return resp, err
}

func (s *Service) GetByIdFollower(ctx context.Context, req *pb.DeleteFollowerReq) (*pb.Follow, error) {
	resp, err := s.Storage.FollowRepository().GetByIdFollowers(req)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Error followerni idsi bo'yicha olishda: %v", err))
		return resp, err
	}

	return resp, nil
}

func (s *Service) Subscribe(ctx context.Context, req *pb.FollowingReq) (*pb.Massage, error) {
	resp, err := s.Storage.FollowRepository().Subscribe(req)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Error following bo'lishda: %v", err))
		return resp, err
	}

	return resp, err
}

func (s *Service) GetFollowing(ctx context.Context, req *pb.GetFollowingReq) (*pb.GetaFollowingRes, error) {
	resp, err := s.Storage.FollowRepository().GetFollowing(req)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Error userni following bo'lganlarini olishda: %v", err))
		return resp, nil
	}

	return resp, err
}


func (s *Service) DeleteFollowing(ctx context.Context, req *pb.DeleteFollowerReq) (*pb.Massage, error) {
	resp, err := s.Storage.FollowRepository().DeleteFollowing(req)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Error followingni o'chirishda: %v", err))
		return resp, err
	}

	return resp, err
}

func (s *Service) GetByIdFollowing(ctx context.Context, req *pb.DeleteFollowerReq) (*pb.Follow, error) {
	resp, err := s.Storage.FollowRepository().GetByIdFollowing(req)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Error followingni id boyicha olishda: %v", err))
		return resp, err
	}

	return resp, err
}

