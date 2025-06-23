package database

import (
	"TrainTracking/internal/features/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	RegisterMigration(&gormigrate.Migration{
		ID: "20250623130412_migration_create_routes_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.Route{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&model.Route{})
		},
	})
}
