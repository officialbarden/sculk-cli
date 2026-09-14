// THIS MODULE CREATES PACK.MCMETA
package create

func CreateSculkProject(args []string, flags map[string]bool) error {
	projectVersion := args[1]

	projectType := ""

	if flags["dp"] == true {
		projectType = "dp"
	} else {
		projectType = "rp"
	}

	// create pack.mcmeta file
	err := CreatePackMcmeta(projectVersion, projectType)
	return err
}
