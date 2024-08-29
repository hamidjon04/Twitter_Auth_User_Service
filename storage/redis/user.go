package redisDb

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore interface {
	AddTokenBlacklisted(ctx context.Context, token string, expirationTime time.Duration) error
	IsTokenBlacklisted(ctx context.Context, token string) (bool, error)
	StoreCode(ctx context.Context, email, code string, exprationTime time.Duration) error
	IsCodeValid(ctx context.Context, email, code string) (bool, error)
}

type redisStoreImpl struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) RedisStore {
	return &redisStoreImpl{
		client: client,
	}
}

func (rdb *redisStoreImpl) AddTokenBlacklisted(ctx context.Context, token string, expirationTime time.Duration) error {
	err := rdb.client.Set(ctx, token, "blacklisted", expirationTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rdb *redisStoreImpl) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	val, err := rdb.client.Get(ctx, token).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return val == "blacklisted", nil
}

func (rdb *redisStoreImpl) StoreCode(ctx context.Context, email, code string, exprationTime time.Duration) error {
	err := rdb.client.Set(ctx, email+":code", code, exprationTime).Err()
	
	return err
}

func (rdb *redisStoreImpl) IsCodeValid(ctx context.Context, email, code string) (bool, error) {
	val, err := rdb.client.Get(ctx, email+":code").Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return val == code, nil
}
