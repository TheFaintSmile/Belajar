package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rumbel/belajar/internal/app/controller/test"
)

func TestRoutes(api *gin.RouterGroup) {
	test := api.Group("/test")
	{
		test.GET("/", controller.NewTestControllerImpl().GetTests)
	}
}
