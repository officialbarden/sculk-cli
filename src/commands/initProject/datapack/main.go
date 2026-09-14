package datapack

import (
	"sculk-cli/src/commands/initProject/packmcmeta"
)

func Create(args []string) error {
	projectVersion := args[1]

	// create pack.mcmeta file
	err := packmcmeta.CreatePackMcmeta(projectVersion)
	return err
}
