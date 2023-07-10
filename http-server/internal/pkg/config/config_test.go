package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test case for updating config with a new service name
func TestUpdate_NewThriftDir(t *testing.T) {
	var config ServiceConfig
	// Get current directory
	wd, _ := os.Getwd()

	err := config.Update(wd)

	assert.Nil(t, err)
	assert.Equal(t, config.IDLDir, wd)
}

func TestUpdate_ResolvesRelativePaths(t *testing.T) {
	var config ServiceConfig

	err := config.Update("../config")
	assert.Nil(t, err)

	wd, _ := os.Getwd()
	assert.Equal(t, config.IDLDir, wd)
}

func TestUpdate_InvalidDir(t *testing.T) {
	var config ServiceConfig
	err := config.Update("path/to/idl")

	assert.NotNil(t, err)
	assert.Equal(t, err, os.ErrNotExist)
}
