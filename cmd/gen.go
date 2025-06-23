/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	gen_cmd "TrainTracking/cmd/gen"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate <option>",
	Short: "Generate your migration, seeder or features scaffold",
	Long:  `Generate your migration, seeder or features scaffolding.`,
	Args:  cobra.ExactArgs(1),
}

func init() {
	generateCmd.AddCommand(gen_cmd.MigrationCMD)
	generateCmd.AddCommand(gen_cmd.SeederCMD)
	generateCmd.AddCommand(gen_cmd.GenerateFeatureCMD)
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
