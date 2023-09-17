package models

import "github.com/jinzhu/gorm"

type CourseList struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Lecturer string `json:"lecturer" binding:"required"`
	LevelID  string `json:"level_id" gorm:"index"`
	Level    Level  `json:"level" gorm:"foreignkey:LevelID"`
}
