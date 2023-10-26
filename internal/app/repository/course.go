package repository

import (
	"github.com/rumbel/belajar/internal/app/dto"
	"github.com/rumbel/belajar/internal/app/models"
	"github.com/rumbel/belajar/internal/app/utils"
)

type CourseRepository struct{}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{}
}

func (repository *CourseRepository) GetCourseList(userLevel int) ([]dto.CourseListResponse, error) {
	var courses []dto.CourseListResponse

	if err := utils.DB.Model(&models.Course{}).Select("id, name, lecturer").Where("level_id = ?", uint(userLevel)).Scan(&courses).Error; err != nil {
		return nil, err
	}

	return courses, nil
}

func (repository *CourseRepository) GetCourseDetail(id uint) (models.Course, error) {
	var course models.Course

	if err := utils.DB.Preload("Weeks").First(&course, id).Preload("Weeks.Tasks").Preload("Weeks.Materials").First(&course, id).Error; err != nil {
		return models.Course{}, err
	}

	return course, nil
}

func (repository *CourseRepository) AddCourse(course dto.AddCourseInput) (dto.AddCourseInput, error) {
	newCourse := models.Course{
		Name:     course.Name,
		Lecturer: course.Lecturer,
		LevelID:  course.LevelID,
	}
	if err := utils.DB.Create(&newCourse).Error; err != nil {
		return dto.AddCourseInput{}, err
	}

	return course, nil
}

func (repository *CourseRepository) AddWeekToCourse(week models.Week) (models.Week, error) {

	if err := utils.DB.Create(&week).Error; err != nil {
		return models.Week{}, err
	}

	return week, nil
}

func (repository *CourseRepository) GetWeekOccurrence(courseID uint) (int, error) {
	var occurrence int64

	if err := utils.DB.Model(&models.Week{}).Where("course_id = ?", courseID).Count(&occurrence).Error; err != nil {
		return 0, err
	}

	return int(occurrence), nil
}

func (repository *CourseRepository) UpdateCourseInformation(courseID uint, course dto.UpdateCourseInformationInput) (dto.UpdateCourseInformationInput, error) {
	var courseInstance models.Course

	if err := utils.DB.First(&courseInstance, courseID).Error; err != nil {
		return dto.UpdateCourseInformationInput{}, err
	}

	updatedCourse := models.Course{
		Name:     course.Name,
		Lecturer: course.Lecturer,
		LevelID:  course.LevelID,
	}

	if err := utils.DB.Model(&courseInstance).Updates(&updatedCourse).Error; err != nil {
		return dto.UpdateCourseInformationInput{}, err
	}

	return course, nil
}

func (repository *CourseRepository) DeleteCourse(id uint) error {
	var course models.Course

	if err := utils.DB.First(&course, id).Error; err != nil {
		return err
	}

	if err := utils.DB.Delete(&course).Error; err != nil {
		return err
	}

	return nil
}

func (repository *CourseRepository) UpdateWeekInCourse(courseID uint, weekID uint, week dto.UpdateWeekInCourseInput) (dto.UpdateWeekInCourseInput, error) {
	var weekInstance models.Week

	if err := utils.DB.Where("course_id = ? AND week_number = ?", courseID, weekID).First(&weekInstance).Error; err != nil {
		return dto.UpdateWeekInCourseInput{}, err
	}

	updatedWeek := models.Week{
		Name: week.Name,
	}

	if err := utils.DB.Model(&weekInstance).Updates(&updatedWeek).Error; err != nil {
		return dto.UpdateWeekInCourseInput{}, err
	}

	return week, nil
}

func (repository *CourseRepository) DeleteWeekInCourse(courseID uint, weekID uint) error {
	var week models.Week

	if err := utils.DB.Where("course_id = ? AND week_number = ?", courseID, weekID).First(&week).Error; err != nil {
		return err
	}

	if err := utils.DB.Delete(&week).Error; err != nil {
		return err
	}

	return nil
}

func (repository *CourseRepository) AddModuleToCourse(courseID uint, weekID uint, module dto.AddModuleToCourse) (dto.AddModuleToCourse, error) {
	var weekInstance models.Week

	if err := utils.DB.Where("course_id = ? AND week_number = ?", courseID, weekID).First(&weekInstance).Error; err != nil {
		return dto.AddModuleToCourse{}, err
	}

	if module.Category == dto.CategoryMaterial {
		newMaterial := models.Material{
			Name:        module.Name,
			Description: module.Description,
			Type:        module.Type,
			Content:     module.Content,
			WeekID:      weekInstance.ID,
		}
		if err := utils.DB.Create(&newMaterial).Error; err != nil {
			return dto.AddModuleToCourse{}, err
		}
	} else {
		newTask := models.Task{
			Name:        module.Name,
			Description: module.Description,
			Type:        module.Type,
			Content:     module.Content,
			WeekID:      weekInstance.ID,
		}
		if err := utils.DB.Create(&newTask).Error; err != nil {
			return dto.AddModuleToCourse{}, err
		}
	}

	return module, nil
}

func (repository *CourseRepository) DeleteMaterialFromCourse(materialID uint) error {
	var material models.Material

	if err := utils.DB.First(&material, materialID).Error; err != nil {
		return err
	}

	if err := utils.DB.Delete(&material).Error; err != nil {
		return err
	}

	return nil
}

func (repository *CourseRepository) DeleteTaskFromCourse(taskID uint) error {
	var task models.Task

	if err := utils.DB.First(&task, taskID).Error; err != nil {
		return err
	}

	if err := utils.DB.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}
