package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Track struct {
		ID            uuid.UUID      `json:"id" gorm:"primaryKey"`
		Sequence      int            `json:"sequence" gorm:"not null"`
		Latitude      float64        `json:"latitude" gorm:"not null"`
		Longitude     float64        `json:"longitude" gorm:"not null"`
		Cost          float64        `json:"cost" gorm:"not null"`
		MaxSpeed      *float64       `json:"max_speed"`
		RouteDetailID uuid.UUID      `json:"route_detail_id" gorm:"not null;index"`
		StationID     *uuid.UUID     `json:"station_id" gorm:"index"`
		CreatedAt     time.Time      `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		UpdatedAt     time.Time      `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		DeletedAt     gorm.DeletedAt `json:"deleted_at"`

		RouteDetail RouteDetail `json:"route_detail" gorm:"foreignKey:RouteDetailID;references:ID"`
		Station     *Station    `json:"station" gorm:"foreignKey:StationID;references:ID"`
	}
)

func (t *Track) BeforeCreate(_ *gorm.DB) error {
	t.ID = uuid.New()

	return nil
}

func (t *Track) BeforeUpdate(_ *gorm.DB) error {
	t.ID = uuid.New()

	return nil
}
