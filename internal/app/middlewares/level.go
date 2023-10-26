package middlewares

import (
	"github.com/jinzhu/gorm"
	"github.com/rumbel/belajar/internal/app/models"
)

func InitializeLevelToDatabase(db *gorm.DB) {
	validLevels := []models.UserLevel{models.LevelSD1, models.LevelSD2, models.LevelSD3, models.LevelSD4, models.LevelSD5, models.LevelSD6, models.LevelSMP, models.LevelSMA}

	for _, levelName := range validLevels {
		var existingLevel models.Level

		// Check if the level already exists in the database
		result := db.Where("name = ?", levelName).First(&existingLevel)
		if result.Error != nil {
			// Level doesn't exist, so create it
			newLevel := models.Level{Name: levelName}
			db.Create(&newLevel)
		}
	}
}
