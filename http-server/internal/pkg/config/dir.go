package config

import (
	"io"
	"log"
	"os"
)

const CONFIG_SUBPATH = "/cwgo/idl"
const CONFIG_FILENAME = "config.json"

type UserConfigDir interface {
	get() string
}

type OsUserConfigDir struct{}

func (OsUserConfigDir) get() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	return configDir
}

var userConfigDir UserConfigDir = OsUserConfigDir{}

// Copies the given file to the config directory
func CopyToConfigDir(srcPath string, name string) error {
	configDir := GetConfigDir()
	destPath := configDir + "/" + name

	err := copyFile(srcPath, destPath)
	return err
}

// Copies the file from source to destination
func copyFile(srcPath string, destPath string) error {
	src, err := os.Open(srcPath)

	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}

	return nil
}

// Creates the config directory if it doesn't exist.
func CreateConfigDir() {
	exists := checkConfigDir()

	if !exists {
		configDir := GetConfigDir()
		err := os.MkdirAll(configDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		config := NewConfig()
		err = config.Write(GetConfigPath())

		if err != nil {
			log.Fatal(err)
		}
	}
}

// Returns the path to the config file on the given system
func GetConfigPath() string {
	configDir := GetConfigDir()
	configFile := configDir + "/" + CONFIG_FILENAME
	return configFile
}

// Returns the path to the config directory on the given system
func GetConfigDir() string {
	return userConfigDir.get() + "/cwgo/idl"
}

func checkConfigDir() bool {
	configDir := GetConfigDir()
	exists, err := exists(configDir)
	if err != nil {
		log.Fatal(err)
	}

	return exists
}

// Check if a given directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if !os.IsNotExist(err) {
		return false, err
	}

	return false, nil
}
