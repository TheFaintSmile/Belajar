package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rumbel/belajar/internal/app/controller"
	"github.com/rumbel/belajar/internal/app/middlewares"
	"github.com/rumbel/belajar/internal/app/repository"
	"github.com/rumbel/belajar/internal/app/service"
	"github.com/rumbel/belajar/internal/app/utils"
)

var (
	courseService *service.CourseService = service.NewCourseService(repository.NewCourseRepository())
)

func CourseRoutes(api *gin.RouterGroup, db *gorm.DB) {
	courseController := controller.NewCourseController(*courseService)

	courseList := api.Group("/course")
	{
		// Will be removed in production later.
		courseList.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Testing routes",
			})
		})
		courseList.GET("/", GetCourseList(courseController))
	}
}

func GetCourseList(courseController *controller.CourseController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := middlewares.TokenValid(ctx)

		if err != nil {
			utils.ErrorResponse(ctx, "unauthorized", nil)
			return
		}

		result, err := courseController.GetCourseList(ctx)
		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}

		utils.SuccessResponse(ctx, "Successfully GET Data.", result)
	}
}
