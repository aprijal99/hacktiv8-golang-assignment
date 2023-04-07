package controller

import (
	"net/http"
	"todolist/database"
	"todolist/entity"
	"todolist/helper"

	"github.com/gin-gonic/gin"
)

func UserSignup(ctx *gin.Context) {
	db := database.GetDB()
	newUser := entity.User{}
	status := http.StatusCreated

	errBind := ctx.ShouldBindJSON(&newUser)
	if errBind != nil {
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, entity.ResponseError(status))
		return
	}

	errCreate := db.Create(&newUser).Error
	if errCreate != nil {
		status = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(status, entity.ResponseError(status))
		return
	}

	ctx.JSON(status, entity.ResponseSuccessNoData(status))
}

func UserLogin(ctx *gin.Context) {
	db := database.GetDB()
	user := entity.User{}
	status := http.StatusFound

	errBind := ctx.ShouldBindJSON(&user)
	if errBind != nil {
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, entity.ResponseError(status))
		return
	}

	password := user.Password

	errFind := db.Where("email = ?", user.Email).Take(&user).Error
	if errFind != nil {
		status = http.StatusNotFound
		ctx.AbortWithStatusJSON(status, entity.ResponseError(status))
		return
	}

	isPasswordValid := helper.ComparePassword([]byte(user.Password), []byte(password))
	if !isPasswordValid {
		status = http.StatusForbidden
		ctx.AbortWithStatusJSON(status, entity.ResponseError(status))
		return
	}

	accessToken := map[string]interface{}{
		"access_token": helper.GenerateToken(user.Id, user.Email),
	}

	ctx.JSON(status, entity.ResponseSuccessWithData(status, accessToken))
}
