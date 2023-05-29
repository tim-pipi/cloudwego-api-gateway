package main

import (
	"context"
	hello "github.com/tim-pipi/cloudwego-api-gateway/rpc-server/kitex_gen/hello"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *hello.Request) (resp *hello.Response, err error) {
	resp = &hello.Response{
		Message: req.Message,
	}
	return // Implicit return
}
