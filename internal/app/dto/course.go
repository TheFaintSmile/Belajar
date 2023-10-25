package dto

import (
	"mime/multipart"

	models "github.com/rumbel/belajar/internal/app/models"
)

type Category string

const (
	CategoryMaterial Category = "material"
	CategoryTask     Category = "task"
)

type (
	CourseListResponse struct {
		ID       uint   `json:"id" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Lecturer string `json:"lecturer" validate:"required"`
	}
)

type AddCourseInput struct {
	Name     string `json:"name" binding:"required"`
	Lecturer string `json:"lecturer" binding:"required"`
	LevelID  uint   `json:"level_id" gorm:"foreign_key" binding:"required"`
}

type UpdateCourseInformationInput struct {
	Name     string `json:"name" binding:"required"`
	Lecturer string `json:"lecturer" binding:"required"`
	LevelID  uint   `json:"level_id" gorm:"foreign_key" binding:"required"`
}

type UpdateWeekInCourseInput struct {
	Name string `json:"name" binding:"required"`
}

type AddModuleToCourse struct {
	Category    Category          `form:"category" binding:"required"`
	Name        string            `form:"name" binding:"required"`
	Description string            `form:"description"`
	Type        models.ModuleType `form:"type" binding:"required"`
	Content     string            `form:"content"`
	File        multipart.File    `form:"file"`
}
