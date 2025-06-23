package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	RegisterMigration(&gormigrate.Migration{
		ID: "20250618114643_migration_create_rail_nodes_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.Exec(`
				CREATE TABLE IF NOT EXISTS rail_nodes (
					id BIGINT PRIMARY KEY,
					lat DOUBLE PRECISION NOT NULL,
					lon DOUBLE PRECISION NOT NULL,
					geom geometry(Point, 4326)
				)
			`).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Exec(`DROP TABLE IF EXISTS rail_nodes`).Error
		},
	})
}
