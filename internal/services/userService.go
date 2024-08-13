package services

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/anandu86130/Microservice_gRPC_user/v2/internal/model"
	utility "github.com/anandu86130/Microservice_gRPC_user/v2/internal/utility"

	"github.com/anandu86130/Microservice_gRPC_user/v2/config"
	generateotp "github.com/anandu86130/Microservice_gRPC_user/v2/internal/generateOTP"

	productpb "github.com/anandu86130/Microservice_gRPC_user/v2/internal/product/pb"
	userrepo "github.com/anandu86130/Microservice_gRPC_user/v2/internal/repositories/interface"
	services_inter "github.com/anandu86130/Microservice_gRPC_user/v2/internal/services/interfaces"

	pb "github.com/anandu86130/Microservice_gRPC_user/v2/internal/pb"
	utils "github.com/anandu86130/Microservice_gRPC_user/v2/internal/utils"
)

type UserService struct {
	repo        userrepo.UserRepository
	client      productpb.ProductServiceClient
	redisClient *config.RedisService
}

// UserProductDetails implements interfaces.UserService.
func (u *UserService) UserProductDetails(*pb.RNoParam) (*pb.ProductList, error) {
	panic("unimplemented")
}

func (u *UserService) Login(userpb *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := u.repo.FindUserByEmail(userpb.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found:%w", err)
	}

	token, err := utils.GenerateToken(user.Email, uint(user.UserID))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}
	return &pb.LoginResponse{Token: token}, nil
}

func (u *UserService) Signup(userpb *pb.SignupRequest) (*pb.SignupResponse, error) {
	existingUser, err := u.repo.FindUserByEmail(userpb.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user already exists")
	}
	otp := generateotp.GenerateOTP(6)
	utility.SendOTPByEmail(userpb.Email, otp)

	create := model.VerifyOTPs{Email: userpb.Email, Otp: otp}
	er := u.repo.CreateOTP(&create)
	if er != nil {
		log.Fatal("Failed to create OTP")
	}
	// user := &model.UserModel{Email: userpb.Email}
	storeotp := &model.VerifyOTPs{Email: userpb.Email, Otp: otp}

	userData, err := json.Marshal(&storeotp)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("user:%s", storeotp.Otp)

	err = u.redisClient.SetDataInRedis(key, userData, time.Minute*10)
	if err != nil {
		return nil, err
	}

	return &pb.SignupResponse{Message: "OTP has send to email"}, nil
}

func (u *UserService) VerifyOTP(userpb *pb.VerifyOTPRequest) (*pb.VerifyOTPResponse, error) {
	err := u.repo.VerifyOTPcheck(userpb.Email, userpb.Otp)
	if err != nil {
		log.Fatal("OTP not verification failed")
	}
	key := fmt.Sprintf("user:%s", userpb.Otp)

	userData, err := u.redisClient.GetFromRedis(key)
	if err != nil {
		return nil, err
	}

	var user model.UserModel
	err = json.Unmarshal([]byte(userData), &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %w", err)
	}

	err = u.repo.CreateUser(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	token, err := utils.GenerateToken(userpb.Email, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &pb.VerifyOTPResponse{Token: token}, nil
}

func NewUserService(repo userrepo.UserRepository, client productpb.ProductServiceClient, redisClient *config.RedisService) services_inter.UserService {
	return &UserService{
		repo:        repo,
		client:      client,
		redisClient: redisClient,
	}
}
