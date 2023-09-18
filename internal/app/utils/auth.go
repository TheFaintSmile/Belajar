package utils

import (
	"fmt"

	"github.com/rumbel/belajar/internal/app/models"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func IsValidLevel(user *models.User) error {
	validLevels := []models.UserLevel{models.LevelSD1, models.LevelSD2, models.LevelSD3, models.LevelSD4, models.LevelSD5, models.LevelSD6, models.LevelSMP, models.LevelSMA}
	for _, level := range validLevels {
		if user.Level == level {
			return nil
		}
	}
	return fmt.Errorf("invalid user level")
}

func HashingPassword(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}