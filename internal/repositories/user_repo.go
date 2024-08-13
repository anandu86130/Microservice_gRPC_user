package repositories

import (
	"fmt"
	userrepo "github.com/anandu86130/Microservice_gRPC_user/v2/internal/repositories/interface"

	"github.com/anandu86130/Microservice_gRPC_user/v2/internal/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) userrepo.UserRepository {
	return &UserRepo{DB: db}
}

func (u *UserRepo) CreateUser(user *model.UserModel) error {
	if err := u.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) CreateOTP(otp *model.VerifyOTPs) error {
	if err := u.DB.Create(&otp).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) FindUserByEmail(email string) (*model.UserModel, error) {
	var user model.UserModel
	err := u.DB.Where("email=?", email).First(&user).Error
	return &user, err
}

func (u *UserRepo) VerifyOTPcheck(email string, otp string) error {
	var user model.VerifyOTPs

	result := u.DB.Where("email = ? AND otp = ?", email, otp).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("email or OTP not found")
		}
		return result.Error
	}

	return nil

}
