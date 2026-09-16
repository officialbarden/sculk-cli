package add

import (
	"encoding/json"
	"os"
	"path/filepath"

	"charm.land/log/v2"
)

func AddToLibrariesJson(libraryIdentifier string) error {
	log.Printf("🚧 Adding '%s' library to libraries.json ...", libraryIdentifier)
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	jsonPath := filepath.Join(workingDir, "/libraries.json")
	file, err := os.Open(jsonPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	librariesJson := ReadLocalLibrariesJson()

	for _, lib := range librariesJson.Libraries {
		if lib.Identifier == libraryIdentifier {
			log.Printf("⚠ Library '%s' already exists in libraries.json", libraryIdentifier)
			return nil
		}
	}
	
	librariesJson.Libraries = append(librariesJson.Libraries, BuildImportedLibraryMetadata(VerifyLibraryIntegrity(libraryIdentifier)))
	combined, err := json.MarshalIndent(librariesJson, "", "   ")

	log.Printf("⚙ Added to libraries.json")
	return os.WriteFile(jsonPath, combined, 0644)
}

func RemoveFromLibrariesJson(libraryIdentifier string) error {
	log.Printf("🚧 Removing '%s' library from libraries.json ...", libraryIdentifier)
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	jsonPath := filepath.Join(workingDir, "/libraries.json")
	file, err := os.Open(jsonPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	librariesJson := ReadLocalLibrariesJson()

	found := false
	filtered := librariesJson.Libraries[:0]
	for _, lib := range librariesJson.Libraries {
		if lib.Identifier == libraryIdentifier {
			found = true
			continue
		}
		filtered = append(filtered, lib)
	}

	if !found {
		log.Printf("⚠ Library '%s' not found in libraries.json", libraryIdentifier)
		return nil
	}

	librariesJson.Libraries = filtered

	combined, err := json.MarshalIndent(librariesJson, "", "   ")
	if err != nil {
		panic(err)
	}

	log.Printf("🗑 Removed from libraries.json")
	return os.WriteFile(jsonPath, combined, 0644)
}