package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Train struct {
		ID           uuid.UUID `gorm:"primaryKey"`
		Name         string    `gorm:"varchar(255)"`
		TrainCode    string    `gorm:"varchar(5);unique"`
		Status       string    `gorm:"varchar(255);not null;default:INACTIVE"`
		TrainClassID uuid.UUID
		CreatedAt    time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		UpdatedAt    time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
		DeletedAt    gorm.DeletedAt

		TrainClass TrainClass `gorm:"foreignKey:TrainClassID;references:ID"`
		Route      *Route     `gorm:"foreignKey:TrainID;references:ID"`
	}

	TrainClass struct {
		ID   uuid.UUID `gorm:"primaryKey"`
		Name string    `gorm:"varchar(255)"`
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
		ID  int64  `json:"id"`
		Ref string `json:"ref"`
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
