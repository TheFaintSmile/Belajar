package service

import (
	"github.com/rumbel/belajar/internal/app/dto"
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

func (s *CourseService) GetCourseList(userLevel int) ([]dto.CourseListResponse, error) {
	courses, err := s.repository.GetCourseList(userLevel)

	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (s *CourseService) GetCourseDetail(id uint) (models.Course, error) {
	course, err := s.repository.GetCourseDetail(id)

	if err != nil {
		return models.Course{}, err
	}

	return course, nil
}

func (s *CourseService) AddCourse(course dto.AddCourseInput) (dto.AddCourseInput, error) {
	result, err := s.repository.AddCourse(course)

	if err != nil {
		return dto.AddCourseInput{}, err
	}

	return result, nil
}

func (s *CourseService) AddWeekToCourse(week models.Week) (models.Week, error) {
	result, err := s.repository.AddWeekToCourse(week)

	if err != nil {
		return models.Week{}, err
	}

	return result, nil
}

func (s *CourseService) GetWeekOccurrence(course_id uint) (int, error) {
	occurrence, err := s.repository.GetWeekOccurrence(course_id)

	if err != nil {
		return 0, err
	}

	return occurrence, nil
}

func (s *CourseService) UpdateCourseInformation(course_id uint, course dto.UpdateCourseInformationInput) (dto.UpdateCourseInformationInput, error) {
	course, err := s.repository.UpdateCourseInformation(course_id, course)

	if err != nil {
		return dto.UpdateCourseInformationInput{}, err
	}

	return course, nil
}

func (s *CourseService) DeleteCourse(course_id uint) error {
	err := s.repository.DeleteCourse(course_id)

	if err != nil {
		return err
	}

	return nil
}

func (s *CourseService) UpdateWeekInCourse(week uint) (int, error) {
	occurrence, err := s.repository.UpdateWeekInCourse(week)

	if err != nil {
		return 0, err
	}

	return occurrence, nil
}

func (s *CourseService) DeleteWeekInCourse(courseID uint, week uint) error {
	err := s.repository.DeleteWeekInCourse(courseID, week)

	if err != nil {
		return err
	}

	return nil
}
