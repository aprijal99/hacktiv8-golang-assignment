package middleware

import (
	"net/http"
	"todolist/helper"

	"github.com/gin-gonic/gin"
)

func AuthorizeRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, err := helper.VerifyToken(ctx)
		if err != nil {
			status := http.StatusUnauthorized
			ctx.AbortWithStatusJSON(status, helper.ResponseError(status))
		}

		ctx.Set("userData", userData)
		ctx.Next()
	}
}
