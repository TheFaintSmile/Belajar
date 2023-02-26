package controller

import "github.com/gin-gonic/gin"

type TestControllerImpl struct {
	text string
}

func NewTestControllerImpl() TestController {
	return &TestControllerImpl{
		text: "asdf",
	}
}

func (controller *TestControllerImpl) GetTests(c *gin.Context) {
	c.JSON(200, "Hello World")
}
