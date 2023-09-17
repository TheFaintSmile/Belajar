package dto

import "github.com/jinzhu/gorm"

type CourseList struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Lecturer string `json:"lecturer" binding:"required"`
}

