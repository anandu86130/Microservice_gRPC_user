package interfaces

import "github.com/anandu86130/Microservice_gRPC_user/v2/internal/model"

type UserRepository interface {
	CreateUser(user *model.UserModel) error
	CreateOTP(user *model.VerifyOTPs) error
	VerifyOTPcheck(email string, otp string) error
	FindUserByEmail(email string) (*model.UserModel, error)
}
