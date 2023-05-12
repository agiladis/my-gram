package service

import (
	"my-gram/entity"
	"my-gram/repository"

	"github.com/go-playground/validator/v10"
)

type CommentService interface {
	Create(commentRequest entity.Comment) (entity.Comment, error)
	GetAll() ([]entity.Comment, error)
}

type commentService struct {
	commentRepository repository.CommentRepository
	Validate          *validator.Validate
}

func NewCommentService(cr repository.CommentRepository, validate *validator.Validate) *commentService {
	return &commentService{
		commentRepository: cr,
		Validate:          validate,
	}
}

func (cs *commentService) Create(commentRequest entity.Comment) (entity.Comment, error) {

	// validate data
	cs.Validate = validator.New()
	if err := cs.Validate.Struct(commentRequest); err != nil {
		return entity.Comment{}, err
	}

	// hit repository
	return cs.commentRepository.Create(commentRequest)
}

func (cs *commentService) GetAll() ([]entity.Comment, error) {
	return cs.commentRepository.GetAll()
}
