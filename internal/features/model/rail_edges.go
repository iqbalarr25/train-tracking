package model

type (
	RailEdge struct {
		ID          int64    `gorm:"primaryKey;autoIncrement"`
		Source      int64    `gorm:"not null"` // OSM node ID
		Target      int64    `gorm:"not null"` // OSM node ID
		Cost        float64  `gorm:"not null"` // Panjang/jarak
		ReverseCost float64  `gorm:""`         // Boleh null
		MaxSpeed    *float64 // nullable
		Geom        string   `gorm:"type:geometry(Point,4326)"` // Geometry LineString

	}
)
