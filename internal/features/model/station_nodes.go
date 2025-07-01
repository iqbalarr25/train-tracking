package model

import "github.com/google/uuid"

type (
	StationNode struct {
		StationID uuid.UUID `gorm:"type:uuid;primaryKey"`
		NodeID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	}
)
