package service

import (
	"auth/model"
	"auth/storage/postgres"
	"log/slog"
)

type AuthRepo struct {
	Auth   *postgres.UserRepo
	Logger *slog.Logger
}

func NewAuthRepo(auth *postgres.UserRepo, logger *slog.Logger) *AuthRepo {
	return &AuthRepo{Auth: auth, Logger: logger}
}

func (a *AuthRepo) ResetPass(in *model.ResetPassReq) (*model.ResetPassResp, error) {
	resp, err := a.Auth.ResetPass(in)
	if err != nil {
		a.Logger.Error("error resetting password", err)
		return nil, err
	}
	return resp, nil
}

func (a *AuthRepo) ChangePass(in *model.ChangePassReq) (model.ChangePassResp, error) {
	resp, err := a.Auth.ChangePass(in)
	if err != nil {
		a.Logger.Error("error changing password", err)
		return model.ChangePassResp{}, err
	}
	return resp, nil
}
