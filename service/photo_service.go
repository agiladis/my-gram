package service

import (
	"errors"
	"my-gram/entity"
	"my-gram/repository"

	"github.com/go-playground/validator/v10"
)

type PhotoService interface {
	Create(photoRequest entity.Photo) (entity.Photo, error)
	GetAll() ([]entity.Photo, error)
	GetById(id int) (entity.Photo, error)
	Update(photoId, userId int, newPhoto entity.PhotoCreateRequest) (entity.Photo, error)
	Delete(photoId, userId int) error
}

type photoService struct {
	photoRepository repository.PhotoRepository
	Validate        *validator.Validate
}

func NewPhotoService(pr repository.PhotoRepository, validate *validator.Validate) *photoService {
	return &photoService{
		photoRepository: pr,
		Validate:        validate,
	}
}

func (ps *photoService) Create(photoRequest entity.Photo) (entity.Photo, error) {

	// validate data
	ps.Validate = validator.New()
	err := ps.Validate.Struct(photoRequest)
	if err != nil {
		return entity.Photo{}, err
	}

	// hit repository
	photo, err := ps.photoRepository.Create(photoRequest)

	return photo, err
}

func (ps *photoService) GetAll() ([]entity.Photo, error) {
	return ps.photoRepository.GetAll()
}

func (ps *photoService) GetById(id int) (entity.Photo, error) {
	return ps.photoRepository.GetById(id)
}

func (ps *photoService) Update(photoId, userId int, newPhoto entity.PhotoCreateRequest) (entity.Photo, error) {
	photo, err := ps.photoRepository.GetById(photoId)
	if err != nil {
		return entity.Photo{}, err
	}

	// authorization check
	if photo.UserID != uint(userId) {
		return entity.Photo{}, errors.New("unauthorized")
	}

	// assign new photo data
	photo.Title = newPhoto.Title
	photo.Caption = newPhoto.Caption
	photo.PhotoURL = newPhoto.PhotoURL

	// hit repository
	err = ps.photoRepository.Update(photoId, photo)
	return photo, err
}

func (ps *photoService) Delete(photoId, userId int) error {
	photo, err := ps.photoRepository.GetById(photoId)
	if err != nil {
		return err
	}

	// authorization check
	if photo.UserID != uint(userId) {
		return errors.New("unauthorized")
	}

	// hit repository
	err = ps.photoRepository.Delete(photoId)
	return err
}
