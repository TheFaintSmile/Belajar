package service

import (
	// "log"
	"fmt"

	"github.com/rumbel/belajar/internal/app/models"
	"github.com/rumbel/belajar/internal/app/repository"
	"github.com/rumbel/belajar/internal/app/utils"
)

type AuthService interface {
	Register(models.User) (string, error)
	Login(models.User) (string, error)
}

type authService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (s *authService) Register(user models.User) (string, error) {
	if s.userRepository.CheckEmailExists(user.Email) {
		return "", fmt.Errorf("email already exists")
	}
	if !utils.IsValidEmail(user.Email) {
		return "", fmt.Errorf("invalid email")
	}
	if utils.ContainsWhiteSpace(user.Password) {
		return "", fmt.Errorf("password cannot contain whitespace")
	}

	u := models.User{}
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Age = user.Age
	u.Email = user.Email
	u.Password = user.Password
	u.Level = user.Level

	_, err := s.userRepository.SaveUser(&u)
	if err != nil {
		return "", err
	}

	return "success", nil
}

func (s *authService) Login(user models.User) (string, error) {
	token, err := s.userRepository.LoginCheck(user.Email, user.Password)
	if err != nil {
		return "", err
	}
	return token, nil
}

