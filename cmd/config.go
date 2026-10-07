/*
Copyright © 2026 BARDEN <theofficialbarden@gmail.com>
*/
package cmd

import (
	"sculk/src/commands/config"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config [param] [value]",
	Short: "Set config for sculk-cli.",
	Long: `Set config for sculk-cli.
params:

doMerge: true/false
author: { name: string }

	`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Main(args)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
