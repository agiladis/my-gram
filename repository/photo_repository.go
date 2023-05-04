package repository

import (
	"my-gram/entity"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo entity.Photo) (entity.Photo, error)
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
