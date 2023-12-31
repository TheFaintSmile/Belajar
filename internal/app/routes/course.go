package routes

import (
	"fmt"

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

func CourseAndModuleRoutes(api *gin.RouterGroup, db *gorm.DB) {
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

		courseList.Use(middlewares.PendidikAdminAuth())
		courseList.PATCH("/:id/week/:weekID/", UpdateWeekInCourse(courseController))
		courseList.DELETE("/:id/week/:weekID/", DeleteWeekInCourse(courseController))
		courseList.PATCH("/:id/", UpdateCourseInformation(courseController))
		courseList.DELETE("/:id/", DeleteCourse(courseController))
		courseList.POST("/", AddCourse(courseController))
		courseList.POST("/week/", AddWeekToCourse(courseController))
		courseList.POST("/:id/week/:weekID/", AddModuleToCourse(courseController))
	}

	moduleList := api.Group("/module")
	{
		moduleList.Use(middlewares.PendidikAdminAuth())
		moduleList.DELETE("/material/:id/", DeleteMaterialFromCourse(courseController))
		moduleList.DELETE("/task/:id/", DeleteTaskFromCourse(courseController))
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

		result, err := courseController.GetCourseDetail(ctx)

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
		result, err := courseController.AddWeekToCourse(ctx)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}

		utils.SuccessResponse(ctx, "Successfully Added Week", result)
	}
}

func UpdateCourseInformation(courseController *controller.CourseController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := courseController.UpdateCourseInformation(ctx)
		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}
		utils.SuccessResponse(ctx, "Successfully Updated Course Information", result)
	}
}
func DeleteCourse(courseController *controller.CourseController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := courseController.DeleteCourse(ctx)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}
		utils.SuccessResponse(ctx, "Successfully Deleted Course", nil)
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
		err := courseController.DeleteWeekInCourse(ctx)

		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}
		utils.SuccessResponse(ctx, "Successfully Deleted Week", nil)
	}
}

func AddModuleToCourse(courseController *controller.CourseController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("masuk")
		result, err := courseController.AddModuleToCourse(ctx)
		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}
		utils.SuccessResponse(ctx, "Sucessfully Added Module", result)
	}
}

func DeleteMaterialFromCourse(courseController *controller.CourseController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := courseController.DeleteMaterialFromCourse(ctx)
		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}
		utils.SuccessResponse(ctx, "Sucessfully Deleted Material", nil)
	}
}

func DeleteTaskFromCourse(courseController *controller.CourseController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := courseController.DeleteTaskFromCourse(ctx)
		if err != nil {
			utils.ErrorResponse(ctx, err.Error(), nil)
			return
		}
		utils.SuccessResponse(ctx, "Sucessfully Deleted Material", nil)
	}
}
