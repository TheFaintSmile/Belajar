package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rumbel/belajar/internal/app/models"
	"github.com/rumbel/belajar/internal/app/service"
)

type CourseController struct {
	service service.CourseService
}

func NewCourseController(service service.CourseService) *CourseController {
	return &CourseController{
		service: service,
	}
}

func (c *CourseController) GetCourseList(ctx *gin.Context, userLevel int) ([]models.Course, error) {
	courses, err := c.service.GetCourseList(userLevel)

	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *CourseController) AddCourse(ctx *gin.Context) (models.Course, error) {
	var course models.Course

	if err := ctx.ShouldBindJSON(&course); err != nil {
		return models.Course{}, err
	}

	result, err := c.service.AddCourse(course)

	if err != nil {
		return models.Course{}, err
	}

	return result, nil
}
