package server

import (
	"log"
	"net"

	"github.com/anandu86130/Microservice_gRPC_user/v2/internal/handlers"
	pb "github.com/anandu86130/Microservice_gRPC_user/v2/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer(l *handlers.UserHandler) {
	log.Println("connecting to gRPC server")
	lis, err := net.Listen("tcp", ":8082")

	if err != nil {
		log.Fatal("error creating listener on port 8082")
	}

	grp := grpc.NewServer()
	pb.RegisterUserServicesServer(grp, l)
	reflection.Register(grp)

	log.Printf("listening on gRPC server on 8082")
	if err := grp.Serve(lis); err != nil {
		log.Fatal("error while connecting to gRPC server")
	}
}
