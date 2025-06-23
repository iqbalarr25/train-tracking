package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	RegisterMigration(&gormigrate.Migration{
		ID: "20250619141410_migration_create_stations_table",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.Exec(`
				CREATE TABLE IF NOT EXISTS stations (
					id BIGINT PRIMARY KEY,
					name TEXT NOT NULL,
					ref  TEXT,
					lat DOUBLE PRECISION NOT NULL,
					lon DOUBLE PRECISION NOT NULL,
					geom geometry(Point, 4326)
				);
			`).Error; err != nil {
				return err
			}

			if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_stations_ref ON stations(ref);`).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			if err := tx.Exec(`DROP INDEX IF EXISTS idx_stations_ref;`).Error; err != nil {
				return err
			}
			return tx.Exec(`DROP TABLE IF EXISTS stations;`).Error
		},
	})
}
