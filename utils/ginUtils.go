package utils

import (
	"github.com/gin-gonic/gin"
	"ms-product/model"
	"net/http"
)

func ResponseError(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, model.ErrorResponse{
		code, err.Error(),
	})
}

func ResponseResult(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func ResponseCreated(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{"data": data})

}
