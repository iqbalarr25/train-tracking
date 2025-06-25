package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Route struct {
		ID              uuid.UUID `gorm:"primaryKey"`
		DepartTime      time.Time `json:"depart_time" gorm:"type:timestamp without time zone"`
		ArriveTime      time.Time `json:"arrive_time" gorm:"type:timestamp without time zone"`
		TrainID         uuid.UUID `json:"train_id"`
		DepartStationID int64     `json:"depart_station_id"`
		ArriveStationID int64     `json:"arrive_station_id"`
		CreatedAt       time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		UpdatedAt       time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		DeletedAt       gorm.DeletedAt

		Train         Train   `json:"train" gorm:"foreignKey:TrainID;references:ID"`
		DepartStation Station `json:"depart_station" gorm:"foreignKey:DepartStationID;references:ID"`
		ArriveStation Station `json:"arrive_station" gorm:"foreignKey:ArriveStationID;references:ID"`
	}
)

func (r *Route) BeforeCreate(_ *gorm.DB) error {
	r.ID = uuid.New()

	return nil
}

func (r *Route) BeforeUpdate(_ *gorm.DB) error {
	r.ID = uuid.New()

	return nil
}
