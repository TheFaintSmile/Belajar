package models

import (
	"github.com/jinzhu/gorm"
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
