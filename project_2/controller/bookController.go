package controller

import (
	"book/database"
	"book/entity"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Books = []entity.Book{
	{Id: 1, Title: "Calculus 1", Author: "John F Bach", Desc: "Calculus 1 covers mostly about Derivatives"},
	{Id: 2, Title: "Introduction of Biochemistry", Author: "Lehninger", Desc: "This is introduction to biomolecules"},
}

func GetAllBooks(ctx *gin.Context) {
	db := database.GetDB()
	var books []entity.Book

	db.Find(&books)

	ctx.JSON(http.StatusOK, books)
}

func GetBookById(ctx *gin.Context) {
	db := database.GetDB()
	var book entity.Book

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": "The id must be a number",
		})
		return
	}

	errFind := db.First(&book, "id = ?", id).Error
	if errFind != nil {
		if errors.Is(errFind, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"status":  "Not Found",
				"message": fmt.Sprintf("The book with id = %d is not found in the database", id),
			})
			return
		}
	}

	ctx.JSON(http.StatusFound, book)
}

func CreateNewBook(ctx *gin.Context) {
	db := database.GetDB()
	var newBook entity.Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": err,
		})
		return
	}

	err := db.Create(&newBook).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"status":  "Internal Server Error",
			"message": "Error creating new book record, there is something error with the server",
		})
		return
	}

	ctx.JSON(http.StatusCreated, newBook)
}

func UpdateBookById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": "The id must be a number",
		})
		return
	}

	isFound := false
	var updatedBook entity.Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": err,
		})
		return
	}

	for idx, book := range Books {
		if book.Id == id {
			Books[idx] = updatedBook
			isFound = true
		}
	}

	if !isFound {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"status":  "Not Found",
			"message": fmt.Sprintf("The book with id = %d is not found in the database", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "Created",
		"message": fmt.Sprintf("Success updating the book with id = %d", id),
	})
}

func DeleteBookById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": "The id must be a number",
		})
		return
	}

	isFound := false
	var bookIdx int

	for idx, book := range Books {
		if book.Id == id {
			bookIdx = idx
			isFound = true
		}
	}

	if !isFound {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"status":  "Not Found",
			"message": fmt.Sprintf("The book with id = %d is not found in the database", id),
		})
		return
	}

	copy(Books[bookIdx:], Books[bookIdx+1:])
	Books[len(Books)-1] = entity.Book{}
	Books = Books[:len(Books)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "Created",
		"message": fmt.Sprintf("Success deleting the book with id = %d", id),
	})
}
