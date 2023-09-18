package models

import (
	"github.com/jinzhu/gorm"
)

type UserLevel string

const (
	LevelSD1 UserLevel = "SD-1"
	LevelSD2 UserLevel = "SD-2"
	LevelSD3 UserLevel = "SD-3"
	LevelSD4 UserLevel = "SD-4"
	LevelSD5 UserLevel = "SD-5"
	LevelSD6 UserLevel = "SD-6"
	LevelSMP UserLevel = "SMP"
	LevelSMA UserLevel = "SMA"
)

type Level struct {
	gorm.Model
	Name    UserLevel `json:"name" gorm:"primary_key"`
	Courses []Course  `gorm:"foreignKey:LevelID"`
	Users   []User    `gorm:"foreignKey:LevelID"`
}
