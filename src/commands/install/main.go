package install

import (
	"encoding/json"
	"os"
	"path"
	"sculk-cli/src/commands/initProject/create"
	"sculk-cli/src/commands/add"
)

func Main() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// find libraries.json
	fullPath := path.Join(workingDir, "/libraries.json")
	// read bytes
	file, err := os.ReadFile(fullPath)

	// parse and store json
	var fileJson create.LibrariesDotJson
	err = json.Unmarshal(file, &fileJson)

	librariesOriginal := fileJson.Libraries
	var libraryIdentifiers []string
	for i := range librariesOriginal {
		libraryIdentifiers = append(libraryIdentifiers, librariesOriginal[i].Identifier)
	}
	
	add.InstallLibraries(libraryIdentifiers)
}
