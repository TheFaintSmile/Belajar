package controller

import "github.com/gin-gonic/gin"

type CourseController struct {
}

func NewCourseController() *CourseController {
	return &CourseController{}
}

func (c *CourseController) CourseListController(ctx *gin.Context) (interface{}, error) {
	return "berhasil", nil
}
