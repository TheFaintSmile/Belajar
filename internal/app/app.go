package app

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rumbel/belajar/internal/app/routes"
	"github.com/rumbel/belajar/internal/config"
	// "github.com/jinzhu/gorm"
)

type App struct {
	api    *gin.RouterGroup
	config *config.Config
	// db     *gorm.DB
	router *gin.Engine
}

func NewApp() *App {
	config := config.LoadConfig()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("\nError loading config: \n", err)
	}

	if os.Getenv("NODE_ENV") == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// db, err := gorm.Open(config.Database.Dialect, config.Database.URL)

	router := gin.Default()
	router.Use(gin.Recovery())
	api := router.Group("/api")

	return &App{
		api:    api,
		config: config,
		// db:     db,
		router: router,
	}
}

// Start App
func (a *App) Run() {
	a.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))
	a.router.SetTrustedProxies(nil)
	serverPort := fmt.Sprintf(":%s", a.config.ServerPort)
	routes.TestRoutes(a.api)

	a.router.Run(serverPort)
}

// func (a *App) Close() {
// 	a.db.Close()
// }