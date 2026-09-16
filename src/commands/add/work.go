package add

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sculk-cli/src/commands/initProject/create"
	"slices"
	"strings"

	"charm.land/log/v2"

	"github.com/go-git/go-billy/v6"
	"github.com/go-git/go-billy/v6/memfs"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/storage/memory"
)

func InstallLibraries(libraries []string, ignoreVersionMismatch bool) {

	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for i := range libraries {

		libraryRepoLink, _, _, fs, err := GetLibrarySource(libraries[i])
		if err != nil {
			panic(err)
		}

		// LOGGER: found library
		log.Printf("📩 Downloading %s", libraryRepoLink)

		// Match Game Versions, if flag doesn't exist.
		isIncompatWithGame, _, _ := gameVersionIncompat(fs)
		if !ignoreVersionMismatch && isIncompatWithGame {
			// LOGGER: Library is Incompatible (and no --ignore flag)
			log.Fatal("The library you're is incompatible with your project's defined game version. Either change the project's version in libraries.json, or use the --ignore flag to install anyway.")
		}

		err = MergeIndividualFiles(fs, "/", workingDir)
		if err != nil {
			panic(err)
		}

		// merge this library to libraries.json
		err = AddToLibrariesJson(libraries)
		if err != nil { panic(err) }
	}
}

func IsPreinstalled(libraryIdentifier string) bool {

	var librariesInstalled []create.Library

	log.Printf("🔍 Checking in libraries.json ...")
	librariesInstalled = ReadLocalLibrariesJson().Libraries	

	for i := range librariesInstalled {
		if librariesInstalled[i].Identifier == libraryIdentifier {
			log.Printf("🔍 Library is already installed.")
			return true
		}
	}

	return false
}


// READS WORKING DIRECTORY'S libraries.json FILE.
func ReadLocalLibrariesJson() create.LibrariesDotJson {
	var libraryJson create.LibrariesDotJson
	// build current path
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := filepath.Join(workDir, "/libraries.json")
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fileData, err := io.ReadAll(file)
	err = json.Unmarshal(fileData, &libraryJson)
	return libraryJson
}

func GetLibrarySource(libraryIdentifier string) (repo string, libraryDotJson create.LibrariesDotJson, fs billy.Filesystem, error error) {
	// GETS THE IDENTIFIER STRING, LIBRARY VERSION AND GAME VERSION
	// AND RETURNS INFORMATION AND RAM-DOWNLOADED fs

	// use in-built hashmap to get src via acronym/identifiers
	log.Printf("🔍 Finding library in internal manager ...")
	repoURL := VerifyLibraryIntegrity(libraryIdentifier).Source

	// clone repo on ram
	filesys := memfs.New()
	_, err := git.Clone(memory.NewStorage(), filesys, &git.CloneOptions{
		URL:   repoURL,
		Depth: 1,
	})
	if err != nil {
		panic(err)
	}

	librariesDotJsonBillyData, err := filesys.Open("/libraries.json")
	if err != nil {panic(err)}
	librariesDotJsonData, err := io.ReadAll(librariesDotJsonBillyData)
	if err != nil {panic(err)}
	var librariesDotJsonParsed create.LibrariesDotJson
	err = json.Unmarshal(librariesDotJsonData, &librariesDotJsonParsed)
	if err != nil {panic(err)}
	
	return repoURL, librariesDotJsonParsed, filesys, err
}


func gameVersionIncompat(fs billy.Filesystem) (isIncompatible bool, existingProjectVersion string, intendedGameVersion string ) {
	var libraryForm create.LibrariesDotJson
	var existingProjectForm create.LibrariesDotJson

	// parse the version string of current project, then parse the version string of the library, compare game_version field.
	log.Print(fs)
	library, err := fs.Open("libraries.json")
	if err != nil {
		log.Error("Library doesn't have a libraries.json file.")
		panic(err)
	}

	libraryContent, err := io.ReadAll(library)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(libraryContent, &libraryForm)
	if err != nil {
		panic(err)
	}

	existingProjectForm = ReadLocalLibrariesJson()

	if strings.EqualFold(existingProjectForm.GameVersion, libraryForm.GameVersion) {
		return false, existingProjectForm.GameVersion, libraryForm.GameVersion
	}

	return true, existingProjectForm.GameVersion, libraryForm.GameVersion
}

// dont merge contents of these files from imported libraries.
func avoidFileName(fileName string) bool {
	var blacklistedFileNames = []string{
		"README.md",
		"LICENSE",
		"pack.mcmeta",
		"libraries.json",
	}
	if slices.Contains(blacklistedFileNames, fileName) {
		return true
	} else {
		return false
	}
}

func MergeIndividualFiles(fs billy.Filesystem, currentPath string, targetDir string) error {

	files, err := fs.ReadDir(currentPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {

		// skip LICENSE file
		if avoidFileName(file.Name()) {
			// LOGGER: Ignoring README.md
			log.Printf("⛔ Ignoring %s", file.Name())
			continue
		}

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
			log.Printf("🎉Merging done!")
		}

	}

	return nil
}

type FunctionTag struct {
	Replace bool     `json:"replace"`
	Values  []string `json:"values"`
}

func handleFileMerging(fs billy.Filesystem, sourcePath string, destinationPath string) error {

	srcFile, err := fs.Open(sourcePath)
	if err != nil {
		log.Error("⚠ Source File Doesn't Exist.")
		panic(err)
	}
	defer srcFile.Close()

	err = os.MkdirAll(filepath.Dir(destinationPath), 0755)
	if err != nil {
		panic(err)
	}

	// if file already exists, append contents to the top of the file.
	if fileInfo, err := os.Stat(destinationPath); err == nil {

		fileName := strings.Split(fileInfo.Name(), ".")
		fileExtension := ""
		for i := range fileName {

			if i == len(fileName)-1 {
				fileExtension = fileName[i]
			}

		}
		if fileExtension == "json" {

			// LOGGER: Merging to load.json & tick.json

			if fileInfo.Name() == "load.json" || fileInfo.Name() == "tick.json" {
				log.Printf("📩 Merging into %s", fileInfo.Name())

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

		existingData, err := os.ReadFile(destinationPath)
		if err != nil {
			fmt.Println("here is panic")
			panic(err)
		}

		newData, err := io.ReadAll(srcFile)
		if err != nil {
			panic(err)
		}

		combined := append(newData, []byte("\n# added by sculk ^^\n")...)
		combined = append(combined, existingData...)

		return os.WriteFile(destinationPath, combined, 0644)
	}

	// if file doesn't exist:
	destinationFile, err := os.OpenFile(destinationPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, srcFile)
	return err

}
