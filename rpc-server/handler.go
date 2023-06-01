package main

import (
	"context"
	api "github.com/tim-pipi/cloudwego-api-gateway/rpc-server/kitex_gen/api"
	"github.com/cloudwego/kitex/pkg/klog"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// HelloMethod implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) HelloMethod(ctx context.Context, request *api.HelloReq) (resp *api.HelloResp, err error) {
	// TODO: Your code here...
	klog.Info("Hello0 Called")
	resp = &api.HelloResp{
		RespBody: "hello, " + request.Name,
	}
	return
}

type HelloServiceImpl1 struct{}

func (s *HelloServiceImpl1) HelloMethod(ctx context.Context, request *api.HelloReq) (resp *api.HelloResp, err error) {
	klog.Info("Hello1 Called")
	resp = &api.HelloResp{
		RespBody: "hello, " + request.Name,
	}
	return
}