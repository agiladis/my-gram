package repository

import (
	"my-gram/entity"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment entity.Comment) (entity.Comment, error)
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
