package dto

import (
	models "github.com/rumbel/belajar/internal/app/models"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type (
	// CreateUserRequest CreateUserRequest
	CreateUserRequest struct {
		FirstName 	string 				`json:"firstname" validate:"required"`
		LastName	string 				`json:"lastname" validate:"required"`
		Age			int8 				`json:"age" validate:"required"`
		Email    	string 				`json:"email" validate:"required,email"`
		Password	string 				`json:"password" validate:"required"`
		Level     	models.UserLevel 	`json:"level" validate:"required"`
	}

	// CreateUserRequest CreateUserRequest
	CreateUserResponse struct {
		FirstName 	string 				`json:"firstname" validate:"required"`
		LastName	string 				`json:"lastname" validate:"required"`
		Age			int8 				`json:"age" validate:"required"`
		Email    	string 				`json:"email" validate:"required,email"`
		Password	string 				`json:"password" validate:"required"`
		Level     	models.UserLevel 	`json:"level" validate:"required"`
	}
)