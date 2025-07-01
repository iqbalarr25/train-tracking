package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	RailEdge struct {
		ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
		Source      uuid.UUID `gorm:"type:uuid;not null"`
		Target      uuid.UUID `gorm:"type:uuid;not null"`
		Cost        float64   `gorm:"not null"` // Panjang/jarak
		ReverseCost float64   `gorm:""`         // Boleh null
		MaxSpeed    *float64  // nullable
		Geom        string    `gorm:"type:geometry(Point,4326)"` // Geometry LineString
	}
)

func (re *RailEdge) BeforeCreate(_ *gorm.DB) error {
	re.ID = uuid.New()

	return nil
}

func (re *RailEdge) BeforeUpdate(_ *gorm.DB) error {
	re.ID = uuid.New()

	return nil
}
