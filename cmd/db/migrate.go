package cmd_db

import (
	"TrainTracking/internal/config"
	database "TrainTracking/internal/database/migrations"
	"fmt"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
)

var MigrateCMD = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting database migration...")

		config.InitDatabase()
		m := gormigrate.New(config.GetDBConnection(), gormigrate.DefaultOptions, database.GetMigrations())

		if err := m.Migrate(); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}

		fmt.Println("Database migrations completed.")
	},
}

func init() {}
