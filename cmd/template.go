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
	Args:  cobra.ExactArgs(1),
	Short: "Create and use existing project as a base/template for future projects.",
	Long:  `Create and use existing project as a base/template for future projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		if createFlag {
			template.Main(args, "create")
		} else if deleteFlag {
			template.Main(args, "delete")
		} else if addFlag {
			template.Main(args, "add")
		}
	},
}

var createFlag bool = false
var deleteFlag bool = false
var addFlag bool = false

func init() {
	rootCmd.AddCommand(templateCmd)

	templateCmd.Flags().BoolVarP(&createFlag, "create", "c", createFlag, "Save the current file-structure as a template.")
	templateCmd.Flags().BoolVarP(&deleteFlag, "delete", "d", deleteFlag, "Delete the named template.")
	templateCmd.Flags().BoolVarP(&addFlag, "add", "a", addFlag, "Add a named template to the current sculk-project.")
}
