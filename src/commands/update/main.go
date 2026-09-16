// To Update
// Check if a newer version exists.
// If it exists, uninstall the current version, and install the new version
package update

import (
	"sculk-cli/src/commands/add"
	"sculk-cli/src/commands/initProject/create"

	"charm.land/log/v2"
	"github.com/go-git/go-billy/v6"
)

func Main(args []string) {
	
	// check whether the user provided a specific identifier
	identifier := args
	
	if len(identifier) > 0 {
		for i := range identifier {
			var libraryDotJson create.LibrariesDotJson
			
			// get source code for that library from git
			_, libraryDotJson, fs, err := add.GetLibrarySource(identifier[i])
			if err != nil {
				return
			}

			// put src-code's libraries.json, identifier and fs in execUpdate
			ExecuteUpdate(libraryDotJson, identifier[i], fs)
		}
	}
}

// based on identifier
func ExecuteUpdate(libraryDotJson create.LibrariesDotJson, libraryIdentifier string, fs billy.Filesystem) {

	isMismatch, oldVersion, newVersion := CheckVersionMismatch(libraryDotJson, libraryIdentifier, fs)	
	if isMismatch {

		// update library and update in libraries.json:
		
		// log
		log.Printf("Library '%s' has been updated [%s -> %s].", libraryIdentifier, oldVersion, newVersion);
	} else {
		log.Printf("Library '%s' is up-to-date (v. %s).", libraryIdentifier, oldVersion);
	}
}

func CheckVersionMismatch(libraryDotJson create.LibrariesDotJson, libraryIdentifier string, fs billy.Filesystem) (isMismatch bool, oldVersion string, newVersion string) {
	
	existingLibraryData := add.ReadLocalLibrariesJson()
	sourceLibraryData := libraryDotJson
	var installedLibraryMetaData create.Library
	
	// find library with identifier.
	for i := range existingLibraryData.Libraries {
		if existingLibraryData.Libraries[i].Identifier == libraryIdentifier {
			// once the identifier metadata is found in local libraries.json, store it in variable to compare later.
			installedLibraryMetaData = existingLibraryData.Libraries[i]
			break
		}
	}

	// check if there's version mismatch
	if sourceLibraryData.Version == installedLibraryMetaData.Version {
		return false, installedLibraryMetaData.Version, sourceLibraryData.Version
	} else { return true, installedLibraryMetaData.Version, sourceLibraryData.Version }
}