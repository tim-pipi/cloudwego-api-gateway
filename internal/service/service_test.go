package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseService_ExistingIDL(t *testing.T) {
	// Invoke the ParseService function with an existing IDL file
	wd, _ := os.Getwd()
	idlPath := wd + "/TestService.thrift"
	services, err := GetServicesFromIDL(idlPath)

	assert.Nil(t, err)
	assert.Equal(t, len(services), 1)
	assert.Equal(t, services[0].Name, "TestService")
}
