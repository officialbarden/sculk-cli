/*
Copyright © 2026 barden <theofficialbarden@gmail.com>
*/
package cmd

import (
	"sculk-cli/src/commands/initProject"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [projectName] [projectVersion]",
	Short: "Initialise a new project — datapack or resourcepack.",
	Long: `
Initialise a new project, add libraries from github, write standard necessary components like id-systems,
	`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {

		flags := map[string]bool{
			"dp": datapackProject,
			"rp": resourcepackProject,
		}

		// Command Execution
		output := initProject.Main(args, flags)

		return output
	},
}

// Flags (and their default values.)
var resourcepackProject bool = false
var datapackProject bool = false

func init() {
	rootCmd.AddCommand(initCmd)
	// either '--dp' or '--rp'
	initCmd.Flags().BoolVar(&resourcepackProject, "rp", resourcepackProject, "Initializes a 'Resourcepack' project.")
	initCmd.Flags().BoolVar(&datapackProject, "dp", datapackProject, "Initializes a 'Datapack' project.")
	initCmd.MarkFlagsMutuallyExclusive("dp", "rp")
}
