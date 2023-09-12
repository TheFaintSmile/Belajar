package models

import (
	"github.com/jinzhu/gorm"
)

type UserLevel string

const (
	LevelSD1      UserLevel = "SD-1"
	LevelSD2      UserLevel = "SD-2"
	LevelSD3 	  UserLevel = "SD-3"
	LevelSD4      UserLevel = "SD-4"
	LevelSD5      UserLevel = "SD-5"
	LevelSD6      UserLevel = "SD-6"
	LevelSMP      UserLevel = "SMP"
	LevelSMA      UserLevel = "SMA"
)

type User struct {
	gorm.Model
	FirstName string     `json:"firstname" binding:"required"`
	LastName  string     `json:"lastname" binding:"required"`
	Age       int8       `json:"age" binding:"gte=1,lte=130"`
	Email     string     `json:"email" binding:"required" gorm:"unique"`
	Password  string     `json:"password" binding:"required,min=6"`
	Level     UserLevel  `json:"level" binding:"required"`
}
