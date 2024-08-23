package postgres

import (
	pb "auth/generated/users"
	"database/sql"
)

type FollowRepository interface {
	AddFollower(in *pb.FollowerReq) (*pb.Massage, error)
	DeleteFollower(in *pb.DeleteFollowers) (*pb.Massage, error)
	GetFollowers(in *pb.GetFollowersReq) (*pb.GetaFollowersRes, error)
	GetByIdFollowers(in *pb.Id) (*pb.Followers, error)
	GetFollowing(in *pb.GetFollowingReq) (*pb.GetaFollowingRes, error)
	GetByIdFollowing(in *pb.Id) (*pb.Following, error)
	AddFollowing(in *pb.FollowingReq) (*pb.Massage, error)
	DeleteFollowing(in *pb.DeleteFollowings) (*pb.Massage, error)
}

type followRepo struct {
	DB *sql.DB
}

func NewFollowRepository(db *sql.DB) FollowRepository {
	return &followRepo{DB: db}
}

func (f *followRepo) AddFollower(in *pb.FollowerReq) (*pb.Massage, error) {
	_, err := f.DB.Exec(`
        INSERT INTO 
		user_followers(
            user_id, 
			follower_id
        )VALUES(
            $1, 
			$2
        )
    `, in.UserId, in.FollowerId)

	if err != nil {
		return nil, err
	}

	return &pb.Massage{Message: "Follower added successfully"}, nil
}

func (f *followRepo) DeleteFollower(in *pb.DeleteFollowers) (*pb.Massage, error) {
	_, err := f.DB.Exec(`
        UPDATE
        	user_followers
		SET
			deleted_at = now()
        WHERE 
            user_id = $1 AND follower_id = $2 AND deleted_at IS NULL
    `, in.UserId, in.FollowerId)

	if err != nil {
		return nil, err
	}

	return &pb.Massage{Message: "Follower deleted successfully"}, nil
}

func (f *followRepo) GetByIdFollowers(in *pb.Id) (*pb.Followers, error) {
	var follower pb.Followers
	err := f.DB.QueryRow(`
        SELECT 
            id, 
            user_id, 
            follower_id,
            created_at,
            updated_at
        FROM 
            user_followers
        WHERE 
            id = $1 AND deleted_at IS NULL
    `, in.Id).Scan(&follower.Id, &follower.UserId, &follower.FollowerId, &follower.CreatedAt, &follower.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &follower, nil
}

func (f *followRepo) GetFollowers(in *pb.GetFollowersReq) (*pb.GetaFollowersRes, error) {
	rows, err := f.DB.Query(`
        SELECT 
            id, 
            user_id, 
            follower_id,
            created_at,
            updated_at
        FROM 
            user_followers
        WHERE 
            deleted_at IS NULL limit $1 offset $2
    `, in.Limit, in.Page)

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

func (f *followRepo) AddFollowing(in *pb.FollowingReq) (*pb.Massage, error) {
	_, err := f.DB.Exec(`
        INSERT INTO 
        user_following(
            user_id, 
            following_id
            ) VALUES(
            $1, 
            $2
        )`, in.UserId, in.FollowingId)

	if err != nil {
		return nil, err
	}

	return &pb.Massage{Message: "Following added successfully"}, nil
}

func (f *followRepo) DeleteFollowing(in *pb.DeleteFollowings) (*pb.Massage, error) {
	_, err := f.DB.Exec(`
        UPDATE
            user_following
        SET
            deleted_at = now()
        WHERE user_id=$1 AND following_id=$2 AND deleted_at IS NULL`, in.UserId, in.FollowingId)

	if err != nil {
		return nil, err

	}

	return &pb.Massage{Message: "Following deleted successfully"}, nil
}

func (f *followRepo) GetByIdFollowing(in *pb.Id) (*pb.Following, error) {
	var following pb.Following
	err := f.DB.QueryRow(`
        SELECT 
            id, 
            user_id, 
            follower_id,
            created_at,
            updated_at
        FROM 
            user_following
        WHERE 
            id = $1 AND deleted_at IS NULL
    `, in.Id).Scan(&following.Id, &following.UserId, &following.FollowingId, &following.CreatedAt, &following.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &following, nil
}

func (f *followRepo) GetFollowing(in *pb.GetFollowingReq) (*pb.GetaFollowingRes, error) {
	rows, err := f.DB.Query(`
        SELECT 
            id, 
            user_id, 
            follower_id,
            created_at,
            updated_at
        FROM 
            user_following
        WHERE 
            deleted_at IS NULL limit $1 offset $2
    `, in.Limit, in.Page)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var followings []*pb.Following

	for rows.Next() {
		var following pb.Following
		err := rows.Scan(&following.Id, &following.UserId, &following.FollowingId, &following.CreatedAt, &following.UpdatedAt)

		if err != nil {
			return nil, err
		}

		followings = append(followings, &following)
	}

	return &pb.GetaFollowingRes{Following: followings}, nil
}
