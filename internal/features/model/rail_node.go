package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	RailNode struct {
		ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
		Lat  float64
		Lon  float64
		Geom string `gorm:"type:geometry(Point,4326)"`

		Stations []*Station `gorm:"many2many:station_nodes"`
	}
)

func (rn *RailNode) BeforeCreate(_ *gorm.DB) error {
	rn.ID = uuid.New()

	return nil
}

func (rn *RailNode) BeforeUpdate(_ *gorm.DB) error {
	rn.ID = uuid.New()

	return nil
}
