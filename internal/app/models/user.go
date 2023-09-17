package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	FirstName string    `json:"firstname" binding:"required"`
	LastName  string    `json:"lastname" binding:"required"`
	Age       int8      `json:"age" binding:"gte=1,lte=130"`
	Email     string    `json:"email" binding:"required" gorm:"unique"`
	Password  string    `json:"password" binding:"required,min=6"`
	LevelID   UserLevel `json:"level_id" gorm:"index"`
	Level     Level     `json:"level" gorm:"foreignkey:LevelID"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
