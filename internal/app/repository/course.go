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

func (repository *CourseRepository) GetWeekOccurrence(course_id uint) (int, error) {
	var occurrence int64

	if err := utils.DB.Model(&models.Week{}).Where("course_id = ?", course_id).Count(&occurrence).Error; err != nil {
		return 0, err
	}

	return int(occurrence), nil
}

func (repository *CourseRepository) UpdateCourseInformation(course_id uint, course dto.UpdateCourseInformationInput) (dto.UpdateCourseInformationInput, error) {
	var course_instance models.Course

	if err := utils.DB.First(&course_instance, course_id).Error; err != nil {
		return dto.UpdateCourseInformationInput{}, err
	}

	updatedCourse := models.Course{
		Name:     course.Name,
		Lecturer: course.Lecturer,
		LevelID:  course.LevelID,
	}

	if err := utils.DB.Model(&course_instance).Updates(&updatedCourse).Error; err != nil {
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

func (repository *CourseRepository) UpdateWeekInCourse(week uint) (int, error) {

	return 0, nil
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
