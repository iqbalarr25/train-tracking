package model

import "github.com/google/uuid"

type (
	LoginAdminRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		Id    uuid.UUID `json:"id"`
		Email string    `json:"email"`
		Token string    `json:"token"`
	}
)
