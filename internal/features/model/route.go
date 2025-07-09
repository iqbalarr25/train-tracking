package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Route struct {
		ID              uuid.UUID      `json:"id" gorm:"primaryKey"`
		DepartTime      time.Time      `json:"depart_time" gorm:"type:timestamp without time zone"`
		ArriveTime      time.Time      `json:"arrive_time" gorm:"type:timestamp without time zone"`
		TrainID         uuid.UUID      `json:"train_id"`
		DepartStationID uuid.UUID      `json:"depart_station_id"`
		ArriveStationID uuid.UUID      `json:"arrive_station_id"`
		CreatedAt       time.Time      `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		UpdatedAt       time.Time      `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		DeletedAt       gorm.DeletedAt `json:"deleted_at"`

		Train         Train         `json:"train" gorm:"foreignKey:TrainID;references:ID"`
		DepartStation Station       `json:"depart_station" gorm:"foreignKey:DepartStationID;references:ID"`
		ArriveStation Station       `json:"arrive_station" gorm:"foreignKey:ArriveStationID;references:ID"`
		RouteDetails  []RouteDetail `json:"route_details" gorm:"foreignKey:RouteID;references:ID"`
	}

	RouteDetail struct {
		ID         uuid.UUID  `json:"id" gorm:"primaryKey"`
		Sequence   int        `json:"sequence"`
		DepartTime *time.Time `json:"depart_time,omitempty" gorm:"type:timestamp without time zone"`
		ArriveTime *time.Time `json:"arrive_time,omitempty" gorm:"type:timestamp without time zone"`
		Note       string     `json:"note" gorm:"type:text"`
		MaxSpeed   *int16     `json:"max_speed,omitempty" gorm:"type:float"`

		RouteID   uuid.UUID      `json:"route_id"`
		StationID uuid.UUID      `json:"station_id"`
		CreatedAt time.Time      `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		UpdatedAt time.Time      `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		DeletedAt gorm.DeletedAt `json:"deleted_at"`

		Route   Route   `json:"route" gorm:"foreignKey:RouteID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
		Station Station `json:"station" gorm:"foreignKey:StationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
		Tracks  []Track `json:"tracks" gorm:"foreignKey:RouteDetailID;references:ID"`
	}

	GetRouteResponse struct {
		ID         uuid.UUID                `json:"id" gorm:"primaryKey"`
		From       string                   `json:"from"`
		To         string                   `json:"to"`
		DepartTime *time.Time               `json:"depart_time"`
		ArriveTime *time.Time               `json:"arrive_time"`
		Detail     []GetRouteDetailResponse `json:"details"`
	}

	GetRouteDetailResponse struct {
		ID         uuid.UUID               `json:"id" gorm:"primaryKey"`
		From       string                  `json:"from"`
		To         string                  `json:"to"`
		DepartTime *time.Time              `json:"depart_time"`
		ArriveTime *time.Time              `json:"arrive_time"`
		Track      []GetRouteTrackResponse `json:"tracks"`
	}

	GetRouteTrackResponse struct {
		ID        uuid.UUID  `json:"id" gorm:"primaryKey"`
		Sequence  int        `json:"sequence"`
		Cost      float32    `json:"cost"`
		MaxSpeed  *int16     `json:"max_speed"`
		Latitude  float64    `json:"latitude"`
		Longitude float64    `json:"longitude"`
		StationID *uuid.UUID `json:"station_id"`
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

func (dr *RouteDetail) BeforeCreate(_ *gorm.DB) error {
	dr.ID = uuid.New()

	return nil
}

func (dr *RouteDetail) BeforeUpdate(_ *gorm.DB) error {
	dr.ID = uuid.New()

	return nil
}
