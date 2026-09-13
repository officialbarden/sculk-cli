/*
Copyright © 2026 barden <theofficialbarden@gmail.com>
*/
package cmd

import (
	"fmt"
	"sculk-cli/src/add"

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

		// libraryName is a list of names so user can download multiple libraries.
		libraryName = args

		// Command Execution
		add.Main()

		fmt.Println("add called")
		return nil
	},
}

var libraryName []string
var ignoreVersioning bool = false

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVar(&ignoreVersioning, "ignore", ignoreVersioning, "Ignore Library Versioning — this will allow you to download libraries meant for other game versions.")
}
