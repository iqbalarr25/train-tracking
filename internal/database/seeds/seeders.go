package database

import (
	"gorm.io/gorm"
)

type Seeder struct {
	Name string
	Run  func(*gorm.DB) error
}

var Seeders []Seeder

func RegisterSeeder(s Seeder) {
	Seeders = append(Seeders, s)
}

func GetSeeders() []Seeder {
	return Seeders
}
