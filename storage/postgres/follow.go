package postgres

import (
	pb "auth/generated/users"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type FollowRepository interface {
	DeleteFollower(in *pb.DeleteFollowerReq) (*pb.Massage, error)
	GetFollowers(in *pb.GetFollowersReq) (*pb.GetaFollowersRes, error)
	GetByIdFollowers(in *pb.DeleteFollowerReq) (*pb.Follow, error)
	GetFollowing(in *pb.GetFollowingReq) (*pb.GetaFollowingRes, error)
	GetByIdFollowing(in *pb.DeleteFollowerReq) (*pb.Follow, error)
	Subscribe(in *pb.FollowingReq) (*pb.Massage, error)
	DeleteFollowing(in *pb.DeleteFollowerReq) (*pb.Massage, error)
}

type followRepo struct {
	DB *sql.DB
}

func NewFollowRepository(db *sql.DB) FollowRepository {
	return &followRepo{DB: db}
}

func (f *followRepo) DeleteFollower(in *pb.DeleteFollowerReq) (*pb.Massage, error) {
	tr, err := f.DB.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	query := `
				UPDATE user_followers SET
					deleted_at = $1
				WHERE 
					user_id = $2 AND follower_id = $3 AND deleted_at is null`
	_, err = tr.Exec(query, time.Now(), in.UserId, in.FollowerId)
	if err != nil {
		log.Println(err)
		tr.Rollback()
		return &pb.Massage{
			Message: "Error",
		}, err
	}
	query = `
				UPDATE user_following SET
					deleted_at = $1
				WHERE 
					user_id = $2 AND following_id = $3 AND deleted_at is null`
	_, err = tr.Exec(query, time.Now(), in.FollowerId, in.UserId)
	if err != nil {
		log.Println(err)
		tr.Rollback()
		return &pb.Massage{
			Message: "error",
		}, err
	}
	tr.Commit()
	return &pb.Massage{
		Message: "Success",
	}, nil
}

func (f *followRepo) GetByIdFollowers(in *pb.DeleteFollowerReq) (*pb.Follow, error) {
	tr, err := f.DB.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := `
				SELECT 
					id
				FROM
					user_followers
				WHERE
					user_id = $1 AND follower_id = $2`
	var id string
	err = tr.QueryRow(query, in.UserId, in.FollowerId).Scan(&id)
	if err != nil {
		log.Println(err)
		tr.Rollback()
		return nil, err
	} else if id == "" {
		log.Println("Followerlar orasida bunday follower mavjud emas")
		tr.Rollback()
		return nil, err
	}
	query = `
				SELECT 
					username, email, name, lastname, birth_day, image
				FROM
					Users
				WHERE 
					id = $1`
	var resp = pb.Follow{Id: in.FollowerId}
	err = tr.QueryRow(query, in.FollowerId).Scan(&resp.Username, &resp.Email, &resp.Name, &resp.Lastname, &resp.BirthDay, &resp.Image)
	if err != nil {
		log.Println(err)
		tr.Rollback()
		return nil, err
	}
	tr.Commit()
	return &resp, nil
}

func (f *followRepo) GetFollowers(in *pb.GetFollowersReq) (*pb.GetaFollowersRes, error) {
	rows, err := f.DB.Query(`
        SELECT 
    		id, user_id, follower_id, created_at
        FROM 
            user_followers
        WHERE 
            user_id = $1 AND deleted_at IS NULL limit $2 offset $3
    `, in.Id, in.Limit, in.Page)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var followers []*pb.Followers

	for rows.Next() {
		var follower pb.Followers
		err := rows.Scan(&follower.Id, &follower.UserId, &follower.FollowerId, &follower.CreatedAt, &follower.UpdatedAt)

		if err != nil {
			return nil, err
		}

		followers = append(followers, &follower)
	}

	return &pb.GetaFollowersRes{Followers: followers}, nil
}

func (f *followRepo) Subscribe(in *pb.FollowingReq) (*pb.Massage, error) {
	tranzaction, err := f.DB.Begin()
	if err != nil {
		log.Println(err)
		tranzaction.Rollback()
		return nil, err
	}
	id1 := uuid.NewString()
	query1 := `
				INSERT INTO user_following(
					id, user_id, following_id)
				VALUES
					($1, $2, $3)`
	_, err = tranzaction.Exec(query1, id1, in.UserId, in.FollowingId)
	if err != nil {
		log.Println(err)
		tranzaction.Rollback()
		return nil, err
	}
	id2 := uuid.NewString()
	query2 := `
				INSERT INTO user_followers(
					id, user_id, follower_id)
				VALUES
					($1, $2, $3)`
	_, err = tranzaction.Exec(query2, id2, in.FollowingId, in.UserId)
	if err != nil {
		log.Println(err)
		tranzaction.Rollback()
		return nil, err
	}
	err = tranzaction.Commit()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.Massage{
		Message: "Succes",
	}, nil
}

func (f *followRepo) DeleteFollowing(in *pb.DeleteFollowerReq) (*pb.Massage, error) {
	tr, err := f.DB.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	query := `
				UPDATE user_followers SET
					deleted_at = $1
				WHERE 
					user_id = $2 AND follower_id = $3 AND deleted_at is null`
	_, err = tr.Exec(query, time.Now(), in.FollowerId, in.UserId)
	if err != nil {
		log.Println(err)
		tr.Rollback()
		return &pb.Massage{
			Message: "Error",
		}, err
	}
	query = `
				UPDATE user_following SET
					deleted_at = $1
				WHERE 
					user_id = $2 AND following_id = $3 AND deleted_at is null`
	_, err = tr.Exec(query, time.Now(), in.UserId, in.FollowerId)
	if err != nil {
		log.Println(err)
		tr.Rollback()
		return &pb.Massage{
			Message: "error",
		}, err
	}
	tr.Commit()
	return &pb.Massage{
		Message: "Success",
	}, nil
}

func (f *followRepo) GetByIdFollowing(in *pb.DeleteFollowerReq) (*pb.Follow, error) {
	tr, err := f.DB.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := `
				SELECT 
					id
				FROM
					user_following
				WHERE
					user_id = $1 AND following_id = $2`
	var id string
	err = tr.QueryRow(query, in.UserId, in.FollowerId).Scan(&id)
	if err != nil {
		log.Println(err)
		tr.Rollback()
		return nil, err
	} else if id == "" {
		log.Println("Followerlar orasida bunday follower mavjud emas")
		tr.Rollback()
		return nil, err
	}
	query = `
				SELECT 
					username, email, name, lastname, birth_day, image
				FROM
					Users
				WHERE 
					id = $1`
	var resp = pb.Follow{Id: in.FollowerId}
	err = tr.QueryRow(query, in.FollowerId).Scan(&resp.Username, &resp.Email, &resp.Name, &resp.Lastname, &resp.BirthDay, &resp.Image)
	if err != nil {
		log.Println(err)
		tr.Rollback()
		return nil, err
	}
	tr.Commit()
	return &resp, nil
}

func (f *followRepo) GetFollowing(in *pb.GetFollowingReq) (*pb.GetaFollowingRes, error) {
	rows, err := f.DB.Query(`
        SELECT 
    		id, user_id, follower_id, created_at
        FROM 
            user_following
        WHERE 
            user_id = $1 AND deleted_at IS NULL limit $2 offset $3
    `, in.Id, in.Limit, in.Page)

	if err != nil {
		return nil, err
	}

	var followings []*pb.Following

	for rows.Next() {
		var following pb.Following
		err := rows.Scan(&following.Id, &following.UserId, &following.FollowingId, &following.CreatedAt)

		if err != nil {
			return nil, err
		}

		followings = append(followings, &following)
	}

	return &pb.GetaFollowingRes{
		Following: followings,
		Limit:     in.Limit,
		Page:      in.Page}, nil
}
