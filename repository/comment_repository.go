package repository

import (
	"my-gram/entity"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment entity.Comment) (entity.Comment, error)
	GetAll() ([]entity.Comment, error)
	GetById(id int) (entity.Comment, error)
}

type commentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(DB *gorm.DB) *commentRepository {
	return &commentRepository{DB}
}

func (cr *commentRepository) Create(comment entity.Comment) (entity.Comment, error) {
	err := cr.DB.Create(&comment).Error
	return comment, err
}

func (cr *commentRepository) GetAll() ([]entity.Comment, error) {
	var comments []entity.Comment
	err := cr.DB.Model(&entity.Comment{}).Find(&comments).Order("updated_at ASC")
	return comments, err.Error
}

func (cr *commentRepository) GetById(id int) (entity.Comment, error) {
	var comment entity.Comment
	err := cr.DB.Where("id = ?", id).First(&comment).Error
	return comment, err
}