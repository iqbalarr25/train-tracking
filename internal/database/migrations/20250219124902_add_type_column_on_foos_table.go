package database

// Example of adding new column on exisiting table
// More Information: https://github.com/go-gormigrate/gormigrate

// import (
// 	"github.com/go-gormigrate/gormigrate/v2"
// 	"gorm.io/gorm"
// )

// func init() {
// 	RegisterMigration(&gormigrate.Migration{
// 		ID: "20250219124902_add_type_column_on_foos_table.go",
// 		Migrate: func(tx *gorm.DB) error {
// 			type Foo struct {
// 				gorm.Model
// 				Type string `gorm:"type:varchar(5)"`
// 			}
// 			return tx.Migrator().AddColumn(&Foo{}, "Type")
// 		},
// 		Rollback: func(tx *gorm.DB) error {
// 			type Foo struct {
// 				gorm.Model
// 				Type string `gorm:"type:varchar(5)"`
// 			}
// 			return tx.Migrator().DropColumn(&Foo{}, "Type")
// 		},
// 	})
// }
