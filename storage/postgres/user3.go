package postgres

import (
	"auth/model"
	"database/sql"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) ResetPass(in *model.ResetPassReq) (*model.ResetPassResp, error) {
	_, err := u.DB.Exec(`
			UPDATE users
			SET password = $1
			WHERE email=$2`, in.Password, in.Email)
	if err != nil {
		return nil, err
	}

	return &model.ResetPassResp{
		Message: "assword reset successfully",
	}, nil
}

func (u *UserRepo) ChangePass(in *model.ChangePassReq) (model.ChangePassResp, error) {
	_, err := u.DB.Exec(`
	    UPDATE users
		SET password = $1
		WHERE password = $2`, in.NewPassword, in.NowPassword)
	if err != nil {
		return model.ChangePassResp{}, err
	}

	return model.ChangePassResp{
		Message: "Password changed successfully",
	}, nil
}
