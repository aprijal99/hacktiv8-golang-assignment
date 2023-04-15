package helper

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParameterId(ctx *gin.Context) int {
	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		status := http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, ResponseErrorNoData(status))
	}

	return id
}
