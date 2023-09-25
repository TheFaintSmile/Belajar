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

// All godoc
// @Tags Auth
// @Summary Register New User
// @Description Put all mandatory parameter
// @Param CreateUserRequest body dto.CreateUserRequest true "CreateUserRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.CreateUserResponse
// @Failure 200 {object} dto.CreateUserResponse
// @Router /auth/register [post]
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
	token, err := c.service.Login(input)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (c *authController) GetUserInfo(userID uint) (*models.User, error) {
	user, err := c.service.GetUserInfo(userID); 
	
	if err != nil {
		return &models.User{}, err
	}

	return user, nil
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