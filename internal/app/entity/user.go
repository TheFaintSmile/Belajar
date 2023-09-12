package entity

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rumbel/belajar/internal/app/utils"
	"github.com/rumbel/belajar/internal/app/utils/token"
	"golang.org/x/crypto/bcrypt"
)

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

type UserLevel string

type User struct {
	gorm.Model
	FirstName string 	`json:"firstname" binding:"required"`
	LastName  string 	`json:"lastname" binding:"required"`
	Age       int8   	`json:"age" binding:"gte=1,lte=130"`
	Email     string 	`json:"email" binding:"required" gorm:"unique"`
	Password  string 	`json:"password" binding:"required,min=6"`
	Level	  UserLevel `json:"level" binding:"required"`
}

func (u *User) SaveUser() (*User, error) {
	validLevels := []UserLevel{LevelSD1, LevelSD2, LevelSD3, LevelSD4, LevelSD5, LevelSD6, LevelSMP, LevelSMA}
    isValidLevel := false
    for _, level := range validLevels {
        if u.Level == level {
            isValidLevel = true
            break
        }
    }
	if !isValidLevel {
        return nil, fmt.Errorf("Invalid user level")
    }
	err := utils.DB.Create(&u).Error	
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
func LoginCheck(email string, password string) (string, error) {
	var err error
	u := User{}
	err = utils.DB.Model(&User{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return "", fmt.Errorf("user not found")
	}
	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", fmt.Errorf("invalid login credentials")
	}
	token, err := token.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}
	return token,nil
}