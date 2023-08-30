package service

import (
	// "log"
	"fmt"

	"github.com/rumbel/belajar/internal/app/entity"
	"github.com/rumbel/belajar/internal/app/utils"
)

type AuthService interface {
	Register(entity.User) (string, error)
	Login(entity.User) (string, error)
}

type authService struct {
	users []entity.User
}

func NewAuthService() AuthService {
	return &authService{
		users: []entity.User{},
	}
}

func (service *authService) Register(user entity.User) (string, error) {
	if service.checkUserExists(user.Email) {
		return "", fmt.Errorf("email already exists")
	}

	u := entity.User{}
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Age = user.Age
	u.Email = user.Email
	u.Password = user.Password

	_, err := u.SaveUser()
	if err != nil {
		return "", err
	}

	service.users = append(service.users, user)
	return "success", nil
}

func (service *authService) Login(user entity.User) (string, error) {
	token, err := entity.LoginCheck(user.Email, user.Password)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (service *authService) checkUserExists(email string) bool {
	var u entity.User
	err := utils.DB.Model(&entity.User{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return false
	}
	return true
}