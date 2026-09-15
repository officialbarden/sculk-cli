package add

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"sculk-cli/src/commands/initProject/create"
)

func AddToLibrariesJson(libraries []string) error {
	
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	jsonPath := filepath.Join(workingDir, "libraries.json")
	file, err := os.Open(jsonPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	fileContent, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var librariesJson create.LibrariesDotJson

	err = json.Unmarshal(fileContent, &librariesJson)
	if err != nil {
		panic(err)
	}

	for i := range libraries {
		librariesJson.Libraries = append(librariesJson.Libraries, VerifyLibraryIntegrity(libraries[i]))
	}
	combined, err := json.MarshalIndent(librariesJson, "", "   ")

	return os.WriteFile(jsonPath, combined, 0644)
}

