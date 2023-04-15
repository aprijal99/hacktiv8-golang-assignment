package router

import (
	"mygram/controller"
	"mygram/middleware"

	_ "mygram/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram API
// @version 1.0
// @description This is a sample service for managing MyGram API
// @termOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:4000
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	// Users
	userRouter := router.Group("/users")
	{
		// Signup
		userRouter.POST("/signup", controller.UserSignup)
		// Login
		userRouter.POST("/login", controller.UserLogin)
	}

	// Photos
	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middleware.AuthenticateRequest())
		// Get All
		photoRouter.GET("/", controller.GetAllPhotos)
		// Get by Id
		photoRouter.GET("/:id", controller.GetOnePhoto)
		// Create
		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.Use(middleware.AuthorizePhoto())
		// Update
		photoRouter.PUT("/:id", controller.UpdatePhoto)
		// Delete
		photoRouter.DELETE("/:id", controller.DeletePhoto)
	}

	// Socialmedia
	socialMediaRouter := router.Group("/socialmedia")
	{
		socialMediaRouter.Use(middleware.AuthenticateRequest())
		// Get All
		socialMediaRouter.GET("/", controller.GetAllSocialMedia)
		// Get by Id
		socialMediaRouter.GET("/:id", controller.GetOneSocialMedia)
		// Create
		socialMediaRouter.POST("/", controller.CreateSocialMedia)
		socialMediaRouter.Use(middleware.AuthorizeSocialMedia())
		// Update
		socialMediaRouter.PUT("/:id", controller.UpdateSocialMedia)
		// Delete
		socialMediaRouter.DELETE("/:id", controller.DeleteSocialMedia)
	}

	// Comments
	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middleware.AuthenticateRequest())
		// Get All
		commentRouter.GET("/", controller.GetAllComments)
		// Get by Id
		commentRouter.GET("/:id", controller.GetOneComment)
		// Create
		commentRouter.POST("/", controller.CreateComment)
		commentRouter.Use(middleware.AuthorizeComment())
		// Update
		commentRouter.PUT("/:id", controller.UpdateComment)
		// Delete
		commentRouter.DELETE("/:id", controller.DeleteComment)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
