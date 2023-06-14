package utils_test

import (
	"net"
	"testing"

	"github.com/tim-pipi/cloudwego-api-gateway/rpc-server/pkg/utils"
)

func TestFindAvailablePort(t *testing.T) {
	addr, err := utils.FindAvailablePort()
	if err != nil {
		t.Errorf("findAvailablePort returned unexpected error: %v", err)
	}
	if addr == nil {
		t.Error("findAvailablePort returned nil address")
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		t.Errorf("failed to listen on returned address: %v", err)
	}
	defer listener.Close()

	// try listening again on the same address and port
	_, err = net.ListenTCP("tcp", addr)
	if err == nil {
		t.Error("failed to find available port")
	}
}
