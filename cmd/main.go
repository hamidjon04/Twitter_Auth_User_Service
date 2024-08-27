package main

import (
	"auth/api"
	"auth/config"
	"auth/generated/users"
	logs "auth/pkg"
	"auth/service"
	"auth/storage"
	"auth/storage/postgres"
	redisDb "auth/storage/redis"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	logger := logs.InitLogger()
	cfg := config.LoadConfig()

	db, err := postgres.ConnectToDB(cfg)
	if err != nil {
		log.Println(err)
		logger.Error(err.Error())
		panic(err)
	}
	defer db.Close()

	rdb := redisDb.ConnectRedis(cfg)
	defer rdb.Close()

	storage := storage.NewStorage(rdb, db)

	listener, err := net.Listen("tcp", cfg.USER_SERVICE)
	if err != nil {
		log.Println(err)
		logger.Error(err.Error())
		panic(err)
	}
	defer listener.Close()

	u := service.NewAuthenticateService(storage, logger)
	serv := service.NewService(storage, logger)
	s := grpc.NewServer()
	users.RegisterUserServiceServer(s, serv)

	go func ()  {
		controller := api.NewController(gin.Default(), logger)
		controller.StartRouter(cfg)	
		controller.SetUpRouter(u, *serv)
	}()

	log.Printf("Service is run: %v", cfg.USER_SERVICE)
	if err = s.Serve(listener); err != nil {
		log.Println(err)
		logger.Error(err.Error())
		panic(err)
	}
}
