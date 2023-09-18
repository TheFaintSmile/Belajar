package service

import (
	"github.com/rumbel/belajar/internal/app/models"
	"github.com/rumbel/belajar/internal/app/repository"
)

type CourseService struct {
	repository *repository.CourseRepository
}

func NewCourseService(repository *repository.CourseRepository) *CourseService {
	return &CourseService{
		repository: repository,
	}
}

func (s *CourseService) GetCourseList() ([]models.Course, error) {
	courses, err := s.repository.GetCourseList()

	if err != nil {
		return nil, err
	}

	return courses, nil
}
