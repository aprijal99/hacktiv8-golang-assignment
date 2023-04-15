package controller

import (
	"mygram/db"
	"mygram/entity"
	"mygram/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllComments(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusFound
	comments := []entity.Comment{}

	if errFind := db.Find(&comments).Error; errFind != nil {
		status = http.StatusNotFound
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessWithData(status, comments))
}

func GetOneComment(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusFound
	comment := entity.Comment{}
	commentId := helper.GetParameterId(ctx)

	if errFind := db.First(&comment, "id = ?", commentId).Error; errFind != nil {
		status = http.StatusNotFound
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessWithData(status, comment))
}

func CreateComment(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusCreated
	newComment := entity.Comment{}
	userId := helper.GetUserIdFromClaims(ctx)

	if errBind := ctx.ShouldBindJSON(&newComment); errBind != nil {
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	newComment.UserId = userId

	if errCreate := db.Create(&newComment).Error; errCreate != nil {
		errs := helper.GetValidationErrorMessage(errCreate)
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorWithData(status, errs))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessNoData(status))
}

func UpdateComment(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusOK
	newComment := entity.Comment{}
	userId := helper.GetUserIdFromClaims(ctx)
	commentId := helper.GetParameterId(ctx)

	if errBind := ctx.ShouldBindJSON(&newComment); errBind != nil {
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	newComment.Id = uint(commentId)
	newComment.UserId = userId

	if errUpdate := db.Model(&newComment).Where("id = ?", newComment).Updates(entity.Comment{Message: newComment.Message}).Error; errUpdate != nil {
		status = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessNoData(status))
}

func DeleteComment(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusOK
	commentId := helper.GetParameterId(ctx)

	if errDelete := db.Delete(&entity.Photo{}, commentId).Error; errDelete != nil {
		status = http.StatusNotFound
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessNoData(status))
}
