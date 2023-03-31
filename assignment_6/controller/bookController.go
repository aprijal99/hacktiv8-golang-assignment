package controller

import (
	"book-service/database"
	"book-service/entity"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(ctx *gin.Context) {
	db := database.GetDB()
	var books []entity.Book

	rows, err := db.Query("select * from books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book entity.Book

		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			panic(err)
		}

		books = append(books, book)
	}

	ctx.JSON(http.StatusOK, books)
}

func GetBookById(ctx *gin.Context) {
	db := database.GetDB()
	var book entity.Book

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	row := db.QueryRow("select * from books where id=$1", id)
	errScan := row.Scan(&book.Id, &book.Title, &book.Author, &book.Desc)
	if errScan != nil {
		panic(errScan)
	}

	ctx.JSON(http.StatusOK, book)
}

func CreateNewBook(ctx *gin.Context) {
	db := database.GetDB()
	var newBook entity.Book

	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		panic(err)
	}

	errInsert := db.QueryRow("insert into books (title, author, description) values ($1, $2, $3)", &newBook.Title, &newBook.Author, &newBook.Desc)
	if errInsert != nil {
		panic(err)
	}

	ctx.JSON(http.StatusCreated, "Created")
}

func UpdateBook(ctx *gin.Context) {
	db := database.GetDB()
	var updatedBook entity.Book

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	errBind := ctx.ShouldBindJSON(&updatedBook)
	if errBind != nil {
		panic(err)
	}

	res, errUpdate := db.Exec("update books set title=$2, author=$3, description=$4 where id=$1", id, &updatedBook.Title, &updatedBook.Author, &updatedBook.Desc)
	if errUpdate != nil {
		panic(errUpdate)
	}

	fmt.Printf("Updated %d row(s)", res)

	ctx.JSON(http.StatusOK, "Updated")
}

func DeleteBook(ctx *gin.Context) {
	db := database.GetDB()

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	res, errDelete := db.Exec("delete from books where id=$1", id)
	if errDelete != nil {
		panic(errDelete)
	}

	fmt.Printf("Deleted %d row(s)", res)

	ctx.JSON(http.StatusOK, "Deleted")
}
