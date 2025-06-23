package cmd

import (
	"TrainTracking/internal/config"
	"TrainTracking/pkg/logger"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   config.GetApp().Name,
	Short: "A CLI tool for managing your application",
	Long:  `Airplane Tracking`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	logger.InitLogger()
}
