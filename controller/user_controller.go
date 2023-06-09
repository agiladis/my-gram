package controller

import (
	"my-gram/entity"
	"my-gram/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (uc *userController) Register(ctx *gin.Context) {
	var userRequest entity.UserCreateRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// hit service
	user, err := uc.userService.Register(userRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"data":    user,
	})
}

func (uc *userController) Login(ctx *gin.Context) {
	userAuth := entity.UserAuthenticate{}

	if err := ctx.ShouldBindJSON(&userAuth); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// hit service
	authRes, err := uc.userService.AuthenticateUser(userAuth)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"data":    authRes,
	})
}
