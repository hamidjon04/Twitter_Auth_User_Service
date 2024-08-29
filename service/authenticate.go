package service

import (
	"auth/model"
	"auth/storage"
	"context"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"strconv"
	"time"
)

type AuthenticateService interface {
	RegisterUser(request *model.RegisterReq) (*model.RegisterResp, error)
	ResetPassword(request *model.ResetPassReq) (*model.ResetPassResp, error)
	ChangePassword(request *model.ChangePassReq) (*model.ChangePassResp, error)
	SaveRefreshToken(request *model.SaveToken) error
	InvalidateRefreshToken(userId string) error
	IsRefreshTokenValid(tokenString string) (bool, error)
	GetUserByEmail(email string) (*model.UserInfo, error)
	AddTokenBlacklisted(ctx context.Context, token string, expirationTime time.Duration) error
	IsTokenBlacklisted(ctx context.Context, token string) (bool, error)
	ForgotPassword(ctx context.Context, req *model.ForgotPassReq) (*model.ForgotPassResp, error)
	IsCodeValid(ctx context.Context, email, code string) (bool, error)
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

func (s *authenticateServiceImpl) ForgotPassword(ctx context.Context, req *model.ForgotPassReq) (*model.ForgotPassResp, error) {
	code := rand.IntN(999999) + 100000

	err := s.storage.RedisUserRepo().StoreCode(ctx, req.Email, strconv.Itoa(code), time.Duration(time.Minute*5))
	if err != nil {
		s.logger.Error(fmt.Sprint("Error codeni redisga saqlashda: ", err))
		return &model.ForgotPassResp{
			Message: err.Error(),
		}, err
	}

	return &model.ForgotPassResp{
		Message: fmt.Sprintf("Code %s emailga yuborildi", req.Email),
	}, nil
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

func (s *authenticateServiceImpl) SaveRefreshToken(request *model.SaveToken) error {
	err := s.storage.UserRepo().SaveRefreshToken(request)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error on token save for user: %v", err))
	}
	return err
}

func (s *authenticateServiceImpl) InvalidateRefreshToken(userID string) error {
	err := s.storage.UserRepo().InvalidateRefreshToken(userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error on token invalidation for the user: %v", err))
		return err
	}

	return nil
}

func (s *authenticateServiceImpl) IsRefreshTokenValid(tokenString string) (bool, error) {
	isValid, err := s.storage.UserRepo().IsRefreshTokenValid(tokenString)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error on token valid token for user: %v", err))
		return isValid, err
	}

	return isValid, err
}

func (s *authenticateServiceImpl) GetUserByEmail(email string) (*model.UserInfo, error) {
	resp, err := s.storage.UserRepo().GetUserByEmail(email)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Datavazadan ma'lumotlani olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (s *authenticateServiceImpl) AddTokenBlacklisted(ctx context.Context, token string, expirationTime time.Duration) error {
	err := s.storage.RedisUserRepo().AddTokenBlacklisted(ctx, token, expirationTime)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error tokenni blacklistga solishda: %v", err))
		return err
	}
	return err
}

func (s *authenticateServiceImpl) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	istokenblacklest, err := s.storage.RedisUserRepo().IsTokenBlacklisted(ctx, token)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error token blacklistda borligini tekshirishda: %v", err))
		return istokenblacklest, err
	}

	return istokenblacklest, err
}

func (s *authenticateServiceImpl) IsCodeValid(ctx context.Context, email, code string) (bool, error) {
	resp, err := s.storage.RedisUserRepo().IsCodeValid(ctx, email, code)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error codeni tekshirishda: %v", err))
		return resp, err
	}

	return resp, err
}
