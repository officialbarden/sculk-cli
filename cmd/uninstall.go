/*
Copyright © 2026 barden <theofficialbarden@gmail.com>
*/
package cmd

import (
	"sculk-cli/src/commands/uninstall"

	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall [...libraryName]?",
	Short: "Uninstall specific/all libraries specified in libraries.json",
	Long:  `Uninstall specific/all libraries specified in libraries.json`,
	Run: func(cmd *cobra.Command, args []string) {
		// Command Execution
		uninstall.Main(args)
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
