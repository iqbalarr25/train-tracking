package database

import (
	"TrainTracking/internal/features/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	RegisterMigration(&gormigrate.Migration{
		ID: "20250619223335_migration_create_station_nodes_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.StationNode{})
		},
		Rollback: func(tx *gorm.DB) error {
			// Define your rollback logic here
			return nil
		},
	})
}
