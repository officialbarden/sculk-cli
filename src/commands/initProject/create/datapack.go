// this .go file handles initializing datapack files.
package create

import ()

func InitDatapack(projectName string, projectVersion string) {
	// mapped as '26.2': []files
	fileStructure := GetDatapackFilesList(projectName, projectVersion)

	// MAKE THE FILES
	CreateDatapackFiles(fileStructure)
}