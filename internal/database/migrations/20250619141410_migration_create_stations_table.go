package database

import (
	"TrainTracking/internal/features/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	RegisterMigration(&gormigrate.Migration{
		ID: "20250619141410_migration_create_stations_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.Station{})
		},
		Rollback: func(tx *gorm.DB) error {
			// Define your rollback logic here
			return tx.Migrator().DropTable(&model.Station{})
		},
	})
}
