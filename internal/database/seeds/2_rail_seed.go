package database

import (
	"TrainTracking/internal/features/model"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"math"
	"os"
	"strconv"
)

type OverpassElement struct {
	Type  string            `json:"type"`
	ID    int64             `json:"id"`
	Lat   float64           `json:"lat,omitempty"`
	Lon   float64           `json:"lon,omitempty"`
	Nodes []int64           `json:"nodes,omitempty"`
	Tags  map[string]string `json:"tags,omitempty"`
}

type OverpassData struct {
	Elements []OverpassElement `json:"elements"`
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

			nodeMap := map[int64]uuid.UUID{}
			nodeIntMap := map[int64]int{}

			// --- Insert Nodes ---
			for _, el := range parsed.Elements {
				if el.Type == "node" {
					node := model.RailNode{
						Lat:  el.Lat,
						Lon:  el.Lon,
						Geom: fmt.Sprintf("SRID=4326;POINT(%f %f)", el.Lon, el.Lat),
					}
					if err := tx.Create(&node).Error; err != nil {
						return fmt.Errorf("insert node error: %w", err)
					}

					nodeMap[el.ID] = node.ID

					// ambil node_int_id
					var intID int
					err = tx.Raw("SELECT node_int_id FROM rail_nodes WHERE id = ?", node.ID).Scan(&intID).Error
					if err != nil {
						return fmt.Errorf("ambil node_int_id gagal: %w", err)
					}
					nodeIntMap[el.ID] = intID
				}
			}

			// --- Insert Edges ---
			for _, el := range parsed.Elements {
				if el.Type == "way" && len(el.Nodes) >= 2 {
					var maxSpeed *float64
					if ms, ok := el.Tags["maxspeed"]; ok {
						if parsedSpeed, err := strconv.ParseFloat(ms, 64); err == nil {
							maxSpeed = &parsedSpeed
						}
					}

					for i := 0; i < len(el.Nodes)-1; i++ {
						srcUUID := nodeMap[el.Nodes[i]]
						dstUUID := nodeMap[el.Nodes[i+1]]

						srcInt := nodeIntMap[el.Nodes[i]]
						dstInt := nodeIntMap[el.Nodes[i+1]]

						src := parsed.FindNodeByID(el.Nodes[i])
						dst := parsed.FindNodeByID(el.Nodes[i+1])
						if src == nil || dst == nil {
							continue
						}

						edge := model.RailEdge{
							Source:      srcUUID,
							Target:      dstUUID,
							SourceInt:   srcInt,
							TargetInt:   dstInt,
							Cost:        haversine(src.Lat, src.Lon, dst.Lat, dst.Lon),
							ReverseCost: haversine(dst.Lat, dst.Lon, src.Lat, src.Lon),
							MaxSpeed:    maxSpeed,
							Geom:        fmt.Sprintf("SRID=4326;LINESTRING(%f %f, %f %f)", src.Lon, src.Lat, dst.Lon, dst.Lat),
						}
						if err := tx.Create(&edge).Error; err != nil {
							log.Printf("insert edge failed: %v", err)
						}
					}
				}
			}

			stationMap := map[int64]uuid.UUID{}

			// --- Insert Stations ---
			for _, el := range parsed.Elements {
				if el.Type == "node" && (el.Tags["railway"] == "station" || el.Tags["public_transport"] == "station") {
					name := el.Tags["name"]
					ref := el.Tags["ref"]
					railRef := el.Tags["railway:ref"]

					station := model.Station{
						Lat:        el.Lat,
						Lon:        el.Lon,
						Name:       name,
						OverpassID: el.ID,
						Geom:       fmt.Sprintf("SRID=4326;POINT(%f %f)", el.Lon, el.Lat),
					}

					if ref != "" {
						station.Ref = ref
					} else if railRef != "" {
						station.Ref = railRef
					}

					if err := tx.Create(&station).Error; err != nil {
						return fmt.Errorf("insert station error: %w", err)
					}

					stationMap[el.ID] = station.ID
				}
			}

			file, err := os.Open("data/station_update.json")
			if err != nil {
				return fmt.Errorf("gagal buka file JSON: %w", err)
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
				}
			}(file)

			// --- Update Station Refs dari station_update.json ---
			file, err = os.Open("data/station_update.json")
			if err != nil {
				return fmt.Errorf("gagal buka file JSON: %w", err)
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
				}
			}(file)

			var updatedStations []model.StationOverpass
			if err := json.NewDecoder(file).Decode(&updatedStations); err != nil {
				return fmt.Errorf("gagal decode JSON: %w", err)
			}

			for _, station := range updatedStations {
				err = tx.Table("stations").Where("overpass_id = ?", station.ID).
					Update("ref", station.Ref).Error
				if err != nil {
					return fmt.Errorf("gagal update stasiun %s: %w", station.ID, err)
				}
			}

			// --- Insert Station <-> Node (Many-to-Many) ---
			for _, el := range parsed.Elements {
				if el.Type == "node" && (el.Tags["railway"] == "stop" || el.Tags["public_transport"] == "stop_position") {
					ref := el.Tags["ref"]
					railRef := el.Tags["railway:ref"]
					name := el.Tags["name"]

					var station model.Station
					found := false

					if ref != "" {
						err = tx.Where("ref = ?", ref).First(&station).Error
						found = err == nil
					}
					if !found && railRef != "" {
						err = tx.Where("ref = ?", railRef).First(&station).Error
						found = err == nil
					}
					if !found && name != "" {
						err = tx.Where("name ILIKE ?", name).First(&station).Error
						found = err == nil
					}
					if !found {
						continue
					}

					nodeUUID, ok := nodeMap[el.ID]
					if !ok {
						continue
					}

					err = tx.Exec(`
						INSERT INTO station_nodes (station_id, node_id)
						VALUES (?, ?)
						ON CONFLICT (station_id, node_id) DO NOTHING
					`, station.ID, nodeUUID).Error
					if err != nil {
						log.Printf("failed to insert station_node (%s, %s): %v", station.ID, nodeUUID, err)
					}
				}
			}

			log.Println("✅ rail seeded successfully")
			return nil
		},
	})
}

func (data *OverpassData) FindNodeByID(id int64) *OverpassElement {
	for _, el := range data.Elements {
		if el.Type == "node" && el.ID == id {
			return &el
		}
	}
	return nil
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
