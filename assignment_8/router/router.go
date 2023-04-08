package router

import (
	"product/controller"
	"product/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/user")
	{
		userRouter.POST("/signup", controller.UserSignup)
		userRouter.POST("/login", controller.UserLogin)
	}

	productRouter := router.Group("/products")
	{
		productRouter.Use(middleware.AuthorizeRequest())
		productRouter.GET("/", controller.GetProductsOwnedByUserId)
		productRouter.POST("/", controller.CreateNewProduct)
	}

	return router
}
