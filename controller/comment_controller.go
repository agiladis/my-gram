package controller

import (
	"my-gram/entity"
	"my-gram/helper"
	"my-gram/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *commentController {
	return &commentController{commentService}
}

func (cc *commentController) CreateComment(ctx *gin.Context) {
	var commentRequest entity.CommentCreateRequest

	if err := ctx.ShouldBindJSON(&commentRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// get user id from ctx
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}

	// hit service
	comment, err := cc.commentService.Create(entity.Comment{
		UserID:  uint(accessClaim.AccessClaims.ID),
		PhotoID: commentRequest.PhotoID,
		Message: commentRequest.Message,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "comment posted",
		"data":    comment,
	})

}
