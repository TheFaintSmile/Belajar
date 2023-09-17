package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rumbel/belajar/internal/app/utils"
)

func CourseListRoutes(api *gin.RouterGroup, db *gorm.DB) {
	courseList := api.Group("/course")
	{
		// Will be removed in production later.
		courseList.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Testing routes",
			})
		})
		courseList.GET("/", GetCourseList())
	}
}

func GetCourseList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		utils.SuccessResponse(ctx, "Berhasil", "")
	}
}
