package repository

import (
	"my-gram/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user entity.User) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *userRepository {
	return &userRepository{DB}
}

func (ur *userRepository) Register(user entity.User) (entity.User, error) {
	err := ur.DB.Create(&user).Error
	return user, err
}

func (ur *userRepository) GetUserByUsername(username string) (entity.User, error) {
	userRes := entity.User{}

	err := ur.DB.Where("username = ?", username).First(&userRes).Error
	return userRes, err
}
