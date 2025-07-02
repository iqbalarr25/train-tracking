package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	RailEdge struct {
		ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
		Source      uuid.UUID `json:"source" gorm:"type:uuid;not null"`
		Target      uuid.UUID `json:"target" gorm:"type:uuid;not null"`
		SourceInt   int       `json:"source_int"`
		TargetInt   int       `json:"target_int"`
		EdgeIntId   int       `json:"edge_int_id" gorm:"->"`
		Cost        float64   `json:"cost" gorm:"not null"`
		ReverseCost float64   `json:"reverse_cost"`
		MaxSpeed    *float64  `json:"max_speed"`
		Geom        string    `json:"geom" gorm:"type:geometry(Point,4326)"`
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
