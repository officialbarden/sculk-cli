// Create the libraries.json file which stores sculk project data
package create

import (
	"os"
	"path/filepath"
	"encoding/json"
)

func CreateLibrariesJson() {
	
	// get target path
	targetPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// create libraries.json
	path := filepath.Join(targetPath + "/libraries.json")
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	
	// close file after program exits
	defer file.Close()

	// file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "   ")

} 

func libraryDataCreate() {}
