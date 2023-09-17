package utils

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/rumbel/belajar/internal/config"
)

var DB *gorm.DB

func ConnectDB() {
	config := config.LoadConfig()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DB, err = gorm.Open(config.DatabaseDriver, config.GetDSN())
	if err != nil {
		fmt.Println("Failed connect to database ", err)
	} else {
		fmt.Println("Connected to database", config.DatabaseDriver)
	}
}
