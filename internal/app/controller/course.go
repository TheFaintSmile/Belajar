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

func (c *CourseController) GetCourseList(ctx *gin.Context) ([]models.Course, error) {
	courses, err := c.service.GetCourseList()

	if err != nil {
		return []models.Course{}, err
	}

	return courses, nil
}
