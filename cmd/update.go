/*
Copyright © 2026 barden <theofficialbarden@gmail.com>
*/
package cmd

import (
	"sculk-cli/src/commands/update"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [...libraryName]?",
	Short: "Update specific/all libraries to their latest subversion.",
	Long: `Update specific/all libraries to their latest subversion for project's specified game version.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Command Execution
		update.Main(args)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
