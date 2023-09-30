package models

import "github.com/jinzhu/gorm"

type ModuleType string

const (
    ModuleTypeFile   ModuleType = "file"
    ModuleTypeLinks  ModuleType = "links"
    ModuleTypeVideo  ModuleType = "video"
)

type Course struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Lecturer string `json:"lecturer" binding:"required"`
	LevelID  uint   `json:"level_id" gorm:"foreign_key" binding:"required"`
	Weeks    []Week `json:"weeks" gorm:"foreignKey:CourseID"`
}

type Week struct {
	gorm.Model
	Name     string   `json:"name" binding:"required"`
	CourseID uint     `json:"course_id" gorm:"foreign_key" binding:"required"`
	Materials  []Material `json:"materials" gorm:"foreignKey:WeekID"`
	Tasks    []Task   `json:"tasks" gorm:"foreignKey:WeekID"`
}

type Material struct {
	gorm.Model
	Name    string `json:"name" binding:"required"`
	Type    ModuleType `json:"type" binding:"required"`
	Content string `json:"content" binding:"required"`
	WeekID  uint   `json:"week_id" gorm:"foreign_key" binding:"required"`
}

type Task struct {
	gorm.Model
	Name    string `json:"name" binding:"required"`
	Type    ModuleType `json:"type" binding:"required"`
	Content string `json:"content" binding:"required"`
	WeekID  uint   `json:"week_id" gorm:"foreign_key" binding:"required"`
}
