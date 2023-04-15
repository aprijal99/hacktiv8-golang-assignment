package controller

import (
	"mygram/db"
	"mygram/entity"
	"mygram/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllPhotos godoc
// @Summary Get details
// @Description Get details of all photos
// @Tags photos
// @Accept json
// @Produce json
// @Success 302 {object} entity.Photo
// @Router /photos [get]
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

// GetOnePhoto godoc
// @Summary Get details for a given Id
// @Description Get details of photo corresponding to the input Id
// @Tags photos
// @Accept json
// @Produce json
// @Param id path int true "Id of the photo"
// @Success 302 {object} entity.Photo
// @Router /photos/{id} [get]
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

// CreatePhoto godoc
// @Summary Post details
// @Description Post details of a new photo
// @Tags photos
// @Accept json
// @Produce json
// @Param entity.Photo body entity.Photo true "Photo details"
// @Success 201 {object} entity.Photo
// @Router /photos [post]
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

// UpdatePhoto godoc
// @Summary Update photo identified by the given Id
// @Description Update the photo corresponding to the input Id
// @Tags photos
// @Accept json
// @Produce json
// @Param id path int true "Id of the photo"
// @Success 200 {object} entity.Photo
// @Router /photos/{id} [put]
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

// DeletePhoto godoc
// @Summary Delete photo identified by the given Id
// @Description Delete the photo corresponding to the input Id
// @Tags photos
// @Accept json
// @Produce json
// @Param id path int true "Id of the photo"
// @Success 200 {object} entity.Photo
// @Router /photos/{id} [delete]
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
