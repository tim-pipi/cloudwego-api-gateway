package main

import (
	"context"
	api "github.com/tim-pipi/cloudwego-api-gateway/rpc-server/kitex_gen/api"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// HelloMethod implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) HelloMethod(ctx context.Context, request *api.HelloReq) (resp *api.HelloResp, err error) {
	// TODO: Your code here...
	resp = &api.HelloResp{
		RespBody: "hello, " + request.Name,
	}
	return
}
