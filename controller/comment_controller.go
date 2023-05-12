package controller

import (
	"my-gram/entity"
	"my-gram/helper"
	"my-gram/service"
	"net/http"
	"strconv"

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

func (cc *commentController) GetAll(ctx *gin.Context) {
	comments, err := cc.commentService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all comment success",
		"data":    comments,
	})
}

func (cc *commentController) GetOne(ctx *gin.Context) {
	commentId := ctx.Param("id")
	commentIdInt, err := strconv.Atoi(commentId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// hit service
	comment, err := cc.commentService.GetById(commentIdInt)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "comment not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get comment success",
		"data":    comment,
	})
}

func (cc *commentController) UpdateComment(ctx *gin.Context) {
	var commentRequest entity.CommentUpdateRequest

	err := ctx.ShouldBindJSON(&commentRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	commentId := ctx.Param("id")
	commentIdInt, err := strconv.Atoi(commentId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// get user info from ctx
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	// hit service
	comment, err := cc.commentService.Update(commentIdInt, accessClaim.AccessClaims.ID, commentRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "update comment success",
		"data":    comment,
	})
}
