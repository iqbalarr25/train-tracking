package model

import "github.com/google/uuid"

type (
	Train struct {
		ID           uuid.UUID `gorm:"primaryKey"`
		Name         string    `gorm:"varchar(255)"`
		TrainClassID uuid.UUID

		TrainClass TrainClass `gorm:"foreignKey:TrainClassID;references:ID"`
	}

	TrainClass struct {
		ID   uuid.UUID `gorm:"primaryKey"`
		Name string    `gorm:"varchar(255)"`
	}
)
