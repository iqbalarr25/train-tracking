package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Train struct {
		ID           uuid.UUID `gorm:"primaryKey"`
		Name         string    `gorm:"varchar(255)"`
		TrainCode    string    `gorm:"varchar(5);unique"`
		TrainClassID uuid.UUID

		TrainClass TrainClass `gorm:"foreignKey:TrainClassID;references:ID"`
	}

	TrainClass struct {
		ID   uuid.UUID `gorm:"primaryKey"`
		Name string    `gorm:"varchar(255)"`
	}
)

func (t *Train) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New()

	return nil
}

func (t *Train) BeforeUpdate(tx *gorm.DB) error {
	t.ID = uuid.New()

	return nil
}

func (tc *TrainClass) BeforeCreate(tx *gorm.DB) error {
	tc.ID = uuid.New()

	return nil
}

func (tc *TrainClass) BeforeUpdate(tx *gorm.DB) error {
	tc.ID = uuid.New()

	return nil
}
