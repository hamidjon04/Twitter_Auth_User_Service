package service

import (
	"auth/model"
	"auth/storage"
	"fmt"
	"log/slog"
)

type AuthenticateService interface {
	RegisterUser(request *model.RegisterReq) (*model.RegisterResp, error) 
	LoginUser(request *model.LoginReq) (*model.UserInfo, error) 
	ResetPassword(request *model.ResetPassReq) (*model.ResetPassResp, error)
	ChangePassword(request *model.ChangePassReq) (*model.ChangePassResp, error) 
	SaveRefreshToken(request *model.SaveToken) (*model.SuccessResponse, error) 
	InvalidateRefreshToken(tokenString string) (*model.SuccessResponse, error) 
	IsRefreshTokenValid(tokenString string) (bool, error) 
	GetUserByEmail(email string)(*model.UserInfo, error)
}

type authenticateServiceImpl struct {
	storage storage.IStorage
	logger  *slog.Logger
}

func NewAuthenticateService(storage storage.IStorage, logger *slog.Logger) AuthenticateService {
	return &authenticateServiceImpl{
		storage: storage,
		logger:  logger,
	}
}

func (s *authenticateServiceImpl) RegisterUser(request *model.RegisterReq) (*model.RegisterResp, error) {
	resp, err := s.storage.UserRepo().Register(request)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error in user register: %v", err))
		return resp, err
	}

	return resp, err
}

func (s *authenticateServiceImpl) LoginUser(request *model.LoginReq) (*model.UserInfo, error) {
	resp, err := s.storage.UserRepo().GetUserByEmail(request.Email)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error on user login %v", err))
		return resp, err
	}

	return resp, err
}

func (s *authenticateServiceImpl) ResetPassword(request *model.ResetPassReq) (*model.ResetPassResp, error) {
	resp, err := s.storage.UserRepo().ResetPass(request)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error on reset password: %v", err))
		return resp, err
	}

	return resp, err
}

func (s *authenticateServiceImpl) ChangePassword(request *model.ChangePassReq) (*model.ChangePassResp, error) {
	resp, err := s.storage.UserRepo().ChangePass(request)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error on user of change password: %v", err))
		return &resp, err
	}

	return &resp, err
}

func (s *authenticateServiceImpl) SaveRefreshToken(request *model.SaveToken) (*model.SuccessResponse, error) {
	err := s.storage.UserRepo().SaveRefreshToken(request)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error on token save for user: %v", err))
		return &model.SuccessResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &model.SuccessResponse{
		Success: true,
		Message: "token saved for user successfully",
	}, err
}

func (s *authenticateServiceImpl) InvalidateRefreshToken(tokenString string) (*model.SuccessResponse, error) {
	err := s.storage.UserRepo().InvalidateRefreshToken(tokenString)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error on token invalidation for the user: %v", err))
		return &model.SuccessResponse{
			Message: err.Error(),
			Success: false,
		}, err
	}

	return &model.SuccessResponse{
		Success: true,
		Message: "Token invalidate for the user successfully",
	}, err
}

func (s *authenticateServiceImpl) IsRefreshTokenValid(tokenString string) (bool, error) {
	isValid, err := s.storage.UserRepo().IsRefreshTokenValid(tokenString)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error on token valid token for user: %v", err))
		return isValid, err
	}

	return isValid, err
}

func (s *authenticateServiceImpl) GetUserByEmail(email string)(*model.UserInfo, error){
	resp, err := s.storage.UserRepo().GetUserByEmail(email)
	if err != nil{
		s.logger.Error(fmt.Sprintf("Datavazadan ma'lumotlani olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}
