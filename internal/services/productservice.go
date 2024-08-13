package services

import (
	pb "github.com/anandu86130/Microservice_gRPC_user/v2/internal/pb"
	product "github.com/anandu86130/Microservice_gRPC_user/v2/internal/product/handlers"
)

func (u *UserService) UserProductByName(userpb *pb.ProductByName) (*pb.ProductDetails, error) {
	result, err := product.FetchProductByNameHandler(u.client, userpb)
	if err != nil {
		return nil, err
	}

	return &pb.ProductDetails{
		Id:          result.Id,
		Category:    result.Category,
		Name:        result.Name,
		Price:       result.Price,
		Imagepath:   result.Imagepath,
		Description: result.Description,
		Size:        result.Size,
		Quantity:    result.Quantity,
	}, nil
}

func (u *UserService) UserProductByID(userpb *pb.ProductID) (*pb.ProductDetails, error) {
	result, err := product.FetchProductByIDHandler(u.client, userpb)
	if err != nil {
		return nil, err
	}

	return &pb.ProductDetails{
		Id:          result.Id,
		Category:    result.Category,
		Name:        result.Name,
		Price:       result.Price,
		Imagepath:   result.Imagepath,
		Description: result.Description,
		Size:        result.Size,
		Quantity:    result.Quantity,
	}, nil
}

func (u *UserService) UserProductList(userpb *pb.RNoParam) (*pb.ProductList, error) {
	result, err := product.FetchByProductsHandler(u.client, userpb)
	if err != nil {
		return nil, err
	}

	var productDetails []*pb.ProductDetails

	for _, i := range result.Item {
		productDetails = append(productDetails, &pb.ProductDetails{
			Id:          i.Id,
			Category:    i.Category,
			Name:        i.Name,
			Price:       i.Price,
			Imagepath:   i.Imagepath,
			Description: i.Description,
			Size:        i.Size,
			Quantity:    i.Quantity,
		})
	}

	return &pb.ProductList{Item: productDetails}, nil
}
