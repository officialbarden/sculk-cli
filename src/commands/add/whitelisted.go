// checks integrity of library
package add

import (
	"sculk-cli/src/commands/initProject/create"
)

func VerifyLibraryIntegrity(identifier string) create.Library {
	// approved libraries: { "identifier": "github.com/identifier" }
	approvedLibraries := map[string]create.Library {
		"id-system": {
			Identifier: "id-system",
			Name: "ID System",
			Source: "http://github.com/officialbarden/id-system",
			LibraryVersion: "1.0.0",
			GameVersion: "26.2",
		},
		"uuid": {
			Identifier: "uuid",
			Name: "UUID",
			Source: "https://github.com/CJDevZ/UUID-Hex",
			LibraryVersion: "1.0.0",
			GameVersion: "26.2",
		},
	}
	return approvedLibraries[identifier]
}