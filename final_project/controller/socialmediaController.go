package controller

import (
	"mygram/db"
	"mygram/entity"
	"mygram/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllSocialMedia godoc
// @Summary Get details
// @Description Get details of all social media
// @Tags social_media
// @Accept json
// @Produce json
// @Success 302 {object} entity.SocialMedia
// @Router /socialmedia [get]
func GetAllSocialMedia(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusFound
	socialMedia := []entity.SocialMedia{}

	if errFind := db.Find(&socialMedia).Error; errFind != nil {
		status = http.StatusNotFound
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessWithData(status, socialMedia))
}

// GetOneSocialMedia godoc
// @Summary Get details for a given Id
// @Description Get details of social media corresponding to the input Id
// @Tags social_media
// @Accept json
// @Produce json
// @Param id path int true "Id of the social media"
// @Success 302 {object} entity.SocialMedia
// @Router /socialmedia/{id} [get]
func GetOneSocialMedia(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusFound
	socialMedia := entity.SocialMedia{}
	socialMediaId := helper.GetParameterId(ctx)

	if errFind := db.First(&socialMedia, "id = ?", socialMediaId).Error; errFind != nil {
		status = http.StatusNotFound
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessWithData(status, socialMedia))
}

// CreateSocialMedia godoc
// @Summary Post details
// @Description Post details of a new social media
// @Tags social_media
// @Accept json
// @Produce json
// @Param entity.SocialMedia body entity.SocialMedia true "Social media details"
// @Success 201 {object} entity.SocialMedia
// @Router /socialmedia [post]
func CreateSocialMedia(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusCreated
	newSocialMedia := entity.SocialMedia{}
	userId := helper.GetUserIdFromClaims(ctx)

	if errBind := ctx.ShouldBindJSON(&newSocialMedia); errBind != nil {
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	newSocialMedia.UserId = userId

	if errCreate := db.Create(&newSocialMedia).Error; errCreate != nil {
		errs := helper.GetValidationErrorMessage(errCreate)
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorWithData(status, errs))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessNoData(status))
}

// UpdateSocialMedia godoc
// @Summary Update social media identified by the given Id
// @Description Update the social media corresponding to the input Id
// @Tags social_media
// @Accept json
// @Produce json
// @Param id path int true "Id of the social media"
// @Success 200 {object} entity.SocialMedia
// @Router /socialmedia/{id} [put]
func UpdateSocialMedia(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusOK
	newSocialMedia := entity.SocialMedia{}
	userId := helper.GetUserIdFromClaims(ctx)
	socialMediaId := helper.GetParameterId(ctx)

	if errBind := ctx.ShouldBindJSON(&newSocialMedia); errBind != nil {
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	newSocialMedia.Id = uint(socialMediaId)
	newSocialMedia.UserId = userId

	if errUpdate := db.Model(&newSocialMedia).Where("id = ?", socialMediaId).Updates(entity.SocialMedia{Name: newSocialMedia.Name, SocialMediaUrl: newSocialMedia.SocialMediaUrl}).Error; errUpdate != nil {
		status = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessNoData(status))
}

// DeleteSocialMedia godoc
// @Summary Delete social media identified by the given Id
// @Description Delete the social media corresponding to the input Id
// @Tags social_media
// @Accept json
// @Produce json
// @Param id path int true "Id of the social media"
// @Success 200 {object} entity.SocialMedia
// @Router /socialmedia/{id} [delete]
func DeleteSocialMedia(ctx *gin.Context) {
	db := db.GetDB()
	status := http.StatusOK
	socialMediaId := helper.GetParameterId(ctx)

	if errDelete := db.Delete(&entity.SocialMedia{}, socialMediaId).Error; errDelete != nil {
		status = http.StatusNotFound
		ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessNoData(status))
}
