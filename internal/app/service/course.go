package service

import (
	"errors"

	"github.com/rumbel/belajar/internal/app/dto"
	"github.com/rumbel/belajar/internal/app/models"
	"github.com/rumbel/belajar/internal/app/repository"
	"github.com/rumbel/belajar/internal/app/utils"
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

func (s *CourseService) GetWeekOccurrence(courseID uint) (int, error) {
	occurrence, err := s.repository.GetWeekOccurrence(courseID)

	if err != nil {
		return 0, err
	}

	return occurrence, nil
}

func (s *CourseService) UpdateCourseInformation(courseID uint, course dto.UpdateCourseInformationInput) (dto.UpdateCourseInformationInput, error) {
	course, err := s.repository.UpdateCourseInformation(courseID, course)

	if err != nil {
		return dto.UpdateCourseInformationInput{}, err
	}

	return course, nil
}

func (s *CourseService) DeleteCourse(courseID uint) error {
	err := s.repository.DeleteCourse(courseID)

	if err != nil {
		return err
	}

	return nil
}

func (s *CourseService) UpdateWeekInCourse(courseID uint, weekID uint, week dto.UpdateWeekInCourseInput) (dto.UpdateWeekInCourseInput, error) {
	res, err := s.repository.UpdateWeekInCourse(courseID, weekID, week)

	if err != nil {
		return dto.UpdateWeekInCourseInput{}, err
	}

	return res, nil
}

func (s *CourseService) DeleteWeekInCourse(courseID uint, weekID uint) error {
	err := s.repository.DeleteWeekInCourse(courseID, weekID)

	if err != nil {
		return err
	}

	return nil
}

func (s *CourseService) AddModuleToCourse(courseID uint, weekID uint, module dto.AddModuleToCourse) (dto.AddModuleToCourse, error) {
	if err := utils.IsValidCategory(&module); err == nil {
		if err := utils.IsValidModuleType(&module); err == nil {
			if module.Type == models.ModuleTypeFile {
				if !(module.File != nil) {
					return dto.AddModuleToCourse{}, errors.New("file is required")
				}
				
				content, err := utils.UploadFile(module.File)
				if err != nil {
					return dto.AddModuleToCourse{}, err
				}
				module.Content = content
			}
		} else {
			return dto.AddModuleToCourse{}, err
		}

		res, err := s.repository.AddModuleToCourse(courseID, weekID, module)

		if err != nil {
			return dto.AddModuleToCourse{}, err
		}

		return res, nil
	} else {
		return dto.AddModuleToCourse{}, err
	}
}

func (s *CourseService) DeleteMaterialFromCourse(materialD uint) error {
	err := s.repository.DeleteCourse(materialD)

	if err != nil {
		return err
	}

	return nil
}