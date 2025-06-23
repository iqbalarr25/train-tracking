package database

import (
	"TrainTracking/internal/features/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	RegisterMigration(&gormigrate.Migration{
		ID: "20250623113631_migration_create_train_class_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.TrainClass{})
		},
		Rollback: func(tx *gorm.DB) error {
			// Define your rollback logic here
			return tx.Migrator().DropTable(&model.TrainClass{})
		},
	})
}
