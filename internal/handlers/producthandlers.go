package handlers

import (
	"context"

	pb "github.com/anandu86130/Microservice_gRPC_user/v2/internal/pb"
)

func (u *UserHandler)UserProductByName(ctx context.Context, p *pb.ProductByName) (*pb.ProductDetails, error) {
	result, err := u.svc.UserProductByName(p)
	if err != nil{
		return nil, err
	}
	return result, nil
}

func (u *UserHandler) UserProductByID(ctx context.Context, p *pb.ProductID) (*pb.ProductDetails, error){
	result, err := u.svc.UserProductByID(p)
	if err != nil{
		return nil, err
	}
	return result,nil
}

func (u *UserHandler) UserProductDetails(ctx context.Context, p *pb.RNoParam) (*pb.ProductList, error){
	result, err := u.svc.UserProductDetails(p)
	if err != nil{
		return nil,	err
	}
	return result, nil
}