package initProject

import (
	"sculk-cli/src/commands/initProject/create"
)

// sculk init [projectName] [projectVersion]
func Main(args []string, flags map[string]bool) error {
	output := create.CreateSculkProject(args, flags)
	return output
}
