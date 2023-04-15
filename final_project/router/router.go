package router

import (
	"mygram/controller"
	"mygram/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/signup", controller.UserSignup)
		userRouter.POST("/login", controller.UserLogin)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middleware.AuthenticateRequest())
		photoRouter.GET("/", controller.GetAllPhotos)
		photoRouter.GET("/:id", controller.GetOnePhoto)
		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.Use(middleware.AuthorizePhoto())
		photoRouter.PUT("/:id", controller.UpdatePhoto)
		photoRouter.DELETE("/:id", controller.DeletePhoto)
	}

	socialMediaRouter := router.Group("/socialmedia")
	{
		socialMediaRouter.Use(middleware.AuthenticateRequest())
		socialMediaRouter.GET("/", controller.GetAllSocialMedia)
		socialMediaRouter.GET("/:id", controller.GetOneSocialMedia)
		socialMediaRouter.POST("/", controller.CreateSocialMedia)
		socialMediaRouter.Use(middleware.AuthorizeSocialMedia())
		socialMediaRouter.PUT("/:id", controller.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", controller.DeleteSocialMedia)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middleware.AuthenticateRequest())
		commentRouter.GET("/", controller.GetAllComments)
		commentRouter.GET("/:id", controller.GetOneComment)
		commentRouter.POST("/", controller.CreateComment)
		commentRouter.Use(middleware.AuthorizeComment())
		commentRouter.PUT("/:id", controller.UpdateComment)
		commentRouter.DELETE("/:id", controller.DeleteComment)
	}

	return router
}
