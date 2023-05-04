package service

import (
	"my-gram/entity"
	"my-gram/repository"

	"github.com/go-playground/validator/v10"
)

type PhotoService interface {
	Create(photoRequest entity.Photo) (entity.PhotoResponse, error)
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

func (ps *photoService) Create(photoRequest entity.Photo) (entity.PhotoResponse, error) {

	// convert request into entity
	// data := entity.Photo{
	// 	Title:    photoRequest.Title,
	// 	Caption:  photoRequest.Caption,
	// 	PhotoURL: photoRequest.PhotoURL,
	// }

	// validate data
	ps.Validate = validator.New()
	err := ps.Validate.Struct(photoRequest)
	if err != nil {
		return entity.PhotoResponse{}, err
	}

	// hit repository
	photo, err := ps.photoRepository.Create(photoRequest)
	photoResponse := entity.PhotoResponse{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   1, // temporary response
	}

	return photoResponse, err
}
