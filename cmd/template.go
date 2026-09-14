/*
Copyright © 2026 BARDEN <theofficialbarden@gmail.com>

*/
package cmd

import (
	"sculk-cli/src/commands/template"
	"github.com/spf13/cobra"
)

// sculk template add/use [templateName]
var templateCmd = &cobra.Command{
	Use:   "template",
	Args: cobra.ExactArgs(2),
	Short: "Create and use existing project as a base/template for future projects.",
	Long: `Create and use existing project as a base/template for future projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		template.Main(args)
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
