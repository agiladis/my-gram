package repository

import (
	"errors"
	"my-gram/entity"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(socialMedia entity.Socialmedia) (entity.Socialmedia, error)
	GetAll() ([]entity.Socialmedia, error)
	GetById(socialMediaId int) (entity.Socialmedia, error)
	Update(socialMediaId int, socialMedia entity.Socialmedia) error
	Delete(socialMediaId int) error
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

func (smr *socialMediaRepository) GetById(socialMediaId int) (entity.Socialmedia, error) {
	var socialMedia entity.Socialmedia
	err := smr.DB.Where("id = ?", socialMediaId).First(&socialMedia).Error
	return socialMedia, err
}

func (smr *socialMediaRepository) Update(socialMediaId int, socialMedia entity.Socialmedia) error {
	result := smr.DB.Model(&entity.Socialmedia{}).Where("id = ?", socialMediaId).Updates(&socialMedia)
	if result.RowsAffected == 0 {
		return errors.New("there is no data to update")
	}

	return nil
}

func (smr *socialMediaRepository) Delete(socialMediaId int) error {
	result := smr.DB.Model(&entity.Socialmedia{}).Where("id = ?", socialMediaId).Delete(&entity.Socialmedia{})
	if result.RowsAffected == 0 {
		return errors.New("there is no data to delete")
	}

	return nil
}
