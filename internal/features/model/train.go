package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Train struct {
		ID           uuid.UUID `gorm:"primaryKey"`
		Name         string    `gorm:"varchar(255)"`
		TrainCode    string    `gorm:"varchar(5);unique"`
		TrainClassID uuid.UUID
		CreatedAt    time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		UpdatedAt    time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		DeletedAt    gorm.DeletedAt

		TrainClass TrainClass `gorm:"foreignKey:TrainClassID;references:ID"`
		Route      *Route     `gorm:"foreignKey:TrainID;references:ID"`
	}

	TrainClass struct {
		ID   uuid.UUID `gorm:"primaryKey"`
		Name string    `gorm:"varchar(255)"`
	}
)

func (t *Train) BeforeCreate(_ *gorm.DB) error {
	t.ID = uuid.New()

	return nil
}

func (t *Train) BeforeUpdate(_ *gorm.DB) error {
	t.ID = uuid.New()

	return nil
}

func (tc *TrainClass) BeforeCreate(_ *gorm.DB) error {
	tc.ID = uuid.New()

	return nil
}

func (tc *TrainClass) BeforeUpdate(_ *gorm.DB) error {
	tc.ID = uuid.New()

	return nil
}
