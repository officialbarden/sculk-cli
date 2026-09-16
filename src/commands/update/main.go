// To update
// Check if a newer version exists.
// If it exists, uninstall the current version, and install the new version
package update

import (
	"sculk-cli/src/commands/add"
	"sculk-cli/src/commands/initProject/create"
	"sculk-cli/src/commands/uninstall"

	"charm.land/log/v2"
	"github.com/go-git/go-billy/v6"
)

func Main(args []string) {
	
	// check whether the user provided a specific identifier
	identifier := args
	
	if len(identifier) > 0 {
		for _, identifierString := range identifier {
			updateInit(identifierString)
		}
	} else {

		// get all identifiers from local /libraries.json:
		libraryFile := add.ReadLocalLibrariesJson()
		var libraryIdentifiers []string
		for _, file := range libraryFile.Libraries {
			libraryIdentifiers = append(libraryIdentifiers, file.Identifier)
		}

		// updateInit each of em:
		for _, identifierString := range libraryIdentifiers {
			updateInit(identifierString)
		}
	}
}

// based on identifier
func ExecuteupdateInit(libraryDotJson create.LibrariesDotJson, libraryIdentifier string, fs billy.Filesystem) {

	isMismatch, oldVersion, newVersion := CheckVersionMismatch(libraryDotJson, libraryIdentifier, fs)	
	if isMismatch {
		log.Printf("🍁 Library '%s' is outdated. Updating [%s -> %s] ...", libraryIdentifier, oldVersion, newVersion);

		// updateInit library and updateInit in libraries.json:
		// uninstall, install
		uninstall.UninstallLibrary(libraryIdentifier)
		add.InstallLibraries(libraryIdentifier, false)
		
		// log
		log.Printf("🍀 Library '%s' has been updateInitd [%s -> %s].", libraryIdentifier, oldVersion, newVersion);
	} else {
		log.Printf("🍀 Library '%s' is up-to-date (v. %s).", libraryIdentifier, oldVersion);
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

func updateInit(libraryIdentifier string) {
	var libraryDotJson create.LibrariesDotJson
	
	// get source code for that library from git
	_, libraryDotJson, fs, err := add.GetLibrarySource(libraryIdentifier)
	if err != nil {
		return
	}

	// put src-code's libraries.json, identifier and fs in execupdateInit
	ExecuteupdateInit(libraryDotJson, libraryIdentifier, fs)
}