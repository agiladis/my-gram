package repository

import (
	"my-gram/entity"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo entity.Photo) (entity.Photo, error)
	GetAll() ([]entity.Photo, error)
}

type photoRepository struct {
	DB *gorm.DB
}

func NewPhotoRepository(DB *gorm.DB) *photoRepository {
	return &photoRepository{DB}
}

func (pr *photoRepository) Create(photo entity.Photo) (entity.Photo, error) {
	err := pr.DB.Create(&photo).Error
	return photo, err
}

func (pr *photoRepository) GetAll() ([]entity.Photo, error) {
	var photos []entity.Photo
	// err := pr.DB.Preload("Users").Find(&photos).Error
	err := pr.DB.Model(&entity.Photo{}).Find(&photos).Order("updated_at ASC")

	return photos, err.Error
}
