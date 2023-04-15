package controller

import (
	"mygram/db"
	"mygram/entity"
	"mygram/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPhotos(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusFound
	photos := []entity.Photo{}

	if errFind := db.Find(&photos).Error; errFind != nil {
		status = http.StatusNotFound
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessWithData(status, photos))
}

func GetOnePhoto(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusFound
	photo := entity.Photo{}
	photoId := helper.GetParameterId(ctx)

	if errFind := db.First(&photo, "id = ?", photoId).Error; errFind != nil {
		status = http.StatusNotFound
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessWithData(status, photo))
}

func CreatePhoto(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusCreated
	newPhoto := entity.Photo{}
	userId := helper.GetUserIdFromClaims(ctx)

	if errBind := ctx.ShouldBindJSON(&newPhoto); errBind != nil {
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	newPhoto.UserId = userId

	if errCreate := db.Create(&newPhoto).Error; errCreate != nil {
		errs := helper.GetValidationErrorMessage(errCreate)
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorWithData(status, errs))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessNoData(status))
}

func UpdatePhoto(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusOK
	newPhoto := entity.Photo{}
	userId := helper.GetUserIdFromClaims(ctx)
	photoId := helper.GetParameterId(ctx)

	if errBind := ctx.ShouldBindJSON(&newPhoto); errBind != nil {
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	newPhoto.Id = uint(photoId)
	newPhoto.UserId = userId

	if errUpdate := db.Model(&newPhoto).Where("id = ?", photoId).Updates(entity.Photo{Title: newPhoto.Title, Caption: newPhoto.Caption}).Error; errUpdate != nil {
		status = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessNoData(status))
}

func DeletePhoto(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusOK
	photoId := helper.GetParameterId(ctx)

	if errDelete := db.Delete(&entity.Photo{}, photoId).Error; errDelete != nil {
		status = http.StatusNotFound
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessNoData(status))
}
