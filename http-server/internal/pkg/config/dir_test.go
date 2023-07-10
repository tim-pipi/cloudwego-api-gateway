package config

import (
	"log"
	"os"
	"testing"

	"github.com/cloudwego/thriftgo/pkg/test"
)

const TEST_DIRECTORY = "/dir_test"

type testUserConfigDir struct{}

func (testUserConfigDir) get() string {
	// Get path of the current directory
	dir, _ := os.Getwd()
	return dir + TEST_DIRECTORY
}

func setup() {
	// Create directory for testUserConfig
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	os.Mkdir(dir+TEST_DIRECTORY, os.ModePerm)
}

func teardown() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	os.RemoveAll(dir + TEST_DIRECTORY)
}

func Test_CreateConfigDir(t *testing.T) {
	// Invoke the CreateConfigDir function
	currConfigDir := userConfigDir
	userConfigDir = testUserConfigDir{}
	defer func() { userConfigDir = currConfigDir }()

	setup()
	defer teardown()

	CreateConfigDir()

	configDir := GetConfigDir()
	e, err := exists(configDir)
	if err != nil {
		log.Fatal(err)
	}

	test.Assert(t, e)

	configFile := configDir + "/" + CONFIG_FILENAME
	e, err = exists(configFile)
	if err != nil {
		log.Fatal(err)
	}
	test.Assert(t, e)
}

