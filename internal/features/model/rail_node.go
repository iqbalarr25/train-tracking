package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	RailNode struct {
		ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
		Lat       float64   `json:"lat"`
		Lon       float64   `json:"lon"`
		NodeIntID int       `json:"node_int_id" gorm:"->"`
		Geom      string    `json:"geom" gorm:"type:geometry(Point,4326)"`

		Stations []*Station `json:"stations" gorm:"many2many:station_nodes"`
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
