package models

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Lecturer string `json:"lecturer" binding:"required"`
	LevelID  uint   `json:"level_id" gorm:"foreign_key" binding:"required"`
	// Modules  []string
}
