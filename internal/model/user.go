package model

type UserModel struct {
	UserID uint   `gorm:"PrimaryKey"`
	Email  string `json:"email" validate:"require"`
	Role   string `json:"role" gorm:"not null;default:'user'"`
}
