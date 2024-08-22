package redisDb

import (
	"auth/config"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(cfg config.Config) *redis.Client {
	addr := fmt.Sprintf("%s:%d", cfg.Redis_HOST, cfg.Redis_PORT)
	log.Println("Connecting to Redis at", addr)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Redis_PASSWORD,
		DB:       cfg.Redis_DB,
	})

	return client
}
