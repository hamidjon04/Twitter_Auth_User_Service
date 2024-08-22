package postgres

import (
	"auth/model"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

type UserRepo interface {
	Register(req *model.RegisterReq) (*model.RegisterResp, error)
}

type userImpl struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userImpl{
		DB: db,
	}
}

func (U *userImpl) Register(req *model.RegisterReq) (*model.RegisterResp, error) {
	id := uuid.NewString()
	query := `
				INSERT INTO users(
					id, email, password)
				VALUES
					($1, $2, $3)`
	_, err := U.DB.Exec(query, id, req.Email, req.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &model.RegisterResp{
		Id:      id,
		Message: "Tizimdan muvvaffaqiyatli ro'yxatdan o'tdingiz",
	}, nil
}

func (U *userImpl) GetUserByEmail(email string) (*model.UserInfo, error) {
	resp := model.UserInfo{}
	query := `
				SELECT 
					id, password, role
				FROM
					users
				WHERE
					email = $1`
	err := U.DB.QueryRow(query, email).
		Scan(&resp.Id, &resp.Password, &resp.Role)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &resp, nil
}

func (U *userImpl) SaveRefreshToken(req *model.SaveToken) error {
	query := `	
				INSERT INTO users(
					user_id, token, expires_at)
				VALUES
					($1, $2, $3)`
	_, err := U.DB.Query(query, req.UserId, req.RefreshToken, req.ExpiresAt)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
