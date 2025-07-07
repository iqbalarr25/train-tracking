package database

import (
	"TrainTracking/internal/features/model"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type PgRoutingData struct {
	OrderFrom     int16      `gorm:"column:order_from"`
	Sequence      int        `gorm:"column:seq"`
	Latitude      float64    `gorm:"column:lat"`
	Longitude     float64    `gorm:"column:lon"`
	MaxSpeed      *float64   `gorm:"column:max_speed"`
	Cost          float64    `gorm:"column:cost"`
	IsStation     bool       `gorm:"column:is_station"`
	RailNodeID    uuid.UUID  `gorm:"column:rail_node_id"`
	RouteDetailID uuid.UUID  `gorm:"column:route_detail_id"`
	StationID     *uuid.UUID `gorm:"column:station_id"`
}

func init() {
	RegisterSeeder(Seeder{
		Name: "TrackSeed",
		Run: func(tx *gorm.DB) error {
			var routes []model.Route
			if err := tx.Find(&routes).Error; err != nil {
				log.Println("❌ Gagal ambil data routes:", err)
			}

			for _, route := range routes {
				var data []PgRoutingData

				const sqlQuery = `
				WITH station_sequence AS (
				  SELECT rd.sequence AS ordering, rd.station_id
				  FROM route_details rd
				  WHERE rd.route_id = @route_id
				),
				station_main_nodes AS (
				  SELECT DISTINCT ON (sn.station_id)
					ss.ordering, sn.station_id, sn.node_id
				  FROM station_sequence ss
				  JOIN station_nodes sn ON sn.station_id = ss.station_id
				  ORDER BY sn.station_id, sn.node_id
				),
				station_pairs AS (
				  SELECT s1.ordering AS order_from, s2.ordering AS order_to,
						 s1.node_id AS source, s2.node_id AS target
				  FROM station_main_nodes s1
				  JOIN station_main_nodes s2 ON s2.ordering = s1.ordering + 1
				),
				station_pairs_int AS (
				  SELECT sp.order_from,
						 rn1.node_int_id AS source_int,
						 rn2.node_int_id AS target_int
				  FROM station_pairs sp
				  JOIN rail_nodes rn1 ON rn1.id = sp.source
				  JOIN rail_nodes rn2 ON rn2.id = sp.target
				),
				paths AS (
				  SELECT sp.order_from, d.seq, d.node, d.edge, d.cost
				  FROM station_pairs_int sp,
					   LATERAL (
						 SELECT * FROM pgr_dijkstra(
						   'SELECT edge_int_id AS id, source_int AS source, target_int AS target, cost, reverse_cost FROM rail_edges',
						   sp.source_int, sp.target_int, false
						 )
					   ) AS d
				),
				paths_with_prev AS (
				  SELECT p.*, LAG(p.node) OVER (PARTITION BY p.order_from ORDER BY p.seq) AS prev_node
				  FROM paths p
				),
				segment_speed AS (
				  SELECT rd.sequence AS order_from, rd.max_speed, rd.id AS route_detail_id
				  FROM route_details rd
				  WHERE rd.route_id = @route_id
				),
				path_result AS (
				  SELECT p.order_from, p.seq, n.id AS rail_node_id, n.lat, n.lon,
						 ss.max_speed, ss.route_detail_id, p.cost,
						 CASE WHEN sn.station_id IS NOT NULL THEN true ELSE false END AS is_station,
						 sn.station_id
				  FROM paths_with_prev p
				  JOIN rail_nodes n ON p.node = n.node_int_id
				  LEFT JOIN station_nodes sn ON sn.node_id = n.id
				  LEFT JOIN segment_speed ss ON ss.order_from = p.order_from
				)
				SELECT * FROM path_result
				ORDER BY order_from, seq
				`

				if err := tx.Raw(sqlQuery, sql.Named("route_id", route.ID)).Scan(&data).Error; err != nil {
					return err
				}

				for _, v := range data {
					fmt.Print(v)
					track := model.Track{
						Sequence:      v.Sequence,
						Latitude:      v.Latitude,
						Longitude:     v.Longitude,
						Cost:          v.Cost,
						MaxSpeed:      v.MaxSpeed,
						RouteDetailID: v.RouteDetailID,
						StationID:     v.StationID,
					}

					err := tx.Create(&track).Error
					if err != nil {
						return err
					}
				}
			}

			log.Println("✅ Semua data selesai diproses")
			return nil
		},
	})
}
