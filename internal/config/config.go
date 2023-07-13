package config

import (
	"os"
)

// ServiceConfig represents the configuration for the API Gateway HTTP server
type ServiceConfig struct {
	IDLDir   string
	EtcdAddr string
	LogLevel string
	LogPath  string
}

// Reads the ServiceConfig from the environment.
// If the environment variable is not set, it will return the default value.
func ReadConfig() *ServiceConfig {
	idlDir := getEnv("IDL_DIR", "/etc/idl")
	etcdAddr := getEnv("ETCD_ADDR", "localhost:2379")
	logLevel := getEnv("LOG_LEVEL", "info")
	logPath := getEnv("LOG_PATH", "/var/log/cloudwego-api-gateway.log")

	return &ServiceConfig{
		IDLDir:   idlDir,
		EtcdAddr: etcdAddr,
		LogLevel: logLevel,
		LogPath:  logPath,
	}
}

// Helper function to retrieve environment variables
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)

	if !exists || value == "" {
		return fallback
	}

	return value
}
