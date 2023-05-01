package service

import (
	"my-gram/entity"
	"my-gram/helper"
	"my-gram/repository"
)

type UserService interface {
	Register(userRequest entity.UserCreateRequest) (entity.UserResponse, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) *userService {
	return &userService{ur}
}

func (us *userService) Register(userRequest entity.UserCreateRequest) (entity.UserResponse, error) {
	// convert request into entity
	data := entity.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Age:      userRequest.Age,
	}

	// hashing password
	data.Password = helper.HashPass(data.Password)

	// hit repository
	user, err := us.userRepository.Register(data)
	userResponse := entity.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}

	return userResponse, err
}
