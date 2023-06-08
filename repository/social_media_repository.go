package repository

import (
	"my-gram/entity"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(socialMedia entity.Socialmedia) (entity.Socialmedia, error)
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
