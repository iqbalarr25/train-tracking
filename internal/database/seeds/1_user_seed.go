package database

import (
	"TrainTracking/internal/features/model"
	"TrainTracking/pkg/helper"
	"gorm.io/gorm"
	"log"
)

func init() {
	RegisterSeeder(Seeder{
		Name: "userSeed",
		Run: func(tx *gorm.DB) error {
			items := []model.User{
				{
					Email:    "admin@transtrack.id",
					Name:     "admin",
					Password: "password",
					Role:     helper.UserRoleAdmin,
				},
			}

			// You can use tx.Model(&items) or bellow:
			if err := tx.Table("users").Create(&items).Error; err != nil {
				return err
			}

			log.Println("✅ user seeded successfully")
			return nil
		},
	})
}
