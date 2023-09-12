package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	controller "github.com/rumbel/belajar/internal/app/controller/auth"
	"github.com/rumbel/belajar/internal/app/middlewares"
	"github.com/rumbel/belajar/internal/app/repository"
	"github.com/rumbel/belajar/internal/app/service"
	utils "github.com/rumbel/belajar/internal/app/utils"
)

var (
	authService service.AuthService = service.NewAuthService(repository.NewUserRepository())
)

func AuthRoutes(api *gin.RouterGroup, db *gorm.DB) {
	authController := controller.NewAuthController(authService, db)

	auth := api.Group("/auth")
	{
		auth.POST("/register", Register(authController))
		auth.POST("/login", Login(authController))
		auth.GET("/credential", Credential(authController))
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

func Credential(authController controller.AuthController) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        userID, err := middlewares.ExtractTokenID(ctx)
        if err != nil {
            utils.ErrorResponse(ctx, err.Error(), nil)
            return
        }
        userInfo, err := authController.GetUserInfo(userID)
        if err != nil {
            utils.ErrorResponse(ctx, err.Error(), nil)
            return
        }
        utils.SuccessResponse(ctx, "User Info Retrieved", userInfo)
    }
}
