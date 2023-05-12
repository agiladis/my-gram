package router

import (
	"my-gram/controller"
	"my-gram/middleware"
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

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository, validate)
	commentController := controller.NewCommentController(commentService)

	app := gin.Default()

	userRouter := app.Group("/users")
	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
	}

	photoRouter := app.Group("/photos")
	{
		photoRouter.Use(middleware.JWTMiddleware())
		photoRouter.GET("/", photoController.GetAll)
		photoRouter.GET("/:id", photoController.GetOne)
		photoRouter.POST("/", photoController.CreatePhoto)
		photoRouter.PUT("/:id", photoController.UpdatePhoto)
		photoRouter.DELETE("/:id", photoController.DeletePhoto)
	}

	commentRouter := app.Group("/comments")
	{
		commentRouter.Use(middleware.JWTMiddleware())
		commentRouter.GET("/", commentController.GetAll)
		commentRouter.GET("/:id", commentController.GetOne)
		commentRouter.POST("/", commentController.CreateComment)
		commentRouter.PUT("/:id", commentController.UpdateComment)
		commentRouter.DELETE("/:id", commentController.DeleteComment)
	}
	return app
}
