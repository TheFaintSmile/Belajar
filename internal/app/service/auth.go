package service

import (
	// "log"
	"fmt"
	"net/mail"
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
	if !service.isValidEmail(user.Email) {
		return "", fmt.Errorf("invalid email")
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
	u.Level = user.Level

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

//* CONSIDERING: move these helper functions to utils package
func (service *authService) checkEmailExists(email string) bool {
	var u entity.User
	err := utils.DB.Model(&entity.User{}).Where("email = ?", email).Take(&u).Error
	return err == nil;
}

func (service *authService) containsWhiteSpace(password string) bool {
	return strings.Contains(password, " ")
}

func (service *authService) isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}