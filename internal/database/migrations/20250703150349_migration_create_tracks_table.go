package database

import (
	"TrainTracking/internal/features/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	RegisterMigration(&gormigrate.Migration{
		ID:      "20250703150349_migration_create_tracks_table",
		Migrate: func(tx *gorm.DB) error { return tx.AutoMigrate(&model.Track{}) },
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&model.Track{})
		},
	})
}
