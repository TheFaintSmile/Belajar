package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx *gin.Context, message string, data interface{}) {
	response := Response{
		Message: message,
		Status:  "SUCCESS",
		Data:    data,
	}
	ctx.JSON(http.StatusOK, response)
}

func ErrorResponse(ctx *gin.Context, message string, data interface{}) {
	response := Response{
		Message: message,
		Status:  "FAILED",
		Data:    data,
	}
	ctx.JSON(http.StatusBadRequest, response)
}