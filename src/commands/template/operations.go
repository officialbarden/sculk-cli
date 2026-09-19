package template

import (
	"os"
	"path/filepath"

	"charm.land/log/v2"
	"github.com/otiai10/copy"
)

func CreateTemplate(templateName string) {
	// working dir
	workingDir, err := os.Getwd()
	if err != nil { panic(err) }

	// sculk-project check
	librariesPath := filepath.Join(workingDir, "libraries.json")
	_, fileExistsCheck := os.Stat(librariesPath)
	if fileExistsCheck != nil {
		// file exists
		log.Error("Cannot create template, Project is not a 'sculk' project (missing libraries.json).")
		return
	}
	
	// C:\Users\User\AppData\Local
	path, err := os.UserCacheDir()
	if err != nil { panic(err) }
	
	// C:\Users\User\AppData\Local\sculk
	templateDir := filepath.Join(path, "sculk", "templates", templateName)
	templateExistsFile := filepath.Join(templateDir, "libraries.json")	

	_, templateExists := os.Stat(templateExistsFile)	
	if templateExists == nil {
		// file exists
		log.Error("A template with this name already exists. Delete the template OR use a unique name.")
		return
	}
	
	err = os.MkdirAll(templateDir, 0755)
	if err != nil { panic(err) }

	// copy paste:
	err = copy.Copy(workingDir, templateDir)
	if err != nil { panic(err) }

	log.Printf("📂 New Template '%s' created.", templateName)
}
func AddTemplate(templateName string) {
	
}
func DeleteTemplate(templateName string) {
	cacheDir, err := os.UserCacheDir()
	if err != nil { panic(err) }

	templateDir := filepath.Join(cacheDir, "sculk", "templates", templateName)
	err = os.RemoveAll(templateDir)
	if err != nil { panic(err) }

	log.Printf("🗑️ Template '%s' has been deleted.", templateName)
}
