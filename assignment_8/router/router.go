package router

import (
	"todolist/controller"
	"todolist/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/user")
	{
		userRouter.POST("/signup", controller.UserSignup)
		userRouter.POST("/login", controller.UserLogin)
	}

	todoRouter := router.Group("/todos")
	{
		todoRouter.Use(middleware.AuthorizeRequest())
		todoRouter.POST("/", controller.CreateProduct)
	}

	return router
}
