package middleware

import (
	"mygram/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func AuthenticateRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, err := helper.VerifyToken(ctx)
		if err != nil {
			status := http.StatusUnauthorized
			ctx.AbortWithStatusJSON(status, helper.ResponseErrorWithData(status, cases.Title(language.AmericanEnglish).String(err.Error())))
			return
		}

		ctx.Set("userData", userData)
		ctx.Next()
	}
}
