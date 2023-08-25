package service

import (
	// "log"

	"github.com/rumbel/belajar/internal/app/entity"
)

type AuthService interface {
	Register(entity.User) entity.User
}

type authService struct {
	users []entity.User
}

func NewAuthService() AuthService {
	return &authService{
		users: []entity.User{},
	}
}

func (service *authService) Register(user entity.User) entity.User {
	u := entity.User{}
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Age = user.Age
	u.Email = user.Email
	u.Password = user.Password

	_, err := u.SaveUser()
	if err != nil {
		return entity.User{}
	}

	service.users = append(service.users, user)
	return user
}