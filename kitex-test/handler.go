package main

import (
	"context"
	api "github.com/tim-pipi/cloudwego-api-gateway/kitex-test/kitex_gen/api"
)

// EchoServiceImpl implements the last service interface defined in the IDL.
type EchoServiceImpl struct{}

// Echo implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) Echo(ctx context.Context, request *api.EchoReq) (resp *api.EchoResp, err error) {
	// TODO: Your code here...
	return
}
