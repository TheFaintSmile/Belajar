package utils

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/rumbel/belajar/internal/app/models"
)

func ContainsWhiteSpace(password string) bool {
	return strings.Contains(password, " ")
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isValidModuleType(moduleType models.ModuleType) error {
	validModules := []models.ModuleType{models.ModuleTypeFile, models.ModuleTypeLinks, models.ModuleTypeVideo}

	for _, module := range validModules {
		if moduleType == module {
			return nil
		}
	}
	return fmt.Errorf("invalid user level")
}
