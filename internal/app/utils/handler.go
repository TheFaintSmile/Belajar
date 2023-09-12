package utils

import (
	"github.com/gin-gonic/gin"
)

// Response is a custom struct to format responses consistently
type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

// sendResponse sends a response with a standardized format
func SendResponse(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(status, Response{
		Message: message,
		Status:  status,
		Data:    data,
	})
}