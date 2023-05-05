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

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository, validate)
	photoController := controller.NewPhotoController(photoService)

	app := gin.Default()

	userRouter := app.Group("/users")
	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
	}

	photoRouter := app.Group("/photos")
	{
		photoRouter.GET("/", photoController.GetAll)
		// photoRouter.Use(middleware.Auth())
		// photoRouter.GET("/:id", photoController.GetPhotoById)
		photoRouter.POST("/", photoController.CreatePhoto)
		// photoRouter.DELETE("/:id", photoController.DeletePhoto)
		// photoRouter.PUT("/:id", photoController.UpdatePhoto)
	}

	return app
}
