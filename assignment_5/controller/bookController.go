package controller

import (
	"book/entity"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Books = []entity.Book{
	{Id: 1, Title: "Calculus 1", Author: "John F Bach", Desc: "Calculus 1 covers mostly about Derivatives"},
	{Id: 2, Title: "Introduction of Biochemistry", Author: "Lehninger", Desc: "This is introduction to biomolecules"},
}

func GetAllBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Books)
}

func GetBookById(ctx *gin.Context) {
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
	var bookData entity.Book

	for _, book := range Books {
		if book.Id == id {
			bookData = book
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

	ctx.JSON(http.StatusFound, bookData)
}

func CreateNewBook(ctx *gin.Context) {
	var newBook entity.Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": err,
		})
		return
	}

	newId := len(Books) + 1
	newBook.Id = newId
	Books = append(Books, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"status":  "Created",
		"message": fmt.Sprintf("Success creating new book with id = %d", newId),
	})
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
