package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rumbel/belajar/internal/app/dto"
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

func (c *CourseController) GetCourseList(ctx *gin.Context, userLevel int) ([]dto.CourseListResponse, error) {
	courses, err := c.service.GetCourseList(userLevel)

	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *CourseController) GetCourseDetail(ctx *gin.Context) (models.Course, error) {
	courseID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		return models.Course{}, err
	}

	course, err := c.service.GetCourseDetail(uint(courseID))

	if err != nil {
		return models.Course{}, err
	}

	return course, nil
}

func (c *CourseController) AddCourse(ctx *gin.Context) (dto.AddCourseInput, error) {
	var course dto.AddCourseInput

	if err := ctx.ShouldBindJSON(&course); err != nil {
		return dto.AddCourseInput{}, err
	}

	result, err := c.service.AddCourse(course)

	if err != nil {
		return dto.AddCourseInput{}, err
	}

	return result, nil
}

func (c *CourseController) AddWeekToCourse(ctx *gin.Context) (models.Week, error) {
	var week models.Week

	if err := ctx.ShouldBindJSON(&week); err != nil {
		return models.Week{}, err
	}

	occurrence, err := c.service.GetWeekOccurrence(week.CourseID)

	if err != nil {
		return models.Week{}, err
	}

	week.WeekNumber = occurrence + 1

	result, err := c.service.AddWeekToCourse(week)

	if err != nil {
		return models.Week{}, err
	}

	return result, nil
}

func (c *CourseController) UpdateCourseInformation(ctx *gin.Context) (dto.UpdateCourseInformationInput, error) {
	courseID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		return dto.UpdateCourseInformationInput{}, err
	}

	var course dto.UpdateCourseInformationInput

	if err := ctx.ShouldBindJSON(&course); err != nil {
		return dto.UpdateCourseInformationInput{}, err
	}

	result, err := c.service.UpdateCourseInformation(uint(courseID), course)

	if err != nil {
		return dto.UpdateCourseInformationInput{}, err
	}

	return result, nil
}

func (c *CourseController) DeleteCourse(ctx *gin.Context) error {
	courseID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		return err
	}

	err = c.service.DeleteCourse(uint(courseID))

	if err != nil {
		return err
	}

	return nil
}

func (c *CourseController) UpdateWeekInCourse(ctx *gin.Context) (dto.UpdateWeekInCourseInput, error) {
	courseID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		return dto.UpdateWeekInCourseInput{}, err
	}

	weekID, err := strconv.ParseUint(ctx.Param("weekID"), 10, 32)

	if err != nil {
		return dto.UpdateWeekInCourseInput{}, err
	}

	var week dto.UpdateWeekInCourseInput

	if err := ctx.ShouldBindJSON(&week); err != nil {
		return dto.UpdateWeekInCourseInput{}, err
	}

	result, err := c.service.UpdateWeekInCourse(uint(courseID), uint(weekID), week)

	if err != nil {
		return dto.UpdateWeekInCourseInput{}, err
	}

	return result, nil
}

func (c *CourseController) DeleteWeekInCourse(ctx *gin.Context) error {
	courseID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		return err
	}

	weekID, err := strconv.ParseUint(ctx.Param("weekID"), 10, 32)

	if err != nil {
		return err
	}

	err = c.service.DeleteWeekInCourse(uint(courseID), uint(weekID))

	if err != nil {
		return err
	}

	return nil
}

func (c *CourseController) AddModuleToCourse(ctx *gin.Context) (dto.AddModuleToCourse, error) {
	courseID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		return dto.AddModuleToCourse{}, err
	}

	weekID, err := strconv.ParseUint(ctx.Param("weekID"), 10, 32)

	if err != nil {
		return dto.AddModuleToCourse{}, err
	}

	fmt.Println(courseID)
	fmt.Println(weekID)

	module := dto.AddModuleToCourse{
		Category:    dto.Category(ctx.PostForm("category")),
		Name:        ctx.PostForm("name"),
		Type:        models.ModuleType(ctx.PostForm("type")),
		Content:     ctx.PostForm("content"),
		Description: ctx.PostForm("description"),
	}
	fmt.Println(module)

	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		return dto.AddModuleToCourse{}, err
	}

	module.File = file

	result, err := c.service.AddModuleToCourse(uint(courseID), uint(weekID), module)

	if err != nil {
		return dto.AddModuleToCourse{}, err
	}

	return result, nil
}

func (c *CourseController) DeleteMaterialFromCourse(ctx *gin.Context) error {
	materialID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		return err
	}

	err = c.service.DeleteMaterialFromCourse(uint(materialID))

	if err != nil {
		return err
	}

	return nil
}

func (c *CourseController) DeleteTaskFromCourse(ctx *gin.Context) error {
	taskID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		return err
	}

	err = c.service.DeleteTaskFromCourse(uint(taskID))

	if err != nil {
		return err
	}

	return nil
}