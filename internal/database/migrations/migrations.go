package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
)

var Migrations []*gormigrate.Migration

func RegisterMigration(m *gormigrate.Migration) {
	Migrations = append(Migrations, m)
}

func GetMigrations() []*gormigrate.Migration {
	return Migrations
}
