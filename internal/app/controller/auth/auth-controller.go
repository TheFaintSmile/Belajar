package controller

import (
	// "log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/rumbel/belajar/internal/app/entity"
	"github.com/rumbel/belajar/internal/app/service"
	"github.com/rumbel/belajar/pkg/validators"
)

type AuthController interface {
	Register(ctx *gin.Context) error
	Login(ctx *gin.Context) (string, error)
}

type authController struct {
	service service.AuthService
	db 	*gorm.DB
}

var authValidate *validator.Validate

func NewAuthController(service service.AuthService, db *gorm.DB) AuthController {
	authValidate = validator.New()
	authValidate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &authController{
		service: service,
		db: db,
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
	_, err = c.service.Register(user)
	if err != nil {
		return err
	}
	return nil
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func (c *authController) Login(ctx *gin.Context) (string, error) {
	var input LoginInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		return "",err
	}
	u := entity.User{}
	u.Email = input.Email
	u.Password = input.Password

	token, err := c.service.Login(u)
	if err != nil {
		return "", err
	}
	return token, nil
}
