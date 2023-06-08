package repository

import (
	"my-gram/entity"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(socialMedia entity.Socialmedia) (entity.Socialmedia, error)
	GetAll() ([]entity.Socialmedia, error)
}

type socialMediaRepository struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(DB *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{DB}
}

func (smr *socialMediaRepository) Create(socialMedia entity.Socialmedia) (entity.Socialmedia, error) {
	err := smr.DB.Create(&socialMedia).Error
	return socialMedia, err
}

func (smr *socialMediaRepository) GetAll() ([]entity.Socialmedia, error) {
	var socialMedia []entity.Socialmedia
	err := smr.DB.Model(&entity.Socialmedia{}).Find(&socialMedia).Order("updated_at ASC")
	return socialMedia, err.Error
}
