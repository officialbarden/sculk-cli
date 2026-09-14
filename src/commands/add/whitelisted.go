// checks integrity of library
package add

func VerifyLibraryIntegrity(identifier string) string {
	// approved libraries: { "identifier": "github.com/identifier" }
	approvedLibraries := map[string]string {
		"id-system": "http://github.com/officialbarden/id-system",
		"uuid": "https://github.com/CJDevZ/UUID-Hex",
	}
	return approvedLibraries[identifier]
}