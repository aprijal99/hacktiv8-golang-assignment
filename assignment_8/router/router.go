package router

import (
	"todolist/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/user/signup", controller.UserSignup)
	router.POST("/user/login", controller.UserLogin)

	return router
}
