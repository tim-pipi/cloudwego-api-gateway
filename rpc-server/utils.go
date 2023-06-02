package main

import (
	"sync"
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
