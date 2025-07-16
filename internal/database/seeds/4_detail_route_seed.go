package database

import (
	"TrainTracking/internal/features/model"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"time"
)

type InputGapeka struct {
	Sequence     int16              `json:"sequence"`
	TrainName    string             `json:"train_name"`
	TrainCode    string             `json:"train_code"`
	RouteDetails []InputRouteDetail `json:"tracks"`
}

type InputRouteDetail struct {
	Number      string `json:"number"`
	StationCode string `json:"station_code"`
	ArriveTime  string `json:"arrive_time"`
	DepartTime  string `json:"depart_time"`
	Note        string `json:"note"`
	MaxSpeed    *int16 `json:"max_speed"`
}

func init() {
	RegisterSeeder(Seeder{
		Name: "detailRouteSeed",
		Run: func(tx *gorm.DB) error {
			file, err := os.Open("data/data_gapeka.json")
			if err != nil {
				return fmt.Errorf("gagal buka file JSON: %w", err)
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
				}
			}(file)

			var inputGapekas []InputGapeka
			if err := json.NewDecoder(file).Decode(&inputGapekas); err != nil {
				return fmt.Errorf("gagal decode JSON: %w", err)
			}

			for _, input := range inputGapekas {
				var train model.Train
				err = tx.Where("train_code = ?", input.TrainCode).Preload("Route").Find(&train).Error
				if err != nil {
					log.Printf("❌ Gagal menemukan train %s", input.TrainCode)
					continue
				}

				for _, routeDetail := range input.RouteDetails {
					loc, _ := time.LoadLocation("Asia/Jakarta")
					layout := "15:04:05"

					var departTimePtr, arriveTimePtr *time.Time

					if routeDetail.DepartTime != "" && routeDetail.DepartTime != "-" {
						parsedDepart, err := time.ParseInLocation(layout, routeDetail.DepartTime, loc)
						if err != nil {
							log.Printf("❌ Gagal parse waktu keberangkatan %s: %v", routeDetail.DepartTime, err)
							continue
						}
						depart := time.Date(2025, 1, 1, parsedDepart.Hour(), parsedDepart.Minute(), parsedDepart.Second(), 0, loc)
						departTimePtr = &depart
					}

					if routeDetail.ArriveTime != "" && routeDetail.ArriveTime != "Ls" {
						parsedArrive, err := time.ParseInLocation(layout, routeDetail.ArriveTime, loc)
						if err != nil {
							log.Printf("❌ Gagal parse waktu kedatangan %s: %v", routeDetail.ArriveTime, err)
							continue
						}
						arrive := time.Date(2025, 1, 1, parsedArrive.Hour(), parsedArrive.Minute(), parsedArrive.Second(), 0, loc)
						arriveTimePtr = &arrive
					}

					var station model.Station
					result := tx.Where("ref = ?", routeDetail.StationCode).First(&station)
					if result.Error != nil {
						log.Printf("❌ Query gagal untuk stasiun %s: %v", routeDetail.StationCode, result.Error)
						continue
					}
					if result.RowsAffected == 0 {
						log.Printf("⚠️ Stasiun %s tidak ditemukan dalam database", routeDetail.StationCode)
						continue
					}

					sequence, err := strconv.Atoi(routeDetail.Number)
					if err != nil {
						log.Printf("❌ Gagal convert sequence %s: %v", routeDetail.Number, err)
						continue
					}

					var routeDetailModel = model.RouteDetail{
						Sequence:   sequence,
						DepartTime: departTimePtr,
						ArriveTime: arriveTimePtr,
						Note:       routeDetail.Note,
						MaxSpeed:   routeDetail.MaxSpeed,
						RouteID:    train.Route.ID,
						StationID:  station.ID,
					}

					if err := tx.Create(&routeDetailModel).Error; err != nil {
						log.Printf("❌ Gagal insert detail route %s & %s: %v", input.TrainName, routeDetailModel.StationID, err)
						break
					}
				}
			}

			log.Println("✅ Semua data selesai diproses")
			return nil
		},
	})
}
