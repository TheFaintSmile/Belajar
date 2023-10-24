package utils

import (
	"errors"

	"github.com/rumbel/belajar/internal/app/dto"
	"github.com/rumbel/belajar/internal/app/models"
)

func IsValidCategory(module *dto.AddModuleToCourse) error {
	validCategory := []dto.Category{dto.CategoryMaterial, dto.CategoryTask}

	for _, category := range validCategory {
		if module.Category == category {
			return nil
		}
	}
	return errors.New("invalid category")
}

func IsValidModuleType(module *dto.AddModuleToCourse) error {
	validModuleType := []models.ModuleType{models.ModuleTypeVideo, models.ModuleTypeLinks, models.ModuleTypeFile}

	for _, moduleType := range validModuleType {
		if module.Type == moduleType {
			return nil
		}
	}
	return errors.New("invalid module type")
}