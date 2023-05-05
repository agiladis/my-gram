package controller

import (
	"my-gram/entity"
	"my-gram/service"
	"net/http"

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

	// hit service
	photo, err := pc.photoService.Create(entity.Photo{
		Title:    photoRequest.Title,
		Caption:  photoRequest.Caption,
		PhotoURL: photoRequest.PhotoURL,
		UserID:   1,
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
