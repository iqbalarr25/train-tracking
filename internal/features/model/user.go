package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type (
	User struct {
		ID        uuid.UUID `gorm:"primary_key" json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Role      string    `json:"role"`
		FcmToken  string    `json:"fcm_token"`
		Token     string    `json:"token"`
		CreatedAt time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		DeletedAt gorm.DeletedAt
	}

	AddUserRequest struct {
		Name     string  `json:"name" validate:"required"`
		Email    string  `json:"email" validate:"required,email"`
		Password string  `json:"password" validate:"required"`
		Role     *string `json:"role"`
	}

	UpdateUserRequest struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	UserListResponse struct {
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Email string    `json:"email"`
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	if u.Password != "" {
		u.Password = u.HashPassword(u.Password)
	}

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.ID = uuid.New()
	if u.Password != "" {
		u.Password = u.HashPassword(u.Password)
	}

	return nil
}

func (u *User) HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
