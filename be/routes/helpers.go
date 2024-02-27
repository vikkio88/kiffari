package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessOr404(ctx *gin.Context, result any, ok bool) {
	if ok {
		ctx.JSON(http.StatusOK, result)
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
}

func SuccessOr400(ctx *gin.Context, result any, ok bool) {
	if ok {
		ctx.JSON(http.StatusOK, result)
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
}

func SuccessOr400WithError(ctx *gin.Context, result any, err error) {
	if err == nil {
		ctx.JSON(http.StatusOK, result)
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
