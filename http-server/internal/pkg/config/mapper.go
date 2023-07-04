package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/cloudwego/thriftgo/parser"
)

// ServiceConfig type
type ServiceConfig struct {
	LastUpdated             string            `json:"last_updated"`
	ServiceNameToThriftFile map[string]string `json:"service_to_thrift"`
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

// Reads the config file from the given path, creating a new one if it doesn't exist
func ReadConfig(path string) (*ServiceConfig, error) {
	out, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c ServiceConfig
	err = json.Unmarshal(out, &c)
	if err != nil {
		return nil, err
	}

	if c.ServiceNameToThriftFile == nil {
		return NewConfig(), nil
	}

	return &c, nil
}

// Creates a new config
func NewConfig() *ServiceConfig {
	c := new(ServiceConfig)
	c.ServiceNameToThriftFile = make(map[string]string)
	return c
}

// Updates the ServiceConfig with the given serviceName and idlPath
func (c *ServiceConfig) Update(serviceName string, idlPath string) error {
	_, ok := c.ServiceNameToThriftFile[serviceName]
	if ok {
		return fmt.Errorf("service already exists")
	}

	c.ServiceNameToThriftFile[serviceName] = idlPath
	return nil
}

// Writes the ServiceConfig to a specified path
func (c *ServiceConfig) Write(path string) error {
	c.LastUpdated = time.Now().Format("2006-01-02 15:04:05")

	b, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
