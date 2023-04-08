package controller

import (
	"net/http"
	"todolist/database"
	"todolist/entity"
	"todolist/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	newTodo := entity.Todo{}
	status := http.StatusCreated

	errBind := ctx.ShouldBindJSON(&newTodo)
	if errBind != nil {
		status := http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseError(status))
		return
	}

	newTodo.UserId = userId

	errCreate := db.Create(&newTodo).Error
	if errCreate != nil {
		status = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(status, helper.ResponseError(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessNoData(status))
}
