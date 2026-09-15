// To Update
// Check if a newer version exists.
// If it exists, uninstall the current version, and install the new version
package update

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sculk-cli/src/commands/add"
	"sculk-cli/src/commands/initProject/create"

	"github.com/go-git/go-billy/v6"
)

func Main(args []string) {
	
	// check whether the user provided a specific identifier
	identifier := args
	
	if len(identifier) > 0 {
		for i := range identifier {
			_, fs, err := add.GetLibrarySource(identifier[i])
			if err != nil {
				return
			}
			ExecuteUpdate(identifier[i], fs)
		}
	}
}

// based on identifier
func ExecuteUpdate(libraryIdentifier string, fs billy.Filesystem) {

	if CheckVersionMismatch(libraryIdentifier, fs) {
		fmt.Printf("Library '%s' has been updated.", libraryIdentifier);
	} else {
		fmt.Printf("Library '%s' is up-2-date.", libraryIdentifier);
	}
}

func CheckVersionMismatch(libraryIdentifier string, fs billy.Filesystem) bool {
	
	var existingLibraryData create.Library
	var sourceLibraryData create.LibrariesDotJson

	sourceLibraryDataFile, err := fs.Open("/libraries.json")
	if err != nil {panic(err)}

	sourceLibraryDataJson, err := io.ReadAll(sourceLibraryDataFile)
	if err != nil {panic(err)}

	err = json.Unmarshal(sourceLibraryDataJson, &sourceLibraryData)
	if err != nil {panic(err)}

	// verify if update is needed.
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	project, err := os.ReadFile(filepath.Join(workDir, "/libraries.json"))
	if err != nil {
		panic(err)
	}

	var projectContent create.LibrariesDotJson
	err = json.Unmarshal(project, &projectContent)
	if err != nil { panic(err) }

	// find library with identifier.
	for i := range projectContent.Libraries {
		if projectContent.Libraries[i].Identifier == libraryIdentifier {
			existingLibraryData = projectContent.Libraries[i]
			break
		}
	}

	// check if there's version mismatch
	if sourceLibraryData.Version == existingLibraryData.LibraryVersion {
		return false
	} else { return true }
	
}