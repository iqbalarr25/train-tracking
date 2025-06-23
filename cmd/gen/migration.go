package gen_cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"TrainTracking/pkg/logger"

	"github.com/spf13/cobra"
)

var MigrationCMD = &cobra.Command{
	Use:   "migration",
	Short: "Generate migration file",
	Run:   generateMigrationFile,
}

func init() {
}

func generateMigrationFile(cmd *cobra.Command, args []string) {
	migrationName := args[0]
	timestamp := time.Now().Format("20060102150405")
	filename := fmt.Sprintf("%s_migration_%s.go", timestamp, migrationName)

	migrationsDir := "internal/database/migrations"
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		err := os.MkdirAll(migrationsDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create migrations directory: %v", err)
		}
	}

	filePath := filepath.Join(migrationsDir, filename)
	fileContent := fmt.Sprintf(`package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	RegisterMigration(&gormigrate.Migration{
		ID: "%s_migration_%s",
		Migrate: func(tx *gorm.DB) error {
			// Define your migration logic here
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Define your rollback logic here
			return nil
		},
	})
}`, timestamp, migrationName)

	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		logger.Fatal("Failed to create migration file: %v", err)
	}

	logger.Info(fmt.Sprintf("✅ Database migration successfully created at %s.", filePath))
}
