package model

import "github.com/google/uuid"

type (
	TrainClass struct {
		ID   uuid.UUID `gorm:"primaryKey"`
		Name string    `gorm:"varchar(255)"`
	}
)
