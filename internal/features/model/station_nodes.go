package model

type (
	StationNode struct {
		StationID int64 `gorm:"primaryKey"`
		NodeID    int64 `gorm:"primaryKey"`
	}
)
