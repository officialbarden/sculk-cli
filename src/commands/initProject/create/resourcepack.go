// this .go file handles initializing resourcepack files.
package create

import ()

func InitResourcepack(projectName string, projectVersion string) {

	// mapped as '26.2': []files
	fileStructure := map[string][]ResourcepackFile{
		"26.2": {
			{"/assets/minecraft/tags/function/", "load.json", []byte("")},
			{"/assets/minecraft/tags/function/", "tick.json", []byte("")},
			{AddProjectNamespace("/assets/", projectName, "/function/global/"), "load.mcfunction", []byte("")},
			{AddProjectNamespace("/assets/", projectName, "/function/global/"), "tick.mcfunction", []byte("")},
		},
	}

	
	// MAKE THE FILES
	CreateResourcepackFiles(fileStructure[projectVersion])

}
