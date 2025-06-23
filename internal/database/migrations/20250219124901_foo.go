package database

// Example of creating a new table
// More Information: https://github.com/go-gormigrate/gormigrate

// import (
// 	"github.com/go-gormigrate/gormigrate/v2"
// 	"gorm.io/gorm"
// )

// func init() {
// 	RegisterMigration(&gormigrate.Migration{
// 		ID: "20250219124901_foo",
// 		Migrate: func(tx *gorm.DB) error {
// 			type Foo struct {
// 				gorm.Model
// 				Name string `gorm:"type:varchar(10)"`
// 			}
// 			return tx.AutoMigrate(&Foo{})
// 		},
// 		Rollback: func(tx *gorm.DB) error {
// 			return tx.Migrator().DropTable("foos")
// 		},
// 	})
// }
