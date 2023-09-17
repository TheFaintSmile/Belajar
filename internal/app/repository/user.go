package repository

import (
	"fmt"

	middlewares "github.com/rumbel/belajar/internal/app/middlewares"
	"github.com/rumbel/belajar/internal/app/models"
	"github.com/rumbel/belajar/internal/app/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) SaveUser(user *models.User) (*models.User, error) {
	validLevels := []models.UserLevel{models.LevelSD1, models.LevelSD2, models.LevelSD3, models.LevelSD4, models.LevelSD5, models.LevelSD6, models.LevelSMP, models.LevelSMA}
	isValidLevel := false
	
	for _, level := range validLevels {
		if user.LevelID == level {
			isValidLevel = true
			break
		}
	}
	if !isValidLevel {
		return nil, fmt.Errorf("Invalid user level")
	}
	err := utils.DB.Create(user).Error
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
	err = utils.DB.Model(&models.User{}).Where("email = ?", email).Take(&user).Error
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
	err := utils.DB.Model(&models.User{}).Where("email = ?", email).Take(&u).Error
	return err == nil
}
