package database

import (
	"TrainTracking/internal/features/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	RegisterMigration(&gormigrate.Migration{
		ID: "20250625171448_migration_create_route_details_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.RouteDetail{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&model.RouteDetail{})
		},
	})
}
