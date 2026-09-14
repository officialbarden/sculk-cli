package packmcmeta

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Defines pack.mcmeta format
type Pack struct {
	Pack_format int    `json:"pack_format"`
	Description string `json:"description"`
}

type FileContent struct {
	PackData Pack `json:"pack"`
}

func CreatePackMcmeta(projectVersion string) error {

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
	writePackMcmeta(file, projectVersion)

	err = file.Close()
	return err
	// fmt.Println(pathName, projectVersion)
}

func writePackMcmeta(file *os.File, projectVersion string) {

	// Pack Format for Datapack
	// source: https://minecraft.wiki/w/Pack_format
	packFormatMap := map[string]int{
		"26.2": 107,
	}

	packFileContent := FileContent{
		PackData: Pack{
			Pack_format: packFormatMap[projectVersion],
			Description: "A Sculk Project",
		},
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(packFileContent)
	if err != nil {
		panic(err)
	}
}
