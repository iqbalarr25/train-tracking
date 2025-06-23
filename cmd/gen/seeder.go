package gen_cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"TrainTracking/pkg/logger"

	"github.com/spf13/cobra"
)

var SeederCMD = &cobra.Command{
	Use:   "seeder",
	Short: "Generate seeder file",
	Run:   generateSeederFile,
}

func init() {
}

func generateSeederFile(cmd *cobra.Command, args []string) {
	seedName := args[0]
	filename := fmt.Sprintf("%sSeed.go", seedName)

	seedDir := "internal/database/seeds"
	if _, err := os.Stat(seedDir); os.IsNotExist(err) {
		err := os.MkdirAll(seedDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create migrations directory: %v", err)
		}
	}

	filePath := filepath.Join(seedDir, filename)
	fileContent := fmt.Sprintf(`package database

import (
	"log"

	"gorm.io/gorm"
)

func init() {
	RegisterSeeder(Seeder{
		Name: "%sSeed",
		Run: func(tx *gorm.DB) error {
			items := []struct {
				Name string
			}{
				{"John"},
				{"Smith"},
			}

			// You can use tx.Model(&items) or bellow:
			if err := tx.Table("your_table").Create(&items).Error; err != nil {
				return err
			}

			log.Println("✅ %s seeded successfully")
			return nil
		},
	})
}`, seedName, seedName)

	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		logger.Fatal("Failed to create migration file: %v", err)
	}

	logger.Info(fmt.Sprintf("✅ Database seed successfully created at %s.", filePath))
}
