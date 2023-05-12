package service

import (
	"errors"
	"my-gram/entity"
	"my-gram/repository"

	"github.com/go-playground/validator/v10"
)

type CommentService interface {
	Create(commentRequest entity.Comment) (entity.Comment, error)
	GetAll() ([]entity.Comment, error)
	GetById(id int) (entity.Comment, error)
	Update(commentId, userId int, newComment entity.CommentUpdateRequest) (entity.Comment, error)
	Delete(commentId, userId int) error
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

func (cs *commentService) GetById(id int) (entity.Comment, error) {
	return cs.commentRepository.GetById(id)
}

func (cs *commentService) Update(commentId, userId int, newComment entity.CommentUpdateRequest) (entity.Comment, error) {
	comment, err := cs.commentRepository.GetById(commentId)
	if err != nil {
		return entity.Comment{}, err
	}

	// authorization check
	if comment.UserID != uint(userId) {
		return entity.Comment{}, errors.New("unauthorized")
	}

	// assign new comment data
	comment.Message = newComment.Message

	// validate data
	cs.Validate = validator.New()
	if err = cs.Validate.Struct(comment); err != nil {
		return entity.Comment{}, err
	}

	// hit repository
	err = cs.commentRepository.Update(commentId, comment)
	return comment, err
}

func (cs *commentService) Delete(commentId, userId int) error {
	comment, err := cs.commentRepository.GetById(commentId)
	if err != nil {
		return err
	}

	// authorization check
	if comment.UserID != uint(userId) {
		return errors.New("unauthorized")
	}

	// hit repository
	err = cs.commentRepository.Delete(commentId)
	return err
}
