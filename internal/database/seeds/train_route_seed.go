package database

import (
	"TrainTracking/internal/features/model"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type Input struct {
	Number         string `json:"number"`
	Code           string `json:"code"`
	Class          string `json:"class"`
	Name           string `json:"name"`
	DepartLocation string `json:"depart_location"`
	ArriveLocation string `json:"arrive_location"`
	DepartTime     string `json:"depart_time"`
	ArriveTime     string `json:"arrive_time"`
	Page           string `json:"page"`
}

func init() {
	RegisterSeeder(Seeder{
		Name: "trainSeed & routeSeed",
		Run: func(db *gorm.DB) error {
			file, err := os.Open("data/train.json")
			if err != nil {
				return fmt.Errorf("gagal buka file JSON: %w", err)
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
				}
			}(file)

			var inputs []Input
			if err := json.NewDecoder(file).Decode(&inputs); err != nil {
				return fmt.Errorf("gagal decode JSON: %w", err)
			}

			for _, input := range inputs {
				tx := db.Begin()
				if tx.Error != nil {
					log.Printf("❌ Gagal memulai transaksi untuk %s: %v", input.Name, tx.Error)
					continue
				}

				var trainClass model.TrainClass
				if err := tx.FirstOrCreate(&trainClass, model.TrainClass{Name: input.Class}).Error; err != nil {
					log.Printf("❌ Gagal buat train class %s: %v", input.Class, err)
					tx.Rollback()
					continue
				}

				train := model.Train{
					Name:         input.Name,
					TrainCode:    input.Code,
					TrainClassID: trainClass.ID,
				}
				if err := tx.Create(&train).Error; err != nil {
					log.Printf("❌ Gagal insert train %s: %v", input.Name, err)
					tx.Rollback()
					continue
				}

				loc, _ := time.LoadLocation("Asia/Jakarta")
				layout := "15:04:05"

				parsedDepart, err := time.ParseInLocation(layout, input.DepartTime, loc)
				if err != nil {
					log.Printf("❌ Gagal parse waktu keberangkatan %s: %v", input.DepartTime, err)
					tx.Rollback()
					continue
				}
				departTime := time.Date(2025, 1, 1, parsedDepart.Hour(), parsedDepart.Minute(), parsedDepart.Second(), 0, loc)

				parsedArrive, err := time.ParseInLocation(layout, input.ArriveTime, loc)
				if err != nil {
					log.Printf("❌ Gagal parse waktu kedatangan %s: %v", input.ArriveTime, err)
					tx.Rollback()
					continue
				}
				arriveTime := time.Date(2025, 1, 1, parsedArrive.Hour(), parsedArrive.Minute(), parsedArrive.Second(), 0, loc)

				var departStation model.Station
				if err := tx.Where("ref = ?", input.DepartLocation).First(&departStation).Error; err != nil {
					log.Printf("❌ Stasiun keberangkatan %s tidak ditemukan: %v", input.DepartLocation, err)
					tx.Rollback()
					continue
				}

				var arriveStation model.Station
				if err := tx.Where("ref = ?", input.ArriveLocation).First(&arriveStation).Error; err != nil {
					log.Printf("❌ Stasiun kedatangan %s tidak ditemukan: %v", input.ArriveLocation, err)
					tx.Rollback()
					continue
				}

				route := model.Route{
					DepartTime:      departTime,
					ArriveTime:      arriveTime,
					TrainID:         train.ID,
					DepartStationID: departStation.ID,
					ArriveStationID: arriveStation.ID,
				}
				if err := tx.Create(&route).Error; err != nil {
					log.Printf("❌ Gagal insert route untuk %s: %v", input.Name, err)
					tx.Rollback()
					continue
				}

				if err := tx.Commit().Error; err != nil {
					log.Printf("❌ Gagal commit transaksi untuk %s: %v", input.Name, err)
					tx.Rollback()
					continue
				}
			}

			log.Println("✅ Semua data selesai diproses")
			return nil
		},
	})
}
