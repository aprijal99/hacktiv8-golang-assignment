package middleware

import (
	"mygram/db"
	"mygram/entity"
	"mygram/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizePhoto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := db.GetDB()
		photoId := helper.GetParameterId(ctx)
		userId := helper.GetUserIdFromClaims(ctx)
		savedPhoto := entity.Photo{}

		if errFind := db.First(&savedPhoto, "id = ?", photoId).Error; errFind != nil {
			status := http.StatusNotFound
			ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
			return
		}

		if userId != savedPhoto.UserId {
			status := http.StatusUnauthorized
			ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
			return
		}

		ctx.Next()
	}
}

func AuthorizeSocialMedia() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := db.GetDB()
		socialMediaId := helper.GetParameterId(ctx)
		userId := helper.GetUserIdFromClaims(ctx)
		savedSocialMedia := entity.SocialMedia{}

		if errFind := db.First(&savedSocialMedia, "id = ?", socialMediaId).Error; errFind != nil {
			status := http.StatusNotFound
			ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
			return
		}

		if userId != savedSocialMedia.UserId {
			status := http.StatusUnauthorized
			ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
			return
		}

		ctx.Next()
	}
}

func AuthorizeComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := db.GetDB()
		commentId := helper.GetParameterId(ctx)
		userId := helper.GetUserIdFromClaims(ctx)
		savedComment := entity.Comment{}

		if errFind := db.First(&savedComment, "id = ?", commentId).Error; errFind != nil {
			status := http.StatusNotFound
			ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
			return
		}

		if userId != savedComment.UserId {
			status := http.StatusUnauthorized
			ctx.AbortWithStatusJSON(status, helper.ResponseErrorNoData(status))
			return
		}

		ctx.Next()
	}
}
