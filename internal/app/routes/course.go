package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rumbel/belajar/internal/app/controller"
	"github.com/rumbel/belajar/internal/app/middlewares"
	"github.com/rumbel/belajar/internal/app/models"
	"github.com/rumbel/belajar/internal/app/repository"
	"github.com/rumbel/belajar/internal/app/service"
	utils "github.com/rumbel/belajar/internal/app/utils"
)

var (
	courseService = service.NewCourseService(repository.NewCourseRepository())
	levelMap      = map[string]int{
		"SD-1": 1,
		"SD-2": 2,
		"SD-3": 3,
		"SD-4": 4,
		"SD-5": 5,
		"SD-6": 6,
		"SMP":  7,
		"SMA":  8,
	}
)

func CourseRoutes(api *gin.RouterGroup, db *gorm.DB) {
	courseController := controller.NewCourseController(*courseService)

	courseList := api.Group("/course")
	{
		courseList.Use(middlewares.Auth())
		// Will be removed in production later.
		courseList.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Testing routes",
			})
		})
		courseList.GET("/", GetCourseList(courseController))
		courseList.GET("/:id/", GetCourseDetail(courseController))
		courseList.PATCH("/:id/", UpdateWeekInCourse(courseController))
		courseList.DELETE("/:id/", DeleteWeekInCourse(courseController))
		courseList.POST("/", AddCourse(courseController))
		courseList.POST("/week/", AddWeekToCourse(courseController))
	}
}

func GetCourseList(courseController *controller.CourseController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, err := middlewares.ExtractTokenID(ctx)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}

		userInfo, err := authService.GetUserInfo(userID)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}

		userLevel := levelMap[string(userInfo.LevelID)]

		result, err := courseController.GetCourseList(ctx, userLevel)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}

		utils.SuccessResponse(ctx, "Successfully GET Data.", result)
	}
}

func GetCourseDetail(courseController *controller.CourseController) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userID, err := middlewares.ExtractTokenID(ctx)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}

		userInfo, err := authService.GetUserInfo(userID)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}

		result, err := courseController.GetCourseDetail(ctx, userInfo)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}
		if (userInfo.Role != models.RoleSiswa) || (levelMap[string(userInfo.LevelID)] == int(result.LevelID)) {
			utils.SuccessResponse(ctx, "Successfully GET Data.", result)
		} else {
			utils.ErrorResponse(ctx, "Forbidden Endpoint", nil)
		}
	}
}

func AddCourse(courseController *controller.CourseController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, err := middlewares.ExtractTokenID(ctx)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}

		userInfo, err := authService.GetUserInfo(userID)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}

		if userInfo.Role != models.RoleAdmin {
			utils.ErrorResponse(ctx, "Forbidden Endpoint", nil)
			return
		}
		result, err := courseController.AddCourse(ctx)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}

		utils.SuccessResponse(ctx, "Successfully POST Data.", result)
	}
}

func AddWeekToCourse(courseController *controller.CourseController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, err := middlewares.ExtractTokenID(ctx)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}

		userInfo, err := authService.GetUserInfo(userID)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}

		if userInfo.Role == models.RoleSiswa {
			utils.ErrorResponse(ctx, "Forbidden Endpoint", nil)
			return
		}

		result, err := courseController.AddWeekToCourse(ctx)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}

		utils.SuccessResponse(ctx, "Successfully Added Week", result)
	}
}

func UpdateWeekInCourse(courseController *controller.CourseController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		result, err := courseController.UpdateWeekInCourse(ctx)
		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}
		utils.SuccessResponse(ctx, "Successfully Updated Week Information", result)
	}
}
func DeleteWeekInCourse(courseController *controller.CourseController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := courseController.DeleteWeekInCourse(ctx)
		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}
		utils.SuccessResponse(ctx, "Successfully Deleted Week", nil)
	}
}
