package config

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/cloudwego/thriftgo/parser"
)

// ServiceConfig type
type ServiceConfig struct {
	LastUpdated string `json:"last_updated"`
	ThriftDir   string `json:"thrift_dir"`
}

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

// Reads the config file from the given path
func ReadConfig(path string) (*ServiceConfig, error) {
	out, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c ServiceConfig
	if json.Unmarshal(out, &c); err != nil {
		return nil, err
	}

	return &c, nil
}

// Creates a new config
func NewConfig() *ServiceConfig {
	return new(ServiceConfig)
}

// Updates the ServiceConfig with the given serviceName and idlPath
func (c *ServiceConfig) Update(thriftDir string) error {
	fullThriftDir, err := filepath.Abs(thriftDir)

	if err != nil {
		return err
	}

	ok, err := exists(fullThriftDir)
	if err != nil {
		return err
	}

	if !ok {
		return os.ErrNotExist
	}

	c.ThriftDir = fullThriftDir
	return nil
}

// Writes the ServiceConfig to a specified path
func (c *ServiceConfig) Write(path string) error {
	c.LastUpdated = time.Now().Format("2006-01-02 15:04:05")

	b, err := json.Marshal(c)
	if err != nil {
		return err
	}

	if os.WriteFile(path, b, 0644); err != nil {
		return err
	}

	return nil
}
