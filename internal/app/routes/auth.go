package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	controller "github.com/rumbel/belajar/internal/app/controller/auth"
	"github.com/rumbel/belajar/internal/app/service"
	utils "github.com/rumbel/belajar/internal/app/utils"
)

var (
	authService service.AuthService = service.NewAuthService()
)

func AuthRoutes(api *gin.RouterGroup, db *gorm.DB) {
	authController := controller.NewAuthController(authService, db)

	auth := api.Group("/auth")
	{
		// test endpoint
		auth.GET("/", func(ctx *gin.Context) {
			utils.SendResponse(ctx, 200, "This is the auth endpoint", nil)
		})
		// register endpoint
		auth.POST("/register", func(ctx *gin.Context) {
			err := authController.Register(ctx)
			if err != nil {
				utils.SendResponse(ctx, 400, err.Error(), nil)
				return
			}
			utils.SendResponse(ctx, 200, "Registration success", nil)
		})
		// login endpoint
		auth.POST("/login", func(ctx *gin.Context) {
			token, err := authController.Login(ctx)
			if err != nil {
				utils.SendResponse(ctx, 400, err.Error(), nil)
				return
			}
			utils.SendResponse(ctx, 200, "Login success", gin.H{"token": token})
		})
	}
}