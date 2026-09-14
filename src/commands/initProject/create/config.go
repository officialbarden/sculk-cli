// important data
package create

import (
	"os"
	"path/filepath"
)

// Defines pack.mcmeta format
type PackMcmetaFileContent struct {
	PackData Pack `json:"pack"`
}

type Pack struct {
	Pack_format int    `json:"pack_format"`
	Description string `json:"description"`
}

type DatapackFile struct {
	DirName string
	FileName string
	Content []byte
}


type ResourcepackFile struct {
	DirName  string
	FileName string
	Content  []byte
}



func AddProjectNamespace(filePathPrefix string, namespace string, filePathSuffix string) string {
	return filePathPrefix + namespace + filePathSuffix
}

func CreateDatapackFiles(files []DatapackFile) {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// make them.
	for i := range files {
		err := os.MkdirAll(filepath.Join(workingDir, files[i].DirName), 0755)
		if err != nil {
			panic(err)
		}
		file, err := os.Create(filepath.Join(workingDir, files[i].DirName, files[i].FileName))
		if err != nil {
			panic(err)
		}
		file.Close()
	}
}

func CreateResourcepackFiles(files []ResourcepackFile) {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// make them.
	for i := range files {
		err := os.MkdirAll(filepath.Join(workingDir, files[i].DirName), 0755)
		if err != nil {
			panic(err)
		}
		file, err := os.Create(filepath.Join(workingDir, files[i].DirName, files[i].FileName))
		if err != nil {
			panic(err)
		}
		file.Close()
	}
}