package di

import (
	"log"

	"github.com/anandu86130/Microservice_gRPC_user/v2/config"
	"github.com/anandu86130/Microservice_gRPC_user/v2/internal/db"
	"github.com/anandu86130/Microservice_gRPC_user/v2/internal/handlers"
	"github.com/anandu86130/Microservice_gRPC_user/v2/internal/product"
	"github.com/anandu86130/Microservice_gRPC_user/v2/internal/repositories"
	"github.com/anandu86130/Microservice_gRPC_user/v2/internal/server"
	"github.com/anandu86130/Microservice_gRPC_user/v2/internal/services"
)

func Init() {
	config.LoadConfig()

	dbConn := db.ConnectDB()

	client, err := product.ClientDial()
	if err != nil {
		log.Fatalf("Error when dialing product service client:%v", err)
	}

	redisService, err := config.SetupRedis()
	if err != nil {
		log.Fatalf("Error when initializing Redis Client:%v", err)
	}

	userRepo := repositories.NewUserRepo(dbConn)

	userSvc := services.NewUserService(userRepo, client, redisService)

	userHandler := handlers.NewUserHandler(userSvc)
	server.NewGrpcServer(userHandler)
}
