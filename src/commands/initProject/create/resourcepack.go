// this .go file handles initializing resourcepack files.
package create

import ()

func InitResourcepack(projectName string, projectVersion string) {
	// mapped as '26.2': []files
	fileStructure := GetResourcepackFilesList(projectName, projectVersion)
	
	// MAKE THE FILES
	CreateResourcepackFiles(fileStructure)
}
