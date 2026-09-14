// THIS MODULE CREATES PACK.MCMETA
package create

import (
	"os/user"
)

func CreateSculkProject(args []string, flags map[string]bool) error {

	projectName := args[0]
	projectVersion := args[1]

	projectType := "dp" // default is datapack project

	if flags["dp"] == true {
		projectType = "dp"
	} else if flags["rp"] == true {
		projectType = "rp"
	}

	// create libraries.json
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
		
	CreateLibrariesJson(user.Name, projectVersion)

	// actually create files now
	switch projectType {
	case "dp":
		InitDatapack(projectName, projectVersion)
	case "rp":
		InitResourcepack(projectName, projectVersion)
	}

	// create pack.mcmeta file
	err = CreatePackMcmeta(projectVersion, projectType)
	return err
}
