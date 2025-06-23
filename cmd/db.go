/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package cmd

import (
	cmd_db "TrainTracking/cmd/db"

	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db <option>",
	Short: "Start db related operations",
	Long: `Start a database operation.
Please key in an option to start. Type 'db -h' for more information.

Popular options are:
- db migrate
- db rollback`,
}

func init() {
	dbCmd.AddCommand(cmd_db.MigrateCMD)
	dbCmd.AddCommand(cmd_db.RollbackCMD)
	dbCmd.AddCommand(cmd_db.SeedCMD)

	rootCmd.AddCommand(dbCmd)
}
