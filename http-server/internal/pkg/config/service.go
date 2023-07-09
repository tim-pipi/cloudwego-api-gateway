package config

import (
	"io/fs"
	"path/filepath"

	"github.com/cloudwego/thriftgo/parser"
)

// Returns a map of service name to IDL files paths for the specified directory
func GetServiceMapFromDir(idlDir string) (map[string]string, error) {
	idls := find(idlDir, ".thrift")

	serviceMap := make(map[string]string)

	for _, idl := range idls {
		serviceNames, err := GetServicesFromIDL(idl)

		if err != nil {
			return nil, err
		}

		for _, serviceName := range serviceNames {
			serviceMap[serviceName] = idl
		}
	}

	return serviceMap, nil
}

func find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

// Returns the services from the given IDL file
func GetServicesFromIDL(idlPath string) ([]string, error) {
	// Use thriftgo to parse the IDL file
	t, err := parser.ParseFile(idlPath, []string{""}, true)
	if err != nil {
		return nil, err
	}

	services := t.Services
	var serviceNames []string
	for _, service := range services {
		serviceNames = append(serviceNames, service.Name)
	}

	return serviceNames, nil
}
