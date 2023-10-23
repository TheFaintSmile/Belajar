package dto

import models "github.com/rumbel/belajar/internal/app/models"

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
	Category    string            `json:"category" binding:"required"`
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description"`
	Type        models.ModuleType `json:"type" binding:"required"`
	Content     string            `json:"content" binding:"required"`
	WeekID      uint              `json:"week_id" gorm:"foreign_key" binding:"required"`
}
