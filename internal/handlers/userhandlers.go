package handlers

import (
	"context"
	"fmt"

	services_inter "github.com/anandu86130/Microservice_gRPC_user/v2/internal/services/interfaces"

	pb "github.com/anandu86130/Microservice_gRPC_user/v2/internal/pb"
)

type UserHandler struct {
	pb.UnimplementedUserServicesServer
	svc services_inter.UserService
}

func NewUserHandler(svc services_inter.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (u *UserHandler) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	result, err := u.svc.Signup(req)
	if err != nil {
		return nil, fmt.Errorf("failed to signup: %v", err)
	}
	return result, nil
}

func (u *UserHandler) VerifyOTP(ctx context.Context, req *pb.VerifyOTPRequest) (*pb.VerifyOTPResponse, error) {
	result, err := u.svc.VerifyOTP(req)
	if err != nil {
		return nil, fmt.Errorf("failed to verify OTP: %v", err)
	}
	return result, nil
}

func (u *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	result, err := u.svc.Login(req)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %v", err)
	}
	return result, nil
}
