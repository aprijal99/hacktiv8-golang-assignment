package controller

import (
	"net/http"
	"product/database"
	"product/entity"
	"product/helper"

	"github.com/gin-gonic/gin"
)

func GetProductsOwnedByUserId(ctx *gin.Context) {
	db := database.GetDB()
	userId := helper.GetUserIdFromClaims(ctx)
	status := http.StatusFound
	products := []entity.Product{}

	err := db.Where("user_id = ?", userId).Find(&products).Error
	if err != nil {
		status = http.StatusNotFound
		ctx.AbortWithStatusJSON(status, helper.ResponseError(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessWithData(status, products))
}

func GetProductById(ctx *gin.Context) {

}

func CreateNewProduct(ctx *gin.Context) {
	db := database.GetDB()
	userId := helper.GetUserIdFromClaims(ctx)
	status := http.StatusCreated
	newProduct := entity.Product{}

	errBind := ctx.ShouldBindJSON(&newProduct)
	if errBind != nil {
		status := http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseError(status))
		return
	}

	newProduct.UserId = userId

	errCreate := db.Create(&newProduct).Error
	if errCreate != nil {
		status = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(status, helper.ResponseError(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessNoData(status))
}
