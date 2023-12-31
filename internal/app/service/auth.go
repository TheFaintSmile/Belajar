package service

import (
	"fmt"

	"github.com/rumbel/belajar/internal/app/dto"
	"github.com/rumbel/belajar/internal/app/models"
	"github.com/rumbel/belajar/internal/app/repository"
	"github.com/rumbel/belajar/internal/app/utils"
)

type AuthService interface {
	Register(models.User) (string, error)
	Login(dto.LoginInput) (string, error)
	GetUserInfo(userID uint) (*models.User, error)
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
	u.LevelID = user.LevelID

	u.Role = user.Role

	if u.Role == models.RoleSiswa {
		err := utils.IsValidLevel(&u)
		if err != nil {
			return "", err
		}
	}

	err := utils.HashingPassword(&u)
	if err != nil {
		return "", err
	}

	_, err = s.userRepository.SaveUser(&u)
	if err != nil {
		return "", err
	}

	return "success", nil
}

func (s *authService) Login(userInput dto.LoginInput) (string, error) {
	token, err := s.userRepository.LoginCheck(userInput.Email, userInput.Password)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *authService) GetUserInfo(userID uint) (*models.User, error) {
	user, err := s.userRepository.GetUserInfo(userID)

	if err != nil {
		return &models.User{}, err
	}

	return user, nil
}
type authService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}
