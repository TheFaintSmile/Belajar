package repository

import "github.com/jinzhu/gorm"

type CourseRepository struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{DB: db}
}

func (repository *CourseRepository) GetCourseList() {

}
