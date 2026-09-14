// this .go file handles initializing datapack files.
package create

import ()


func InitDatapack(projectName string, projectVersion string) {
	// mapped as '26.2': []files
	fileStructure := map[string][]DatapackFile{
		"26.2": {
			{"/data/minecraft/tags/function/", "load.json", []byte("")},
			{"/data/minecraft/tags/function/", "tick.json", []byte("")},
			{AddProjectNamespace("/data/", projectName, "/function/global/"), "load.mcfunction", []byte("")},
			{AddProjectNamespace("/data/", projectName, "/function/global/"), "tick.mcfunction", []byte("")},
		},
	}

	// MAKE THE FILES
	CreateDatapackFiles(fileStructure[projectVersion])
	
}
