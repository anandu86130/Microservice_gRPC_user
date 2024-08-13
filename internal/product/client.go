package product

import (
	productpb "github.com/anandu86130/Microservice_gRPC_user/v2/internal/product/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial() (productpb.ProductServiceClient, error) {
	grpc, err := grpc.NewClient("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error when connecting to gRPC client:8083")
		return nil, err
	}
	log.Printf("Successfully connected to gRPC client:8083")
	return productpb.NewProductServiceClient(grpc), nil
}
