package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rumbel/belajar/internal/app/models"
)

type CourseRepository struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{DB: db}
}

func (repository *CourseRepository) GetCourseList() ([]models.Course, error) {
	var level models.Level

	if err := repository.DB.Preload("Courses").First(&level, 6).Error; err != nil {
		return nil, err
	}

	fmt.Println(level.Courses)

	return level.Courses, nil
}

func (repository *CourseRepository) AddCourse(course models.Course) (models.Course, error) {

	if err := repository.DB.Create(&course).Error; err != nil {
		return models.Course{}, err
	}

	return course, nil
}
