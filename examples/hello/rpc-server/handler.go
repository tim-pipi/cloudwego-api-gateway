package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	api "github.com/tim-pipi/cloudwego-api-gateway/examples/hello/rpc-server/kitex_gen/api"
)

var _ = klog.Info

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// HelloMethod implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) HelloMethod(ctx context.Context, request *api.HelloReq) (resp *api.HelloResp, err error) {
	if request.Name == "" {
		err = fmt.Errorf("name is required")
		return
	}	
	
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