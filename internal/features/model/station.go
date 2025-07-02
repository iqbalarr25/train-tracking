package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Station struct {
		ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
		OverpassID int64     `json:"overpass_id" json:"overpass_id" gorm:"index"`
		Name       string    `json:"name" gorm:"not null"`
		Ref        string    `json:"ref" gorm:"index"`
		Lat        float64   `json:"lat"`
		Lon        float64   `json:"lon"`
		Geom       string    `json:"geom" gorm:"type:geometry(Point,4326)"`

		RailNodes []*RailNode `json:"rail_nodes" gorm:"many2many:station_nodes"`
	}

	StationOverpass struct {
		ID   int64
		Name string
		Ref  string
		Lat  float64
		Lon  float64
		Geom string
	}
)

func (s *Station) BeforeCreate(_ *gorm.DB) error {
	s.ID = uuid.New()

	return nil
}

func (s *Station) BeforeUpdate(_ *gorm.DB) error {
	s.ID = uuid.New()

	return nil
}
