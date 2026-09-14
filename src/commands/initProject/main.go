package initProject

import "sculk-cli/src/commands/initProject/datapack"

// sculk init [projectName] [projectVersion]
func Main(args []string) error {
	output := datapack.Create(args)
	return output
}
