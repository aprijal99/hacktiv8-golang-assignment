package router

import (
	"book-service/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controller.GetAllBooks)
	router.GET("/books/:id", controller.GetBookById)
	router.POST("/books", controller.CreateNewBook)
	router.PUT("/books/:id", controller.UpdateBook)
	router.DELETE("/books/:id", controller.DeleteBook)

	return router
}
