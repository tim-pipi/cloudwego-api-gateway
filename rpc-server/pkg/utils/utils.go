package utils

import (
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

type Counter struct {
	counter int
	mu      sync.Mutex
}

func (c *Counter) Increment() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter += 1
	return c.counter
}

func FindAvailablePort() (*net.TCPAddr, error) {
	// Create a range of ports to try
	startPort := 2000
	endPort := 65535

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	port := r.Intn(endPort-startPort+1) + startPort
	tries := 0
	MAX_TRIES := 1000

	for {
		addr := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", addr)
		if err == nil {
			// Port is available, close the listener and return the port number
			_ = listener.Close()

			return net.ResolveTCPAddr("tcp", addr)
		}
		port = rand.Intn(endPort-startPort+1) + startPort

		tries++
		if tries > MAX_TRIES {
			break
		}
	}

	// No available port found
	return nil, fmt.Errorf("no available port found")
}
