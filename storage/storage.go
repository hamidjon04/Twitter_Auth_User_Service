package storage

import (
	"auth/storage/postgres"
	redisDb "auth/storage/redis"
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type IStorage interface {
	UserRepo()postgres.UserRepo
	RedisUserRepo() redisDb.RedisStore
}

type storageImpl struct {
	redis    *redis.Client
	postgres *sql.DB
}

func NewStorage(client *redis.Client, db *sql.DB) IStorage {
	return &storageImpl{redis: client, postgres: db}
}

func(u *storageImpl) UserRepo()postgres.UserRepo{
	return postgres.NewUserRepo(u.postgres)
}

func(u *storageImpl) RedisUserRepo() redisDb.RedisStore{
	return redisDb.NewRedisStore(u.redis)
}
