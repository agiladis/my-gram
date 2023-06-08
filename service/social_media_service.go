package service

import (
	"my-gram/entity"
	"my-gram/repository"

	"github.com/go-playground/validator/v10"
)

type SocialMediaService interface {
	Create(socialMediaRequest entity.Socialmedia) (entity.Socialmedia, error)
}

type socialMediaService struct {
	socialMediaRepository repository.SocialMediaRepository
	Validate              *validator.Validate
}

func NewSocialMediaService(smr repository.SocialMediaRepository, validate *validator.Validate) *socialMediaService {
	return &socialMediaService{
		socialMediaRepository: smr,
		Validate:              validate,
	}
}

func (sms *socialMediaService) Create(socialMediaRequest entity.Socialmedia) (entity.Socialmedia, error) {
	// validate data
	sms.Validate = validator.New()
	err := sms.Validate.Struct(socialMediaRequest)
	if err != nil {
		return entity.Socialmedia{}, err
	}

	// hit repository
	return sms.socialMediaRepository.Create(socialMediaRequest)
}
