package create

import (
	"sculk-cli/src/commands/initProject/packmcmeta"
)

func CreateSculkProject(args []string, flags map[string]bool) error {
	projectVersion := args[1]

	projectType := ""

	if flags["dp"] == true {
		projectType = "dp"
	} else {
		projectType = "rp"
	}

	// create pack.mcmeta file
	err := packmcmeta.CreatePackMcmeta(projectVersion, projectType)
	return err
}
