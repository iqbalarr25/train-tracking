package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Train struct {
		ID           uuid.UUID      `json:"id" gorm:"primaryKey"`
		Name         string         `json:"name" gorm:"varchar(255)"`
		TrainCode    string         `json:"train_code" gorm:"varchar(5);unique"`
		Status       string         `json:"status" gorm:"varchar(255);not null;default:Inactive"`
		TrainClassID uuid.UUID      `json:"train_class_id"`
		CreatedAt    time.Time      `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		UpdatedAt    time.Time      `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		DeletedAt    gorm.DeletedAt `json:"deleted_at"`

		TrainClass TrainClass `json:"train_class" gorm:"foreignKey:TrainClassID;references:ID"`
		Route      *Route     `json:"route" gorm:"foreignKey:TrainID;references:ID"`
	}

	TrainClass struct {
		ID   uuid.UUID `json:"id" gorm:"primaryKey"`
		Name string    `json:"name" gorm:"varchar(255)"`
	}

	TrainListResponse struct {
		ID                 uuid.UUID              `json:"id"`
		Name               string                 `json:"name"`
		TrainCode          string                 `json:"train_code"`
		Status             string                 `json:"status"`
		TrainClass         TrainClass             `json:"train_class"`
		TrainRouteResponse TrainListRouteResponse `json:"route"`
	}

	TrainListRouteResponse struct {
		ID            uuid.UUID                `json:"id"`
		DepartTime    time.Time                `json:"depart_time"`
		ArriveTime    time.Time                `json:"arrive_time"`
		DepartStation TrainListStationResponse `json:"depart_station"`
		ArriveStation TrainListStationResponse `json:"arrive_station"`
	}

	TrainListStationResponse struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Ref  string    `json:"ref"`
	}

	GetTrainPositionResponse struct {
		Lat      float64 `json:"lat"`
		Lon      float64 `json:"lon"`
		Progress float64 `json:"progress"`
		Status   string  `json:"status"`
	}

	TrainPositionRouteDetailSummary struct {
		DepartTime      time.Time `json:"depart_time"`
		ArriveTime      time.Time `json:"arrive_time"`
		ID              uuid.UUID `json:"id"`
		CurrentSequence int64     `json:"current_sequence"`
	}

	TrainLatLon struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	}
)

func (t *Train) BeforeCreate(_ *gorm.DB) error {
	t.ID = uuid.New()

	return nil
}

func (t *Train) BeforeUpdate(_ *gorm.DB) error {
	t.ID = uuid.New()

	return nil
}

func (tc *TrainClass) BeforeCreate(_ *gorm.DB) error {
	tc.ID = uuid.New()

	return nil
}

func (tc *TrainClass) BeforeUpdate(_ *gorm.DB) error {
	tc.ID = uuid.New()

	return nil
}
