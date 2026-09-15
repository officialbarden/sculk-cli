package add

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-billy/v6"
	"github.com/go-git/go-billy/v6/memfs"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/storage/memory"
)

func InstallLibraries(libraries []string) {

	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for i := range libraries {
		// use in-built hashmap to get src via acronym/identifiers
		repoURL := VerifyLibraryIntegrity(libraries[i]).Source
		
		// clone repo on ram
		fs := memfs.New()
		_, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
			URL: repoURL,
			Depth: 1,
		})
		if err != nil {
			panic(err)
		}

		err = MergeIndividualFiles(fs, "/", workingDir)
	}
}

func MergeIndividualFiles(fs billy.Filesystem, currentPath string, targetDir string) error {
	files, err := fs.ReadDir(currentPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {

		// skip LICENSE file
		// if file.Name() == "LICENSE" {
		// 	continue
		// }
		
		// store memory path and local path
		memoryPath := filepath.Join(currentPath, file.Name())
		localPath := filepath.Join(targetDir, file.Name())

		if file.IsDir() {
			// if found directory, go inside that directory and recurse. 
			err := MergeIndividualFiles(fs, memoryPath, localPath)
			if err != nil {
				panic(err)
			}
		} else {
			err := handleFileMerging(fs, memoryPath, localPath)
			if err != nil {
				panic(err)
			}
		}
		
	}

	return nil
}

type FunctionTag struct {
	Replace bool `json:"replace"`
	Values []string `json:"values"`
}

func handleFileMerging(fs billy.Filesystem, sourcePath string, destinationPath string) error {
	srcFile, err := fs.Open(sourcePath)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	err = os.MkdirAll(filepath.Dir(destinationPath), 0755)
	if err != nil {
		panic(err)
	}


	// if file already exists, append contents to the top of the file.
	if fileInfo, err := os.Stat(destinationPath); err == nil {

		fileName := strings.Split(fileInfo.Name(), ".");
		fileExtension := ""
		for i := range fileName {

			if i == len(fileName) - 1 {
				fileExtension = fileName[i];
			}

		}
		if fileExtension == "json" {

			if fileInfo.Name() == "load.json" || fileInfo.Name() == "tick.json" {

				var sourceTagContent FunctionTag
				var existingTagContent FunctionTag
				
				existingContent, err := os.ReadFile(destinationPath)
				if err != nil {
					panic(err)
				}
				
				sourceContent, err := io.ReadAll(srcFile)
				if err != nil {
					panic(err)
				}
				
				err = json.Unmarshal(sourceContent, &sourceTagContent)
				if err != nil {
					panic(err)
				}

				err = json.Unmarshal(existingContent, &existingTagContent)
				if err != nil {
					panic(err)
				}

				existingTagContent.Values = append(existingTagContent.Values, sourceTagContent.Values...)
				combined, err := json.Marshal(existingTagContent)
				if err != nil {
					panic(err)
				}
				return os.WriteFile(destinationPath, combined, 0644)
				
			} else {
				return nil
			}
			// return nil	// to break out of the overarching loop
		}
		
		existingData, err := os.ReadFile(destinationPath);
		if err != nil {
			fmt.Println("here is panic")
			panic(err)
		}

		newData, err := io.ReadAll(srcFile);
		if err != nil {
			panic(err)
		}

		combined := append(newData, []byte("\n# added by sculk ^^\n")...)
		combined = append(combined, existingData...)

		return os.WriteFile(destinationPath, combined, 0644);
	}

	// if file doesn't exist:
	destinationFile, err := os.OpenFile(destinationPath, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, srcFile);
	return err
	
	
}
