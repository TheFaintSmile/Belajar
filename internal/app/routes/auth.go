package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	controller "github.com/rumbel/belajar/internal/app/controller/auth"
	"github.com/rumbel/belajar/internal/app/service"
)

var (
	authService service.AuthService = service.NewAuthService()
)

// Response is a custom struct to format responses consistently
type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

// sendResponse sends a response with a standardized format
func sendResponse(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(status, Response{
		Message: message,
		Status:  status,
		Data:    data,
	})
}

func AuthRoutes(api *gin.RouterGroup, db *gorm.DB) {
	authController := controller.NewAuthController(authService, db)

	auth := api.Group("/auth")
	{
		// test endpoint
		auth.GET("/", func(ctx *gin.Context) {
			sendResponse(ctx, 200, "This is the auth endpoint", nil)
		})
		// register endpoint
		auth.POST("/register", func(ctx *gin.Context) {
			err := authController.Register(ctx)
			if err != nil {
				sendResponse(ctx, 400, err.Error(), nil)
				return
			}
			sendResponse(ctx, 200, "Registration success", nil)
		})
		// login endpoint
		auth.POST("/login", func(ctx *gin.Context) {
			token, err := authController.Login(ctx)
			if err != nil {
				sendResponse(ctx, 400, err.Error(), nil)
				return
			}
			sendResponse(ctx, 200, "Login success", gin.H{"token": token})
		})
	}
}
