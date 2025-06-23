package cmd_db

import (
	"TrainTracking/internal/config"
	seed "TrainTracking/internal/database/seeds"
	"TrainTracking/pkg/logger"

	"github.com/spf13/cobra"
)

var seedName string // Flag for specific seeder

var SeedCMD = &cobra.Command{
	Use:   "seed",
	Short: "Run database seeders",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("🌱 Running database seeders...")

		config.InitDatabase()
		db := config.GetDBConnection()
		if db == nil {
			logger.Fatal("❌ Database connection failed, DB is nil")
		}

		// Run a specific seeder if --name is provided
		if seedName != "" {
			for _, seeder := range seed.GetSeeders() {
				if seeder.Name == seedName {
					logger.Info("🌱 Running seeder: %s", seeder.Name)
					if err := seeder.Run(db); err != nil {
						logger.Log.WithError(err).Fatalf("❌ Seeder %s failed", seeder.Name)
					}
					logger.Info("✅ Seeder %s completed", seeder.Name)
					return
				}
			}
			logger.Fatal("❌ Seeder with name '%s' not found", seedName)
		}

		// Run all seeders if no specific seeder is provided
		for _, seeder := range seed.GetSeeders() {
			logger.Info("🌱 Running seeder: %s", seeder.Name)
			if err := seeder.Run(db); err != nil {
				logger.Log.WithError(err).Fatalf("❌ Seeder %s failed", seeder.Name)
			}
		}

		logger.Info("✅ Database seeding completed.")
	},
}

func init() {
	SeedCMD.Flags().StringVarP(&seedName, "name", "n", "", "Run a specific seeder by name")
}
