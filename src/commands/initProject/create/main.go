// THIS MODULE CREATES PACK.MCMETA
package create

import (
	"os/user"

	"charm.land/log/v2"
)

func CreateSculkProject(args []string, flags map[string]bool) error {

	projectName := args[0]
	projectVersion := args[1]
	log.Printf("🚧  Creating Sculk Project %s for Minecraft %s", projectName, projectVersion)

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
	log.Printf("🚧  Creating pack.mcmeta ...")
	err = CreatePackMcmeta(projectVersion, projectType)
	log.Printf("✅  Created pack.mcmeta")
	return err
}
