package controller

import (
	"mygram/db"
	"mygram/entity"
	"mygram/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllComments godoc
// @Summary Get details
// @Description Get details of all comments
// @Tags comments
// @Accept json
// @Produce json
// @Success 302 {object} entity.Comment
// @Router /comments [get]
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

// GetOneComment godoc
// @Summary Get details for a given Id
// @Description Get details of comment corresponding to the input Id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Id of the comment"
// @Success 302 {object} entity.Comment
// @Router /comments/{id} [get]
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

// CreateComment godoc
// @Summary Post details
// @Description Post details of a new comment
// @Tags comments
// @Accept json
// @Produce json
// @Param entity.Comment body entity.Comment true "Comment details"
// @Success 201 {object} entity.Comment
// @Router /comments [post]
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

// UpdateComment godoc
// @Summary Update comment identified by the given Id
// @Description Update the comment corresponding to the input Id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Id of the comment"
// @Success 200 {object} entity.Comment
// @Router /comments/{id} [put]
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

// DeleteComment godoc
// @Summary Delete comment identified by the given Id
// @Description Delete the comment corresponding to the input Id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Id of the comment"
// @Success 200 {object} entity.Comment
// @Router /comments/{id} [delete]
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
