/*
Copyright © 2026 BARDEN <theofficialbarden@gmail.com>
*/
package cmd

import (
	"sculk-cli/src/links"

	"github.com/spf13/cobra"
)


// linksCmd represents the links command
var linksCmd = &cobra.Command{
	Use:   "links [...providers]?",
	Short: "Get official links affiliated with sculk.",
	Long:  `Get official links affiliated with sculk.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Command Execution
		output := links.Main(args)
		return output
	},
}

func init() {
	rootCmd.AddCommand(linksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// linksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// linksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
