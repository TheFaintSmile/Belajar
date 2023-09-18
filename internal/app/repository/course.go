package repository

import (
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
	var courses []models.Course

	if err := repository.DB.Find(&courses).Error; err != nil {
		return nil, err
	}

	return courses, nil
}
