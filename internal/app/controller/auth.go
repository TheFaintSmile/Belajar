package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/rumbel/belajar/internal/app/dto"
	"github.com/rumbel/belajar/internal/app/models"
	"github.com/rumbel/belajar/internal/app/service"
	"github.com/rumbel/belajar/pkg/validators"
)

var authValidate *validator.Validate

type AuthController interface {
	Register(ctx *gin.Context) error
	Login(ctx *gin.Context) (string, error)
	GetUserInfo(userID uint) (*models.User, error)
}

func (c *authController) Register(ctx *gin.Context) error {
	var user models.User
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

func (c *authController) Login(ctx *gin.Context) (string, error) {
	var input dto.LoginInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		return "",err
	}
	u := models.User{}
	u.Email = input.Email
	u.Password = input.Password

	token, err := c.service.Login(u)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (c *authController) GetUserInfo(userID uint) (*models.User, error) {
	var user models.User
	if err := c.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

type authController struct {
	service service.AuthService
	db 	*gorm.DB
}

func NewAuthController(service service.AuthService, db *gorm.DB) AuthController {
	authValidate = validator.New()
	authValidate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &authController{
		service: service,
		db: db,
	}
}