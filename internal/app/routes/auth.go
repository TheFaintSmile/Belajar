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
		auth.POST("/register", Register(authController))
		auth.POST("/login", Login(authController))
	}
}

func Register(authController controller.AuthController) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        err := authController.Register(ctx)
        if err != nil {
            utils.ErrorResponse(ctx, err.Error(), nil)
            return
        }
        utils.SuccessResponse(ctx, "Registration success", nil)
    }
}

func Login(authController controller.AuthController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := authController.Login(ctx)
		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}
		utils.SuccessResponse(ctx, "Login success", gin.H{"token": token})
	}
}
