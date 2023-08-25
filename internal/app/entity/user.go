package entity

import (
	"github.com/jinzhu/gorm"
	"github.com/rumbel/belajar/internal/app/utils"
	"github.com/rumbel/belajar/internal/app/utils/token"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

func (u *User) SaveUser() (*User, error) {
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
		return "", err
	}
	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := token.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}
	return token,nil
}