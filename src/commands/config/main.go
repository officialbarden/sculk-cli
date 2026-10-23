package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"charm.land/log/v2"
)

type Config struct {
	DoMerge      bool            `json:"doMerge"`
	Author       Author          `json:"author"`
	Sculk        SculkVersioning `json:"sculk"`
	InitTemplate string          `json:"initTemplate"`
}

type SculkVersioning struct {
	Version string `json:"version"`
}

type Author struct {
	Name string `json:"name"`
}

func Main(args []string) {}

func InitConfig() {
	// config doesn't exist, create one.
	appDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	sculkConfigFilePath := filepath.Join(appDir, ".sculk")
	err = os.MkdirAll(sculkConfigFilePath, 0755)
	if err != nil {
		panic(err)
	}

	sculkConfigBase := Config{
		DoMerge: true, // Merge to Existing Datapack
		Author: Author{
			Name: "Sculk Author", // Author Name
		},
		Sculk:        SculkVersioning{Version: "1.0.0"},
		InitTemplate: "none", // Init Template, incase other people have different file-structure practices.
	}

	// indentation for pretty-printing.
	sculkConfigBaseJson, err := json.MarshalIndent(sculkConfigBase, "", "	")

	// write to config.
	err = os.WriteFile(filepath.Join(sculkConfigFilePath, "config.json"), sculkConfigBaseJson, 0644)
	if err != nil {
		panic(err)
	}
}

func GetSculkConfig() (configJsonData *Config, filePath string, error error) {
	appDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	sculkConfigFile := filepath.Join(appDir, ".sculk", "config.json")

	configDataFile, err := os.ReadFile(sculkConfigFile)
	if err != nil {
		return nil, "", errors.New("Config File Doesn't Exist.")
	}

	var configData Config
	err = json.Unmarshal(configDataFile, &configData)

	return &configData, sculkConfigFile, err
}

func ConfigExists() bool {
	_, _, err := GetSculkConfig()
	if err != nil {
		log.Print("Couldn't find config.json, creating a new base config. Use 'sculk config' to change values.")
		InitConfig()
		return false
	} else {
		return true
	}
}

func setParamValue(param string, value string) {}
