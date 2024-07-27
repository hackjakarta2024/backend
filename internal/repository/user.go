package repository

import (
	"github.com/hackjakarta2024/backend/internal/model"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	CreateUser(user model.User) error
	GetUserByEmail(email string) (model.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) CreateUser(user model.User) error {
	err := r.DB.Create(&user).Error
	return err
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return user, err
}
