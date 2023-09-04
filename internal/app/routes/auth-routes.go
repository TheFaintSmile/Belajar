package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	controller "github.com/rumbel/belajar/internal/app/controller/auth"
	"github.com/rumbel/belajar/internal/app/service"
)

var (
    authService    service.AuthService    = service.NewAuthService()
)

func AuthRoutes(api *gin.RouterGroup, db *gorm.DB) {
	authController := controller.NewAuthController(authService, db)
	
	auth := api.Group("/auth")
	{
		// test endpoint
		auth.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "this is auth endpoint",
			})
		})
		// register endpoint
		auth.POST("/register", func(ctx *gin.Context) {
			err := authController.Register(ctx)
			if err != nil {
				ctx.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(200, gin.H{
				"message": "registration success",
			})
		})
		// login endpoint
		auth.POST("/login", func(ctx *gin.Context) {
			token, err := authController.Login(ctx)
			if err != nil {
				ctx.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(200, gin.H{
				"message": "login success",
				"token": token,
			})
		})
	}
}
