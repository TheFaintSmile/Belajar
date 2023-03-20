package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort       string
	DatabaseDriver   string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
}

func NewConfig() *Config {
	// Default value
	return &Config{
		ServerPort:       "8080",
		DatabaseDriver:   "psql",
		DatabaseHost:     "postgres",
		DatabasePort:     "5432",
		DatabaseName:     "rumbel",
		DatabaseUser:     "postgres",
		DatabasePassword: "",
	}
}

func LoadConfig() *Config {
	cfg := NewConfig()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("\nError loading config: \n", err)
	}

	if port := os.Getenv("SERVER_PORT"); port != "" {
		cfg.ServerPort = port
	}

	if driver := os.Getenv("DATABASE_DRIVER"); driver != "" {
		cfg.DatabaseDriver = driver
	}

	if host := os.Getenv("DATABASE_HOST"); host != "" {
		cfg.DatabaseHost = host
	}

	if port := os.Getenv("DATABASE_PORT"); port != "" {
		cfg.DatabasePort = port
	}

	if name := os.Getenv("DATABASE_NAME"); name != "" {
		cfg.DatabaseName = name
	}

	if user := os.Getenv("DATABASE_USER"); user != "" {
		cfg.DatabaseUser = user
	}

	if password := os.Getenv("DATABASE_PASSWORD"); password != "" {
		cfg.DatabasePassword = password
	}

	return cfg
}

func (cfg *Config) GetDSN() string {
	return cfg.DatabaseUser + ":" + cfg.DatabasePassword + "@tcp(" + cfg.DatabaseHost + ":" + cfg.DatabasePort + ")/" + cfg.DatabaseName
}

func (cfg *Config) GetServerPort() int {
	port, err := strconv.Atoi(cfg.ServerPort)
	if err != nil {
		panic("Invalid server port: " + cfg.ServerPort)
	}
	return port
}
