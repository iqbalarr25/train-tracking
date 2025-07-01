package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	RegisterMigration(&gormigrate.Migration{
		ID: "20250618142332_migration_create_rail_edges_table",
		Migrate: func(tx *gorm.DB) error {
			// Define your migration logic here
			return tx.Exec(`
				CREATE EXTENSION IF NOT EXISTS "pgcrypto";

				CREATE TABLE IF NOT EXISTS rail_edges (
					id UUID PRIMARY KEY,
					source UUID NOT NULL,
					target UUID NOT NULL,
					cost DOUBLE PRECISION NOT NULL,
					reverse_cost DOUBLE PRECISION,
					max_speed DOUBLE PRECISION,
					geom geometry(LineString, 4326)
				)
			`).Error
		},
		Rollback: func(tx *gorm.DB) error {
			// Define your rollback logic here
			return tx.Exec(`DROP TABLE IF EXISTS rail_edges`).Error
		},
	})
}
