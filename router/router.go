package router

import (
	"my-gram/controller"
	"my-gram/repository"
	"my-gram/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine {

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	app := gin.Default()

	userRouter := app.Group("/user")
	{
		userRouter.POST("/register", userController.Register)
	}

	return app
}
