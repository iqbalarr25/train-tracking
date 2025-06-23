package cmd_db

import (
	"fmt"

	"TrainTracking/internal/config"
	database "TrainTracking/internal/database/migrations"
	"TrainTracking/pkg/logger"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
)

var RollbackCMD = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback last migration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Rolling back last migration...")

		config.InitDatabase()
		m := gormigrate.New(config.GetDBConnection(), gormigrate.DefaultOptions, database.GetMigrations())

		if err := m.RollbackLast(); err != nil {
			logger.Error("Rollback failed: %v", err)
		}

		logger.Info("✅ Last migration rolled back sucessfully")
	},
}

func init() {}
