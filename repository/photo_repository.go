package repository

import (
	"errors"
	"my-gram/entity"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo entity.Photo) (entity.Photo, error)
	GetAll() ([]entity.Photo, error)
	GetById(id int) (entity.Photo, error)
	Update(id int, photo entity.Photo) error
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

func (pr *photoRepository) GetById(id int) (entity.Photo, error) {
	var photo entity.Photo
	err := pr.DB.Where("id = ?", id).First(&photo).Error
	return photo, err
}

func (pr *photoRepository) Update(id int, photo entity.Photo) error {
	result := pr.DB.Model(&entity.Photo{}).Where("id = ?", id).Updates(&photo)
	if result.RowsAffected == 0 {
		return errors.New("there is no data to update")
	}

	return nil
}
