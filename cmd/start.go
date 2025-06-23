package cmd

import (
	"fmt"
	"log"

	"TrainTracking/pkg/server"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the application...")
		err := server.StartApp()
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
