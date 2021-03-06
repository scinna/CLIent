package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

func findConfigFolder() string {
	var configPath string

	usr, _ := user.Current()
	if runtime.GOOS == "windows" {
		configPath = filepath.Join(usr.HomeDir, "AppData\\Roaming\\scinna")
	} else if runtime.GOOS == "darwin" {
		// @TODO: TEST on osx
		configPath = filepath.Join(usr.HomeDir, "Library/Preferences/scinna")
	} else {
		// Supposing linux, please don't try to run it on weird things like android
		configPath = filepath.Join(usr.HomeDir, ".config/scinna")
	}

	return configPath
}

func getConfigFile() string {
	path := findConfigFolder()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("Could not create config folder: ", err)
			os.Exit(2)
		}
	}

	return filepath.Join(path, "CLIent.json")
}

func createDefaultConfiguration() error {
	cfg := Config{
		DefaultTitle: StringCommand("${date \"+%Y-%m-%d %H:%M:%S\"}"),
		DefaultDescription: "",
		DefaultVisibility:  VisibilityFromString("UNLISTED"),
	}

	return WriteConfiguration(&cfg)
}