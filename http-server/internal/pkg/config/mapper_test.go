package config

import (
	"os"
	"reflect"
	"testing"

	"github.com/cloudwego/thriftgo/pkg/test"
)

// Test case for updating config with a new service name
func TestUpdate_NewServiceName(t *testing.T) {
	// Invoke the UpdateConfig function with a new service name
	var config ServiceConfig
	config.ServiceNameToThriftFile = make(map[string]string)

	err := config.Update("NewService", "path/to/idl")
	test.Assert(t, err == nil)

	data := make(map[string]string)
	data["NewService"] = "path/to/idl"

	test.Assert(t, reflect.DeepEqual(config.ServiceNameToThriftFile, data))
}

// Test case for updating config with an existing service name
func TestUpdate_ExistingServiceName(t *testing.T) {
	// Invoke the UpdateConfig function with an existing service name
	var config ServiceConfig
	config.ServiceNameToThriftFile = make(map[string]string)
	config.ServiceNameToThriftFile["ExistingService"] = "path/to/idl"

	err := config.Update("ExistingService", "new/path/to/idl")

	errMsg := "service already exists"
	test.Assert(t, err != nil)
	test.Assert(t, err.Error() == errMsg)
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
