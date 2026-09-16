package add

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"sculk-cli/src/commands/initProject/create"

	"charm.land/log/v2"
)

func AddToLibrariesJson(libraryIdentifier string) error {
	log.Printf("🚧 Adding %s to libraries.json ...", libraryIdentifier)
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

	for i := range librariesJson.Libraries {

		// do not create duplicate library entries
		if librariesJson.Libraries[i].Identifier == libraryIdentifier {
			continue
		}

		librariesJson.Libraries = append(librariesJson.Libraries, VerifyLibraryIntegrity(libraries[i]))
	}
	combined, err := json.MarshalIndent(librariesJson, "", "   ")

	log.Printf("⚙ Added to libraries.json")
	return os.WriteFile(jsonPath, combined, 0644)
}
