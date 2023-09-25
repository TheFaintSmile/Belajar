package models

import (
	"github.com/jinzhu/gorm"
)

type UserRole string

const (
	RoleSiswa      UserRole = "Siswa"
	RolePendidik   UserRole = "Pendidik"
	RoleAdmin      UserRole = "Admin"
)

type User struct {
	gorm.Model
	FirstName string    `json:"firstname" binding:"required"`
	LastName  string    `json:"lastname" binding:"required"`
	Age       int8      `json:"age" binding:"gte=1,lte=130"`
	Email     string    `json:"email" binding:"required" gorm:"unique"`
	Password  string    `json:"password" binding:"required,min=6"`
	LevelID   UserLevel `json:"level_id"`
	Role      UserRole   `json:"role" binding:"required" gorm:"index"`
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
