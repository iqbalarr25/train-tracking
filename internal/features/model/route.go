package model

import (
	"github.com/google/uuid"
	"time"
)

type (
	Route struct {
		ID              uuid.UUID `gorm:"primaryKey"`
		Name            string    `gorm:"varchar(255)"`
		TrainID         uuid.UUID `json:"train_id"`
		TimeDepart      time.Time `json:"time_depart"`
		TimeArrive      time.Time `json:"time_arrive"`
		DepartStationID uuid.UUID `json:"depart_station_id"`
		ArriveStationID uuid.UUID `json:"arrive_station_id"`

		Train         Train   `json:"train" gorm:"foreignKey:TrainID;references:ID"`
		DepartStation Station `json:"depart_station" gorm:"foreignKey:DepartStationID;references:ID"`
		ArriveStation Station `json:"arrive_station" gorm:"foreignKey:ArriveStationID;references:ID"`
	}
)
