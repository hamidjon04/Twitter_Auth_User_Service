package postgres

import (
	pb "auth/generated/users"
	"auth/model"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type UserRepo interface {
	Register(req *model.RegisterReq) (*model.RegisterResp, error)
	GetUserByEmail(email string) (*model.UserInfo, error)
	SaveRefreshToken(req *model.SaveToken) error
	ResetPass(in *model.ResetPassReq) (*model.ResetPassResp, error)
	ChangePass(in *model.ChangePassReq) (model.ChangePassResp, error)
	InvalidateRefreshToken(token string) error
	IsRefreshTokenValid(token string) (bool, error)
	GetUser(in *pb.GetUserReq) (*pb.GetUserRes, error)
	DeleteUsers(in *pb.Id) (*pb.Massage, error)
	GetByIdUsers(in *pb.Id) (*pb.User, error)
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
					id, email, password, username)
				VALUES
					($1, $2, $3,$4)`
	_, err := U.DB.Exec(query, id, req.Email, req.Password, req.Username)
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
		DELETE FROM	
			refresh_tokens
		WHERE
			user_id = $1
	`
	_, err := U.DB.Exec(query, req.UserId)
	if err != nil {
		return err
	}

	query = `	
				INSERT INTO refresh_tokens(
					user_id, token, expires_at)
				VALUES
					($1, $2, $3)`
	_, err = U.DB.Exec(query, req.UserId, req.RefreshToken, req.ExpiresAt)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (U *userImpl) ResetPass(in *model.ResetPassReq) (*model.ResetPassResp, error) {
	_, err := U.DB.Exec(`
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

func (U *userImpl) ChangePass(in *model.ChangePassReq) (model.ChangePassResp, error) {
	_, err := U.DB.Exec(`
	    UPDATE users
		SET password = $1
		WHERE password = $2 and id=$3`, in.NewPassword, in.NowPassword, in.UserId)
	if err != nil {
		return model.ChangePassResp{
			Message: "error updating password",
		}, err
	}

	return model.ChangePassResp{
		Message: "Password changed successfully",
	}, nil
}

func (U *userImpl) InvalidateRefreshToken(userId string) error {
	_, err := U.DB.Exec(`
		DELETE FROM
			refresh_tokens
		WHERE
			user_id = $1
	`, userId)

	return err
}

func (U *userImpl) IsRefreshTokenValid(user_id string) (bool, error) {
	var count int
	err := U.DB.QueryRow(`
        SELECT 
            count(*)
        FROM 
            refresh_tokens 
        WHERE 
            user_id=$1 AND 
            expires_at > CURRENT_TIMESTAMP
			
    `, user_id).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (u *userImpl) GetUser(in *pb.GetUserReq) (*pb.GetUserRes, error) {
	rows, err := u.DB.Query(`
	    SELECT 
			id, 
			email, 
			username,
			name,
			password,
			lastname,
            birth_day,
            image,
            created_at,
            updated_at
		FROM users 
		WHERE deleted_at IS NULL
			limit=$1 offset=$2
			`)

	if err != nil {
		return nil, err
	}
	var users []*pb.User
	for rows.Next() {
		var user pb.User
		err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Username,
			&user.Name,
			&user.Password,
			&user.Lastname,
			&user.BirthDay,
			&user.Image,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return &pb.GetUserRes{Users: users}, nil
}

func (u *userImpl) DeleteUsers(in *pb.Id) (*pb.Massage, error) {
	deleted_at := time.Now().Format("2006/01/02")
	_, err := u.DB.Exec(`UPDATE users SET deleted_at=$1 WHERE id=$2 and deleted_at IS NULL`, deleted_at, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Massage{Message: "User deleted successfully"}, nil
}

func (u *userImpl) GetByIdUsers(in *pb.Id) (*pb.User, error) {
	var user pb.User
	err := u.DB.QueryRow(`
        SELECT 
            id, 
            email, 
            username,
            name,
            password,
            lastname,
            birth_day,
            image,
            created_at,
            updated_at
        FROM users 
		WHERE id=$1 AND 
			deleted_at IS NULL
            `, in.Id).Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.Name,
		&user.Password,
		&user.Lastname,
		&user.BirthDay,
		&user.Image,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
