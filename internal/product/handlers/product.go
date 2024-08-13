package handlers

import (
	"context"
	"log"

	pb "github.com/anandu86130/Microservice_gRPC_user/v2/internal/pb"
	productpb "github.com/anandu86130/Microservice_gRPC_user/v2/internal/product/pb"
)

func FetchProductByIDHandler(client productpb.ProductServiceClient, p *pb.ProductID) (*productpb.ProductDetails, error) {
	ctx := context.Background()
	response, err := client.FetchProductByID(ctx, &productpb.ProductID{Id: p.Id})
	if err != nil {
		log.Printf("error while fetching the product by ID")
		return nil, err
	}
	return response, nil
}

func FetchProductByNameHandler(client productpb.ProductServiceClient, p *pb.ProductByName) (*productpb.ProductDetails, error) {
	ctx := context.Background()
	response, err := client.FetchProductByName(ctx, &productpb.ProductByName{Name: p.Name})
	if err != nil {
		log.Printf("error while fetching the product by name: %v", err)
		return nil, err
	}
	return response, nil
}

func FetchByProductsHandler(client productpb.ProductServiceClient, p *pb.RNoParam) (*productpb.ProductList, error) {
	ctx := context.Background()
	response, err := client.FetchProducts(ctx, &productpb.NoParam{})
	if err != nil {
		log.Printf("error while fetching the product details: %v", err)
		return nil, err
	}
	return response, nil
}
