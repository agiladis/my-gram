package router

import (
	"my-gram/controller"
	"my-gram/repository"
	"my-gram/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine {
	var validate *validator.Validate

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validate)
	userController := controller.NewUserController(userService)

	app := gin.Default()

	userRouter := app.Group("/users")
	{
		userRouter.POST("/register", userController.Register)
	}

	return app
}
