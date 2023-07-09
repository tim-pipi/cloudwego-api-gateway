package config

import (
	"os"
	"testing"

	"github.com/cloudwego/thriftgo/pkg/test"
)

// Test case for updating config with a new service name
func TestUpdate_NewThriftDir(t *testing.T) {
	var config ServiceConfig
	// Get current directory
	wd, _ := os.Getwd()

	err := config.Update(wd)

	test.Assert(t, err == nil)
	test.Assert(t, config.ThriftDir == wd)
}

func TestUpdate_ResolvesRelativePaths(t *testing.T) {
	var config ServiceConfig

	err := config.Update("../config")
	test.Assert(t, err == nil)

	wd, _ := os.Getwd()
	test.Assert(t, config.ThriftDir == wd)
}

func TestUpdate_InvalidDir(t *testing.T) {
	var config ServiceConfig
	err := config.Update("path/to/idl")

	test.Assert(t, err != nil)
	test.Assert(t, err == os.ErrNotExist)
}

func TestParseService_ExistingIDL(t *testing.T) {
	// Invoke the ParseService function with an existing IDL file
	wd, _ := os.Getwd()
	idlPath := wd + "/TestService.thrift"
	serviceNames, err := GetServicesFromIDL(idlPath)

	test.Assert(t, err == nil)
	test.Assert(t, len(serviceNames) == 1)
	test.Assert(t, serviceNames[0] == "TestService")
}
