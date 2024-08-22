package storage

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type IStorage interface {
}

type storageImpl struct {
	redis    *redis.Client
	postgres *sql.DB
}

func NewStorage(client *redis.Client, db *sql.DB) IStorage {
	return &storageImpl{redis: client, postgres: db}
}
