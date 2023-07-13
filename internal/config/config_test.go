package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestReadConfig tests the ReadConfig function
func TestReadConfig(t *testing.T) {
	// Create test cases using different environment variable values
	testCases := []struct {
		desc             string
		idlDirEnv        string
		etcdAddrEnv      string
		logLevelEnv      string
		logPathEnv       string
		expectedIDLDir   string
		expectedEtcdAddr string
		expectedLogLevel string
		expectedLogPath  string
	}{
		{
			desc:             "All environment variables present",
			idlDirEnv:        "/path/to/idl",
			etcdAddrEnv:      "example.com:2379",
			logLevelEnv:      "debug",
			logPathEnv:       "/var/log/test.log",
			expectedIDLDir:   "/path/to/idl",
			expectedEtcdAddr: "example.com:2379",
			expectedLogLevel: "debug",
			expectedLogPath:  "/var/log/test.log",
		},
		{
			desc:             "Missing IDL_DIR environment variable",
			etcdAddrEnv:      "example.com:2379",
			logLevelEnv:      "debug",
			logPathEnv:       "/var/log/test.log",
			expectedIDLDir:   "/etc/idl",
			expectedEtcdAddr: "example.com:2379",
			expectedLogLevel: "debug",
			expectedLogPath:  "/var/log/test.log",
		},
		// Add more test cases here if needed
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// Set the environment variables for each test case
			os.Setenv("IDL_DIR", tc.idlDirEnv)
			os.Setenv("ETCD_ADDR", tc.etcdAddrEnv)
			os.Setenv("LOG_LEVEL", tc.logLevelEnv)
			os.Setenv("LOG_PATH", tc.logPathEnv)

			// Call ReadConfig and check the returned ServiceConfig values
			config := ReadConfig()

			require.Equal(t, tc.expectedIDLDir, config.IDLDir)
			require.Equal(t, tc.expectedEtcdAddr, config.EtcdAddr)
			require.Equal(t, tc.expectedLogLevel, config.LogLevel)
			require.Equal(t, tc.expectedLogPath, config.LogPath)
		})
	}
}
