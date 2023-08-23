package controller

import (
	// "log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rumbel/belajar/internal/app/entity"
	"github.com/rumbel/belajar/internal/app/service"
	"github.com/rumbel/belajar/pkg/validators"
)

type AuthController interface {
	Register(ctx *gin.Context) error
}

type authController struct {
	service service.AuthService
}

var authValidate *validator.Validate

func NewAuthController(service service.AuthService) AuthController {
	authValidate = validator.New()
	authValidate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &authController{
		service: service,
	}
}

func (c *authController) Register(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	err = authValidate.Struct(user)
	if err != nil {
		return err
	}
	// log.Println("user in controller:", user)
	c.service.Register(user)
	return nil
}