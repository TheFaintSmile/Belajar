package repository

import (
	"fmt"

	middlewares "github.com/rumbel/belajar/internal/app/middlewares"
	"github.com/rumbel/belajar/internal/app/models"
	"github.com/rumbel/belajar/internal/app/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) SaveUser(user *models.User) (*models.User, error) {
	err := utils.DB.Create(user).Error
	if err != nil {
		return &models.User{}, err
	}
	if user.Role == models.RoleSiswa {
		siswa := models.Siswa{}
		siswa.Email = user.Email
		err = utils.DB.Create(&siswa).Error
		if err != nil {
			return &models.User{}, err
		}
	} else if user.Role == models.RolePendidik {
		pendidik := models.Pendidik{}
		pendidik.Email = user.Email
		err = utils.DB.Create(&pendidik).Error
		if err != nil {
			return &models.User{}, err
		}
	} else if user.Role == models.RoleAdmin {
		admin := models.Admin{}
		admin.Email = user.Email
		err = utils.DB.Create(&admin).Error
		if err != nil {
			return &models.User{}, err
		}
	}
	return user, nil
}

func (ur *UserRepository) LoginCheck(email, password string) (string, error) {
	var err error
	user := models.User{}
	err = utils.DB.Model(&models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", fmt.Errorf("user not found")
	}
	err = utils.VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", fmt.Errorf("invalid login credentials")
	}
	token, err := middlewares.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (ur *UserRepository) CheckEmailExists(email string) bool {
	var u models.User
	err := utils.DB.Model(&models.User{}).Where("email = ?", email).Take(&u).Error
	return err == nil;
}