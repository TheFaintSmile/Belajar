package routes

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	controller "github.com/rumbel/belajar/internal/app/controller/auth"
	"github.com/rumbel/belajar/internal/app/service"
)

var (
	authService service.AuthService = service.NewAuthService()
	authController controller.AuthController = controller.NewAuthController(authService)
)

func AuthRoutes(api *gin.RouterGroup) {
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
			// Log the entire request details
			data, _ := ioutil.ReadAll(ctx.Request.Body)
			ctx.JSON(200, gin.H{
				"message": "success",
				"data": data,
			})
		})
	}
}
