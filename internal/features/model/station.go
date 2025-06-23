package model

type (
	Station struct {
		ID   int64  `gorm:"primaryKey"`
		Name string `gorm:"not null"`
		Ref  string `gorm:"index"`
		Lat  float64
		Lon  float64
		Geom string `gorm:"type:geometry(Point,4326)"`
	
		RailNodes []*RailNode `gorm:"many2many:station_nodes"`
	}
)
