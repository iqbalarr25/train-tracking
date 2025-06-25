package database

import (
	"TrainTracking/internal/features/model"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"math"
	"os"
	"strconv"
)

type OverpassData struct {
	Elements []struct {
		Type  string            `json:"type"`
		ID    int64             `json:"id"`
		Lat   float64           `json:"lat,omitempty"`
		Lon   float64           `json:"lon,omitempty"`
		Nodes []int64           `json:"nodes,omitempty"`
		Tags  map[string]string `json:"tags,omitempty"`
	} `json:"elements"`
}

func init() {
	RegisterSeeder(Seeder{
		Name: "railSeed",
		Run: func(tx *gorm.DB) error {
			raw, err := os.ReadFile("data/rail.json")
			if err != nil {
				log.Fatal("read file error:", err)
			}

			var parsed OverpassData
			if err := json.Unmarshal(raw, &parsed); err != nil {
				log.Fatal("parse json error:", err)
			}

			nodeMap := map[int64]model.RailNode{}

			// --- Insert Nodes ---
			for _, el := range parsed.Elements {
				if el.Type == "node" {
					node := model.RailNode{ID: el.ID, Lat: el.Lat, Lon: el.Lon}
					tx.Exec(`
					  INSERT INTO rail_nodes (id, lat, lon, geom)
					  VALUES (?, ?, ?, ST_SetSRID(ST_MakePoint(?, ?), 4326))
					`, node.ID, node.Lat, node.Lon, node.Lon, node.Lat)
					nodeMap[el.ID] = node
				}
			}

			for _, el := range parsed.Elements {
				if el.Type == "way" && len(el.Nodes) >= 2 {
					var maxSpeed *float64
					if ms, ok := el.Tags["maxspeed"]; ok {
						if parsed, err := strconv.ParseFloat(ms, 64); err == nil {
							maxSpeed = &parsed
						}
					}

					for i := 0; i < len(el.Nodes)-1; i++ {
						src := nodeMap[el.Nodes[i]]
						dst := nodeMap[el.Nodes[i+1]]

						if src.ID == 0 || dst.ID == 0 {
							continue
						}

						edge := model.RailEdge{
							Source:      src.ID,
							Target:      dst.ID,
							Cost:        haversine(src.Lat, src.Lon, dst.Lat, dst.Lon),
							ReverseCost: haversine(dst.Lat, dst.Lon, src.Lat, src.Lon),
							MaxSpeed:    maxSpeed,
						}

						tx.Exec(`
							INSERT INTO rail_edges (source, target, cost, reverse_cost, max_speed, geom)
							VALUES (?, ?, ?, ?, ?, ST_SetSRID(ST_MakeLine(
								ST_MakePoint(?, ?),
								ST_MakePoint(?, ?)
							), 4326))
						`, edge.Source, edge.Target, edge.Cost, edge.ReverseCost, edge.MaxSpeed, src.Lon, src.Lat, dst.Lon, dst.Lat)
					}
				}
			}

			// --- Insert Stations ---x
			for _, el := range parsed.Elements {
				if el.Type == "node" && (el.Tags["railway"] == "station" || el.Tags["public_transport"] == "station") {
					name := el.Tags["name"]
					ref := el.Tags["railway:ref"]
					station := model.Station{
						ID:   el.ID,
						Lat:  el.Lat,
						Lon:  el.Lon,
						Name: name,
						Ref:  ref,
					}

					tx.Exec(`
					  INSERT INTO stations (id, name, ref, lat, lon, geom)
					  VALUES (?, ?, ?, ?, ?, ST_SetSRID(ST_MakePoint(?, ?), 4326))
					`, station.ID, station.Name, station.Ref, station.Lat, station.Lon, station.Lon, station.Lat)
				}
			}

			// --- Insert Station <-> Node Ref Many-to-Many ---
			for _, el := range parsed.Elements {
				if el.Type == "node" && (el.Tags["railway"] == "stop" || el.Tags["public_transport"] == "stop_position") {
					ref := el.Tags["ref"]
					railRef := el.Tags["railway:ref"]
					name := el.Tags["name"]

					var station model.Station
					var err error

					// Coba cari stasiun pakai ref
					if ref != "" {
						err = tx.Raw(`SELECT id FROM stations WHERE ref = ? LIMIT 1`, ref).Scan(&station).Error
						if err != nil {
							log.Printf("error querying station by ref (%s): %v", ref, err)
							continue
						}
					}

					// Kalau belum ketemu dan masih kosong, coba pakai railway:ref
					if station.ID == 0 && railRef != "" {
						err = tx.Raw(`SELECT id FROM stations WHERE ref = ? LIMIT 1`, railRef).Scan(&station).Error
						if err != nil {
							log.Printf("error querying station by railway:ref (%s): %v", railRef, err)
							continue
						}
					}

					// Kalau masih belum ketemu dan ada name, coba pakai name
					if station.ID == 0 && name != "" {
						err = tx.Raw(`SELECT id FROM stations WHERE name ILIKE ? LIMIT 1`, name).Scan(&station).Error
						if err != nil {
							log.Printf("error querying station by name (%s): %v", name, err)
							continue
						}
					}

					// Kalau gak ketemu, skip
					if station.ID == 0 {
						continue
					}

					// Masukkan relasi station_id dan node_id
					err = tx.Exec(`
						INSERT INTO station_nodes (station_id, node_id)
						VALUES (?, ?)
						ON CONFLICT (station_id, node_id) DO NOTHING
					`, station.ID, el.ID).Error

					if err != nil {
						log.Printf("failed to insert station_node (%d, %d): %v", station.ID, el.ID, err)
					}
				}
			}

			log.Println("✅ rail seeded successfully")
			return nil
		},
	})
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371e3
	φ1 := lat1 * math.Pi / 180
	φ2 := lat2 * math.Pi / 180
	dφ := (lat2 - lat1) * math.Pi / 180
	dλ := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(dφ/2)*math.Sin(dφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*math.Sin(dλ/2)*math.Sin(dλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
