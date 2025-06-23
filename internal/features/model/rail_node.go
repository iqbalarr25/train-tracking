package model

type (
	RailNode struct {
		ID   int64 `gorm:"primaryKey"`
		Lat  float64
		Lon  float64
		Geom string `gorm:"type:geometry(Point,4326)"`

		Stations []*Station `gorm:"many2many:station_nodes"`
	}
)
