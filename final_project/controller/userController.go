package controller

import (
	"mygram/db"
	"mygram/entity"
	"mygram/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserSignup godoc
// @Summary User signup
// @Description User signup for new user
// @Tags users
// @Accept json
// @Produce json
// @Success 201 {object} entity.User
// @Router /users/signup [post]
func UserSignup(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusCreated
	newUser := entity.User{}

	if errBind := ctx.ShouldBindJSON(&newUser); errBind != nil {
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	if errCreate := db.Create(&newUser).Error; errCreate != nil {
		errs := helper.GetValidationErrorMessage(errCreate)
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorWithData(status, errs))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessNoData(status))
}

// UserLogin godoc
// @Summary User login
// @Description User login and get access token
// @Tags users
// @Accept json
// @Produce json
// @Success 302 {object} entity.User
// @Router /users/login [post]
func UserLogin(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusFound
	user := entity.User{}

	if errBind := ctx.ShouldBindJSON(&user); errBind != nil {
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	password := user.Password

	if errFind := db.Where("email = ?", user.Email).Take(&user).Error; errFind != nil {
		status = http.StatusNotFound
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	if isPasswordValid := helper.ComparePassword([]byte(user.Password), []byte(password)); !isPasswordValid {
		status = http.StatusForbidden
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	accessToken := map[string]interface{}{
		"access_token": helper.GenerateToken(user.Id, user.Email),
	}

	ctx.JSON(status, helper.ResponseSuccessWithData(status, accessToken))
}
