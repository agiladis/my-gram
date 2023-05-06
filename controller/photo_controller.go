package controller

import (
	"my-gram/entity"
	"my-gram/helper"
	"my-gram/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type photoController struct {
	photoService service.PhotoService
}

func NewPhotoController(photoService service.PhotoService) *photoController {
	return &photoController{photoService}
}

func (pc *photoController) CreatePhoto(ctx *gin.Context) {
	var photoRequest entity.PhotoCreateRequest

	err := ctx.ShouldBindJSON(&photoRequest)
	if err != nil {
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
	photo, err := pc.photoService.Create(entity.Photo{
		Title:    photoRequest.Title,
		Caption:  photoRequest.Caption,
		PhotoURL: photoRequest.PhotoURL,
		UserID:   uint(accessClaim.AccessClaims.ID),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "photo posted",
		"data":    photo,
	})
}

func (pc *photoController) GetAll(ctx *gin.Context) {
	photos, err := pc.photoService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all photos success",
		"data":    photos,
	})

}

func (pc *photoController) GetOne(ctx *gin.Context) {
	photoId := ctx.Param("id")
	photoIdInt, err := strconv.Atoi(photoId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// hit service
	photo, err := pc.photoService.GetById(photoIdInt)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "photo not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "photo found",
		"data":    photo,
	})

}

func (pc *photoController) UpdatePhoto(ctx *gin.Context) {
	var photoRequest entity.PhotoCreateRequest

	err := ctx.ShouldBindJSON(&photoRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	photoId := ctx.Param("id")
	photoIdInt, err := strconv.Atoi(photoId)
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
	photo, err := pc.photoService.Update(photoIdInt, accessClaim.AccessClaims.ID, photoRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "update foto success",
		"data":    photo,
	})
}
