package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	USER_ROUTER          string
	USER_SERVICE         string
	DB_HOST              string
	DB_PORT              string
	DB_USER              string
	DB_PASSWORD          string
	DB_NAME              string
	ACCESS_TOKEN_SECRET  string
	REFRESH_TOKEN_SECRET string
	SIGNING_KEY          string
	Redis_HOST           string
	Redis_PASSWORD       string
	Redis_DB             int
	Redis_PORT           int
}

func LoadConfig() Config {
	if err := godotenv.Load("./../../.env"); err != nil {
		log.Println("error loading .env file or not found", err)
	}

	config := Config{}

	config.USER_SERVICE = cast.ToString(coalesce("USER_SERVICE", ":50050"))
	config.USER_ROUTER = cast.ToString(coalesce("USER_ROUTER", ":1596"))
	config.SIGNING_KEY = cast.ToString(coalesce("SIGNING_KEY", "secret"))

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "your-pasword"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "your-db-name"))

	config.Redis_HOST = cast.ToString(coalesce("REDIS_HOST", "localhost"))
	config.Redis_PASSWORD = cast.ToString(coalesce("REDIS_PASSWORD", "your-redis-password"))
	config.Redis_DB = cast.ToInt(coalesce("REDIS_DB", 0))
	config.Redis_PORT = cast.ToInt(coalesce("REDIS_PORT", 6379))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
