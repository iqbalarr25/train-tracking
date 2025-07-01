package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Station struct {
		ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
		OverpassID int64     `json:"overpass_id" gorm:"index"`
		Name       string    `gorm:"not null"`
		Ref        string    `gorm:"index"`
		Lat        float64
		Lon        float64
		Geom       string `gorm:"type:geometry(Point,4326)"`

		RailNodes []*RailNode `gorm:"many2many:station_nodes"`
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
