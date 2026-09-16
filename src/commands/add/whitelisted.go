// checks integrity of library
package add

import (
	"sculk-cli/src/commands/initProject/create"
)

type libraryBlock struct {
	Identifier string
	Source string
}

func VerifyLibraryIntegrity(identifier string) libraryBlock {

	// approved libraries: { "identifier": "github.com/identifier" }
	approvedLibraries := map[string]libraryBlock{
		"id-system": {
			Identifier:     "id-system",
			Source:         "http://github.com/officialbarden/id-system",
		},
		"uuid": {
			Identifier:     "uuid",
			Source:         "https://github.com/CJDevZ/UUID-Hex",
		},
	}
	return approvedLibraries[identifier]
}



func BuildImportedLibraryMetadata(libraryBlock libraryBlock) create.Library {
	var sourceLibraryDotJson create.LibrariesDotJson

	_, sourceLibraryDotJson, _, err := GetLibrarySource(libraryBlock.Identifier)	
	if err != nil {panic(err)}

	return create.Library{
		Identifier: libraryBlock.Identifier,
		Source: libraryBlock.Source,
		Version: sourceLibraryDotJson.Version,
		GameVersion: sourceLibraryDotJson.GameVersion,
	}
}