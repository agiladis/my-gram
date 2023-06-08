package service

import (
	"errors"
	"my-gram/entity"
	"my-gram/repository"

	"github.com/go-playground/validator/v10"
)

type SocialMediaService interface {
	Create(socialMediaRequest entity.Socialmedia) (entity.Socialmedia, error)
	GetAll() ([]entity.Socialmedia, error)
	GetById(socialMediaId int) (entity.Socialmedia, error)
	Update(socialMediaId, userId int, newSocialMedia entity.SocialMediaRequest) (entity.Socialmedia, error)
	Delete(socialMediaId, userId int) error
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

func (sms *socialMediaService) GetAll() ([]entity.Socialmedia, error) {
	return sms.socialMediaRepository.GetAll()
}

func (sms *socialMediaService) GetById(socialMediaId int) (entity.Socialmedia, error) {
	return sms.socialMediaRepository.GetById(socialMediaId)
}

func (sms *socialMediaService) Update(socialMediaId, userId int, newSocialMedia entity.SocialMediaRequest) (entity.Socialmedia, error) {
	socialMedia, err := sms.socialMediaRepository.GetById(socialMediaId)
	if err != nil {
		return entity.Socialmedia{}, err
	}

	// authorization check
	if socialMedia.UserID != uint(userId) {
		return entity.Socialmedia{}, errors.New("unauthorized")
	}

	// assign new social media data
	socialMedia.Name = newSocialMedia.Name
	socialMedia.SocialMediaURL = newSocialMedia.SocialmediaURL

	// validate data
	sms.Validate = validator.New()
	err = sms.Validate.Struct(socialMedia)
	if err != nil {
		return entity.Socialmedia{}, err
	}

	// hit repository
	err = sms.socialMediaRepository.Update(socialMediaId, socialMedia)
	return socialMedia, err
}

func (sms *socialMediaService) Delete(socialMediaId, userId int) error {
	socialMedia, err := sms.socialMediaRepository.GetById(socialMediaId)
	if err != nil {
		return err
	}

	// authorization check
	if socialMedia.UserID != uint(userId) {
		return errors.New("unauthorized")
	}

	// hit repository
	err = sms.socialMediaRepository.Delete(socialMediaId)
	return err
}
