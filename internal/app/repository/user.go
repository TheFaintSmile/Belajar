package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	middlewares "github.com/rumbel/belajar/internal/app/middlewares"
	"github.com/rumbel/belajar/internal/app/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) SaveUser(user *models.User) (*models.User, error) {
	err := ur.DB.Create(user).Error
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (ur *UserRepository) VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (ur *UserRepository) LoginCheck(email, password string) (string, error) {
	var err error
	user := models.User{}
	err = ur.DB.Model(&models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", fmt.Errorf("user not found")
	}
	err = ur.VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", fmt.Errorf("invalid login credentials")
	}
	token, err := middlewares.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (ur *UserRepository) CheckEmailExists(email string) bool {
	var u models.User
	err := ur.DB.Model(&models.User{}).Where("email = ?", email).Take(&u).Error
	return err == nil
}
