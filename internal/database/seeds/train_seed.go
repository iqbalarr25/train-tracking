package database

import (
	"TrainTracking/internal/features/model"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
)

type TrainInput struct {
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
		Name: "trainSeed",
		Run: func(tx *gorm.DB) error {
			file, err := os.Open("data/train.json")
			if err != nil {
				return fmt.Errorf("gagal buka file: %w", err)
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {

				}
			}(file)

			var inputs []TrainInput
			decoder := json.NewDecoder(file)
			if err := decoder.Decode(&inputs); err != nil {
				return fmt.Errorf("gagal decode json: %w", err)
			}

			for _, input := range inputs {
				var trainClass model.TrainClass
				err := tx.FirstOrCreate(&trainClass, model.TrainClass{
					Name: input.Name,
				}).Error

				if err != nil {
					return fmt.Errorf("failed open train class: %w", err)
				}

				train := model.Train{
					Name:         input.Name,
					TrainCode:    input.Code,
					TrainClassID: trainClass.ID,
				}

				if err := tx.Create(&train).Error; err != nil {
					log.Printf("❌ Failed insert train %s: %v", input.Name, err)
					continue
				}
			}

			log.Println("✅ train seeded successfully")
			return nil
		},
	})
}
