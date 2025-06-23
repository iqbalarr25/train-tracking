package database

import (
	"TrainTracking/internal/features/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	RegisterMigration(&gormigrate.Migration{
		ID: "20250305073317_migration_create_users_table",
		Migrate: func(tx *gorm.DB) error {
			// Define your migration logic here
			return tx.AutoMigrate(&model.User{})
		},
		Rollback: func(tx *gorm.DB) error {
			// Define your rollback logic here
			return tx.Migrator().DropTable(&model.User{})
		},
	})
}
