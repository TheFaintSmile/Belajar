package repository

import (
	"github.com/rumbel/belajar/internal/app/models"
	"github.com/rumbel/belajar/internal/app/utils"
)

type CourseRepository struct{}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{}
}

func (repository *CourseRepository) GetCourseList(userLevel int) ([]models.Course, error) {
	var level models.Level

	if err := utils.DB.Preload("Courses").First(&level, userLevel).Error; err != nil {
		return nil, err
	}

	return level.Courses, nil
}

func (repository *CourseRepository) AddCourse(course models.Course) (models.Course, error) {

	if err := utils.DB.Create(&course).Error; err != nil {
		return models.Course{}, err
	}

	return course, nil
}
