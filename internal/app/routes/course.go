package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rumbel/belajar/internal/app/controller"
	"github.com/rumbel/belajar/internal/app/utils"
)

func CourseRoutes(api *gin.RouterGroup, db *gorm.DB) {
	courseController := controller.NewCourseController()

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
		result, err := courseController.CourseListController(ctx)
		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}

		utils.SuccessResponse(ctx, "Berhasil", result)
	}
}
