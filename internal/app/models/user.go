package models

import (
	"github.com/jinzhu/gorm"
)

type UserLevel string
type UserRole string

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

const (
	RoleSiswa      UserRole = "Siswa"
	RolePendidik   UserRole = "Pendidik"
	RoleAdmin      UserRole = "Admin"
)

type User struct {
	gorm.Model
	FirstName string     `json:"firstname" binding:"required"`
	LastName  string     `json:"lastname" binding:"required"`
	Age       int8       `json:"age" binding:"gte=1,lte=130"`
	Email     string     `json:"email" binding:"required" gorm:"unique"`
	Password  string     `json:"password" binding:"required,min=6"`
	Level     UserLevel  `json:"level"`
	Role      UserRole   `json:"role" binding:"required"`
}

type Siswa struct {
	gorm.Model
	Email    string    `json:"email" binding:"required" gorm:"unique"`
}

type Pendidik struct {
	gorm.Model
	Email    string   `json:"email" binding:"required" gorm:"unique"`
}

type Admin struct {
	gorm.Model
	Email    string  `json:"email" binding:"required" gorm:"unique"`
}
