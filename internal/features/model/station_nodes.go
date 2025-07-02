package model

import "github.com/google/uuid"

type (
	StationNode struct {
		StationID uuid.UUID `json:"station_id" gorm:"type:uuid;primaryKey"`
		NodeID    uuid.UUID `json:"node_id" gorm:"type:uuid;primaryKey"`
	}
)
