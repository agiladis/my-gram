package controller

import (
	"my-gram/entity"
	"my-gram/helper"
	"my-gram/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type socialMediaController struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaController(socialMediaService service.SocialMediaService) *socialMediaController {
	return &socialMediaController{socialMediaService}
}

func (smc *socialMediaController) CreateSocialMedia(ctx *gin.Context) {
	var socialMediaRequest entity.SocialMediaRequest

	if err := ctx.ShouldBindJSON(&socialMediaRequest); err != nil {
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
	socialMedia, err := smc.socialMediaService.Create(entity.Socialmedia{
		Name:           socialMediaRequest.Name,
		SocialMediaURL: socialMediaRequest.SocialmediaURL,
		UserID:         uint(accessClaim.AccessClaims.ID),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "social media created",
		"data":    socialMedia,
	})
}

func (smc *socialMediaController) GetAll(ctx *gin.Context) {
	socialMedia, err := smc.socialMediaService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get all social media success",
		"data":    socialMedia,
	})
}

func (smc *socialMediaController) GetOne(ctx *gin.Context) {
	socialMediaId := ctx.Param("id")
	socialMediaIdInt, err := strconv.Atoi(socialMediaId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// hit service
	socialMedia, err := smc.socialMediaService.GetById(socialMediaIdInt)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "social media not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "social media found",
		"data":    socialMedia,
	})
}
