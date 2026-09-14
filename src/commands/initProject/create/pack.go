package create

import (
	"encoding/json"
	"os"
	"path/filepath"
)



func CreatePackMcmeta(projectVersion string, projectType string) error {

	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// CREATE pack.mcmeta
	pathName := filepath.Join(workingDir + "/pack.mcmeta")
	file, err := os.Create(pathName)
	if err != nil {
		panic(err)
	}

	// WRITE to pack.mcmeta
	writePackMcmeta(file, projectVersion, projectType)

	err = file.Close()
	return err
	// fmt.Println(pathName, projectVersion)
}

func writePackMcmeta(file *os.File, projectVersion string, projectType string) {

	// Pack Format for Datapack
	// source: https://minecraft.wiki/w/Pack_format
	datapackPackFormatMap := map[string]int{
		"26.2": 107,
	}
	resourcepackPackFormatMap := map[string]int{
		"26.2": 88,
	}

	var packFormatNumber int
	var packDescription string

	// Different pack_format value based on resourcepack/datapack
	if projectType == "dp" {
		packFormatNumber = datapackPackFormatMap[projectVersion]
		packDescription = "a Sculk Project [datapack]"
	} else {
		packFormatNumber = resourcepackPackFormatMap[projectVersion]
		packDescription = "a Sculk Project [resourcepack]"
	}
	
	// Construct pack.mcmeta
	packFileContent := PackMcmetaFileContent{
		PackData: Pack{
			Pack_format: packFormatNumber,
			Description: packDescription,
		},
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(packFileContent)
	if err != nil {
		panic(err)
	}
}
