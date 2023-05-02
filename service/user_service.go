package service

import (
	"my-gram/entity"
	"my-gram/helper"
	"my-gram/repository"

	"github.com/go-playground/validator/v10"
)

type UserService interface {
	Register(userRequest entity.UserCreateRequest) (entity.UserResponse, error)
}

type userService struct {
	userRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(ur repository.UserRepository, validate *validator.Validate) *userService {
	return &userService{
		userRepository: ur,
		Validate:       validate,
	}
}

func (us *userService) Register(userRequest entity.UserCreateRequest) (entity.UserResponse, error) {

	// convert request into entity
	data := entity.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Age:      userRequest.Age,
	}

	// validate data
	us.Validate = validator.New()
	err := us.Validate.Struct(data)
	if err != nil {
		return entity.UserResponse{}, err
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
