package reids

import "github.com/redis/go-redis/v9"

type RedisStore interface {
}

type redisStoreImpl struct {
	client *redis.Client
}


