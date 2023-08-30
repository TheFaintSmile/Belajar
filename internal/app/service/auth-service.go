package service

import (
	// "log"
	"fmt"
	"strings"

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
	if service.checkEmailExists(user.Email) {
		return "", fmt.Errorf("email already exists")
	}
	if service.checkNameExists(user.FirstName, user.LastName) {
		return "", fmt.Errorf("name already exists")
	}
	if service.containsWhiteSpace(user.Password) {
		return "", fmt.Errorf("password cannot contain whitespace")
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

func (service *authService) checkEmailExists(email string) bool {
	var u entity.User
	err := utils.DB.Model(&entity.User{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return false
	}
	return true
}

func (service *authService) checkNameExists(fname string, lname string) bool {
	var u entity.User
	err := utils.DB.Model(&entity.User{}).Where("first_name = ? AND last_name = ?", fname, lname).Take(&u).Error
	if err != nil {
		return false
	}
	return true
}

func (service *authService) containsWhiteSpace(password string) bool {
	return strings.Contains(password, " ")
}