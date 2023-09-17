package app

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/rumbel/belajar/internal/app/middlewares"
	"github.com/rumbel/belajar/internal/app/models"
	"github.com/rumbel/belajar/internal/app/routes"
	"github.com/rumbel/belajar/internal/app/utils"
	"github.com/rumbel/belajar/internal/config"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	api    *gin.RouterGroup
	config *config.Config
	db     *gorm.DB
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

	// db, err := gorm.Open(config.DatabaseDriver, config.GetDSN())
	// if err != nil {
	//     log.Fatal("Error connecting to the database:", err)
	// }
	// db.AutoMigrate(&entity.User{})
	utils.ConnectDB()
	utils.DB.AutoMigrate(&models.User{})
	utils.DB.AutoMigrate(&models.Level{})
	utils.DB.AutoMigrate(&models.CourseList{})

	middlewares.InitializeLevelToDatabase(utils.DB)
	// Serve Swagger documentation
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middlewares.LogRequest)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")

	return &App{
		api:    api,
		config: config,
		db:     utils.DB,
		router: router,
	}
}

func (a *App) Run() {
	a.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))
	api := a.api
	db := a.db

	a.router.SetTrustedProxies(nil)
	serverPort := fmt.Sprintf(":%s", a.config.ServerPort)
	routes.AuthRoutes(api, db)
	routes.CourseRoutes(api, db)

	a.router.Run(serverPort)
}
