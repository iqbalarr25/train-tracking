package database

import (
	"TrainTracking/internal/features/model"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"log"
	"math"
	"net/http"
	"strings"
)

type OsrmData struct {
	Code      string     `json:"code"`
	Routes    []Routes   `json:"routes"`
	Waypoints []Waypoint `json:"waypoints"`
}

type Routes struct {
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Coordinates [][]float64 `json:"coordinates"`
	Type        string      `json:"type"`
}

type Waypoint struct {
	Hint     string    `json:"hint"`
	Distance float64   `json:"distance"`
	Name     string    `json:"name"`
	Location []float64 `json:"location"`
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
				var routeDetails []model.RouteDetail
				if err := tx.Where("route_id = ?", route.ID).
					Order("sequence ASC").
					Preload("Station").
					Find(&routeDetails).Error; err != nil {
					log.Println("❌ Gagal ambil data route details:", err)
					continue
				}

				var coords []string
				for _, rd := range routeDetails {
					if rd.Station.Lon != 0 && rd.Station.Lat != 0 {
						coord := fmt.Sprintf("%.14f,%.14f", rd.Station.Lon, rd.Station.Lat)
						coords = append(coords, coord)
					}
				}

				coordString := strings.Join(coords, ";")

				url := fmt.Sprintf(
					"https://mule-open-titmouse.ngrok-free.app/route/v1/train/%s?alternatives=true&overview=full&geometries=geojson",
					coordString,
				)
				fmt.Println(url)

				resp, err := http.Get(url)
				if err != nil {
					fmt.Println("❌ Error saat request:", err)
					return err
				}
				defer func(Body io.ReadCloser) {
					err := Body.Close()
					if err != nil {
					}
				}(resp.Body)

				var osrmData OsrmData
				if err := json.NewDecoder(resp.Body).Decode(&osrmData); err != nil {
					return fmt.Errorf("gagal decode JSON: %w", err)
				}

				if len(osrmData.Routes) == 0 {
					log.Printf("⚠️ Tidak ada route ditemukan dari OSRM untuk %s", route.ID)
					continue
				}

				geometry := osrmData.Routes[0].Geometry
				waypoints := osrmData.Waypoints

				index := 0
				for i, coord := range geometry.Coordinates {
					var stationID *uuid.UUID
					if i == 0 {
						stationID = &routeDetails[index].StationID
					} else if index+1 < len(waypoints) &&
						isSameCoordinate(coord, waypoints[index+1].Location) &&
						index+1 < len(routeDetails) {
						index++
						stationID = &routeDetails[index].StationID
					}

					var cost float64
					if i > 0 {
						prev := geometry.Coordinates[i-1]
						lat1, lon1 := prev[1], prev[0]
						lat2, lon2 := coord[1], coord[0]
						cost = haversine(lat1, lon1, lat2, lon2)
					}

					track := model.Track{
						Sequence:      i + 1,
						Longitude:     coord[0],
						Latitude:      coord[1],
						RouteDetailID: routeDetails[index].ID,
						MaxSpeed:      routeDetails[index].MaxSpeed,
						StationID:     stationID,
						Cost:          cost,
					}

					if err := tx.Create(&track).Error; err != nil {
						log.Printf("❌ Gagal insert track untuk route %s, index %d: %v", route.ID, i, err)
					}
				}
			}

			return nil
		},
	})
}

func isSameCoordinate(coord1, coord2 []float64) bool {
	if len(coord1) != 2 || len(coord2) != 2 {
		return false
	}
	const tolerance = 1e-6
	return math.Abs(coord1[0]-coord2[0]) < tolerance && math.Abs(coord1[1]-coord2[1]) < tolerance
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // Radius bumi dalam meter
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1Rad)*math.Cos(lat2Rad)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
