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
	var level models.Level

	if err := repository.DB.Model(&models.Level{}).Preload("Courses").Find(&level).Error; err != nil {
		return nil, err
	}

	return level.Courses, nil
}
