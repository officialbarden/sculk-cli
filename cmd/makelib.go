/*
Copyright © 2026 BARDEN <theofficialbarden@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"sculk/src/commands/makelib"
)

// makelibCmd represents the makelib command
var makelibCmd = &cobra.Command{
	Use:   "makelib",
	Short: "Add a libraries.json to your existing datapack/resourcepack",
	Long: `Adds a libraries.json file to your datapack/resourcepack so that others can download it using sculk-cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		makelib.Main(args)
	},
}

func init() {
	rootCmd.AddCommand(makelibCmd)
}
