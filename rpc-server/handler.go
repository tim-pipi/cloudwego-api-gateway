package main

import (
	"context"
	api "github.com/tim-pipi/cloudwego-api-gateway/rpc-server/kitex_gen/api"
	"github.com/cloudwego/kitex/pkg/klog"
)

var _ = klog.Info

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

// Echo Method returns the message sent by the client
func (s *HelloServiceImpl) Echo(ctx context.Context, request *api.EchoReq) (resp *api.EchoResp, err error) {
	resp = &api.EchoResp{
		Response: request.Message,
	}
	return
}