package controller

import "github.com/gin-gonic/gin"

type TestController interface {
	GetTests(c *gin.Context)
}
