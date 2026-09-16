package uninstall

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"sculk-cli/src/commands/add"
	"strings"

	"charm.land/log/v2"
	"github.com/go-git/go-billy/v6"
)

// Similar to `sculk add` command, walk to every file, then delete instead of merging.
func Main(args []string) {

	libraryIdentifiers := args
	for _, libraryIdentifier := range libraryIdentifiers {
		// drop repo-link and librariesjson information
		_, _, sourceCodeFilesystem, err := add.GetLibrarySource(libraryIdentifier)
		if err != nil {panic(err)}
		
		getWorkingDir, err := os.Getwd()
		// start from '/'
		err = TraverseSourceCode(sourceCodeFilesystem, "/", getWorkingDir)
		if err != nil {panic(err)}
		err = add.RemoveFromLibrariesJson(libraryIdentifier)
		if err != nil {panic(err)}
		
	}

	log.Printf("Note: Please delete the empty directories manually.")
}

func TraverseSourceCode(fs billy.Filesystem, currentPath string, targetDir string) error {
	files, err := fs.ReadDir(currentPath)

	for _, file := range files {
		// store memory path and local path
		memoryPath := filepath.Join(currentPath, file.Name())
		localPath := filepath.Join(targetDir, file.Name())

		if file.IsDir() {
			// if found directory, go inside that directory and recurse.
			err = TraverseSourceCode(fs, memoryPath, localPath)
			if err != nil {
				panic(err)
			}

		} else {

			// blacklisted filenames
			blacklistedFiles := []string{
				"libraries.json",
			}
			
			err := handleFileDeletion(fs, memoryPath, localPath, blacklistedFiles)
			if err != nil {
				panic(err)
			}
			log.Printf("🚮 Deleted file '%s'.", file.Name())
		}
	}
	return nil
}


func handleFileDeletion(fs billy.Filesystem, sourcePath string, destinationPath string, blacklist []string) error {

	// open source file (contains the lines/entries to remove)
	srcFile, err := fs.Open(sourcePath)
	if err != nil {
		log.Error("⚠ Source File Doesn't Exist.")
		return err
	}
	defer srcFile.Close()

	// nothing to delete from if destination doesn't exist
	fileInfo, err := os.Stat(destinationPath)
	if err != nil {
		log.Printf("⚠ Destination File Doesn't Exist, nothing to delete: %s", destinationPath)
		return nil
	}

	if isBlacklisted(fileInfo.Name(), blacklist) {
		log.Printf("🚫 %s is blacklisted, skipping deletion", fileInfo.Name())
		return nil
	}

	fileExtension := strings.TrimPrefix(filepath.Ext(fileInfo.Name()), ".")

	sourceContent, err := io.ReadAll(srcFile)
	if err != nil {
		return err
	}

	rootDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if fileExtension == "json" {

		if fileInfo.Name() == "load.json" || fileInfo.Name() == "tick.json" {
			log.Printf("🗑 Removing entries from %s", fileInfo.Name())

			var sourceTagContent add.FunctionTag
			var existingTagContent add.FunctionTag

			existingContent, err := os.ReadFile(destinationPath)
			if err != nil {
				return err
			}

			if err := json.Unmarshal(sourceContent, &sourceTagContent); err != nil {
				return err
			}
			if err := json.Unmarshal(existingContent, &existingTagContent); err != nil {
				return err
			}

			// build a lookup of values to remove
			toRemove := make(map[string]bool, len(sourceTagContent.Values))
			for _, v := range sourceTagContent.Values {
				toRemove[v] = true
			}

			filtered := existingTagContent.Values[:0]
			for _, v := range existingTagContent.Values {
				if !toRemove[v] {
					filtered = append(filtered, v)
				}
			}
			existingTagContent.Values = filtered

			// if nothing left, delete the file entirely (and any now-empty parent dirs)
			if len(existingTagContent.Values) == 0 {
				log.Printf("🗑 %s is now empty, deleting file", fileInfo.Name())
				if err := os.Remove(destinationPath); err != nil {
					return err
				}
				return removeEmptyDirs(filepath.Dir(destinationPath), rootDir, blacklist)
			}

			combined, err := json.Marshal(existingTagContent)
			if err != nil {
				return err
			}
			return os.WriteFile(destinationPath, combined, 0644)
		} else {
			return os.Remove(destinationPath)
		}
	}

	// generic line-based deletion
	existingData, err := os.ReadFile(destinationPath)
	if err != nil {
		return err
	}

	// lines to remove, exact match
	toRemove := make(map[string]bool)
	for _, line := range strings.Split(string(sourceContent), "\n") {
		toRemove[line] = true
	}

	var kept []string
	for _, line := range strings.Split(string(existingData), "\n") {
		if !toRemove[line] {
			kept = append(kept, line)
		}
	}

	remaining := strings.TrimSpace(strings.Join(kept, "\n"))

	// if file becomes empty, delete it entirely (and any now-empty parent dirs)
	if remaining == "" {
		log.Printf("🗑 %s is now empty, deleting file", fileInfo.Name())
		if err := os.Remove(destinationPath); err != nil {
			return err
		}
		return removeEmptyDirs(filepath.Dir(destinationPath), rootDir, blacklist)
	}

	return os.WriteFile(destinationPath, []byte(remaining+"\n"), 0644)
}

func isBlacklisted(name string, blacklist []string) bool {
	for _, b := range blacklist {
		if b == name {
			return true
		}
	}
	return false
}

func removeEmptyDirs(dir string, root string, blacklist []string) error {
	root = filepath.Clean(root)

	for {
		dir = filepath.Clean(dir)

		if dir == root || dir == "." || dir == string(filepath.Separator) {
			return nil
		}
		rel, err := filepath.Rel(root, dir)
		if err != nil || strings.HasPrefix(rel, "..") {
			return nil
		}

		if isBlacklisted(filepath.Base(dir), blacklist) {
			log.Printf("🚫 %s is blacklisted, stopping cleanup here", dir)
			return nil
		}

		empty, err := isEmptyDirTree(dir, blacklist)
		if err != nil {
			if os.IsNotExist(err) {
				return nil
			}
			return err
		}

		if !empty {
			// contains at least one real (or blacklisted) file somewhere, stop climbing
			return nil
		}

		log.Printf("🗑 Removing empty directory tree %s", dir)
		if err := os.RemoveAll(dir); err != nil {
			return err
		}

		dir = filepath.Dir(dir)
	}
}

func isEmptyDirTree(dir string, blacklist []string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}

	for _, entry := range entries {
		if isBlacklisted(entry.Name(), blacklist) {
			return false, nil
		}

		if !entry.IsDir() {
			return false, nil
		}

		childEmpty, err := isEmptyDirTree(filepath.Join(dir, entry.Name()), blacklist)
		if err != nil {
			return false, err
		}
		if !childEmpty {
			return false, nil
		}
	}

	return true, nil
}