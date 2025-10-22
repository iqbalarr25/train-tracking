package model

import "github.com/google/uuid"

type (
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		Id    uuid.UUID `json:"id"`
		Email string    `json:"email"`
		Name  string    `json:"name"`
		Token string    `json:"token"`
	}

	RegisterRequest struct {
		Email                string `json:"email" validate:"required,email"`
		Name                 string `json:"name" validate:"required"`
		Password             string `json:"password" validate:"required"`
		ConfirmationPassword string `json:"confirmation_password" validate:"required"`
	}

	RegisterResponse struct {
		Id    uuid.UUID `json:"id"`
		Email string    `json:"email"`
		Name  string    `json:"name"`
		Token string    `json:"token"`
	}
)
