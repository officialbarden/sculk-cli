/*
Copyright © 2026 barden <theofficialbarden@gmail.com>
*/
package cmd

import (
	"sculk-cli/src/commands/add"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [...libraryName]",
	Short: "Add a library (or multiple libraries) to your project.",
	Args:  cobra.MinimumNArgs(1),
	Long: `
Add a library to your project. See more here: officialbarden.github.io/sculk/libraries
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Command Execution
		output := add.Main(ignoreVersionMismatch, args)
		return output
	},
}

var libraryName []string
var ignoreVersionMismatch bool = false

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVar(&ignoreVersionMismatch, "ignore", ignoreVersionMismatch, "Ignore Library Versioning — this will allow you to download libraries meant for other game versions.")
}
