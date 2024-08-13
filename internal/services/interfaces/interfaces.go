package interfaces

import pb "github.com/anandu86130/Microservice_gRPC_user/v2/internal/pb"

type UserService interface {
	Signup(userpb *pb.SignupRequest) (*pb.SignupResponse, error)
	VerifyOTP(userpb *pb.VerifyOTPRequest) (*pb.VerifyOTPResponse, error)
	Login(userpb *pb.LoginRequest) (*pb.LoginResponse, error)
	UserProductByName(*pb.ProductByName) (*pb.ProductDetails, error)
	UserProductByID(*pb.ProductID) (*pb.ProductDetails, error)
	UserProductDetails(*pb.RNoParam) (*pb.ProductList, error)
}
