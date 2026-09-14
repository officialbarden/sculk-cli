/*
Copyright © 2026 BARDEN <theofficialbarden@gmail.com>
*/
package cmd

import (
	"sculk-cli/src/commands/remove"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [...libraryName]",
	Short: "Remove (atleast) one library from the project.",
	Args:  cobra.MinimumNArgs(1),
	Long: `The remove command requires users to input specific libraries they want removed from their project,
which is essential in-order to not accidently uninstall all libraries. Hence, this command works as a safe alternative to 'uninstall'.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Command Execution
		remove.Main()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
