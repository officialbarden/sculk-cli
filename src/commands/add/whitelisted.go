// checks integrity of library
package add

import ()

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
