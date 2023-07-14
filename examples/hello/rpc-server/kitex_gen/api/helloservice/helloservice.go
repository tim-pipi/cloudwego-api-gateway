// Code generated by Kitex v0.6.0. DO NOT EDIT.

package helloservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	api "github.com/tim-pipi/cloudwego-api-gateway/examples/hello/rpc-server/kitex_gen/api"
)

func serviceInfo() *kitex.ServiceInfo {
	return helloServiceServiceInfo
}

var helloServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "HelloService"
	handlerType := (*api.HelloService)(nil)
	methods := map[string]kitex.MethodInfo{
		"HelloMethod": kitex.NewMethodInfo(helloMethodHandler, newHelloServiceHelloMethodArgs, newHelloServiceHelloMethodResult, false),
		"echo":        kitex.NewMethodInfo(echoHandler, newHelloServiceEchoArgs, newHelloServiceEchoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.0",
		Extra:           extra,
	}
	return svcInfo
}

func helloMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.HelloServiceHelloMethodArgs)
	realResult := result.(*api.HelloServiceHelloMethodResult)
	success, err := handler.(api.HelloService).HelloMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloServiceHelloMethodArgs() interface{} {
	return api.NewHelloServiceHelloMethodArgs()
}

func newHelloServiceHelloMethodResult() interface{} {
	return api.NewHelloServiceHelloMethodResult()
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.HelloServiceEchoArgs)
	realResult := result.(*api.HelloServiceEchoResult)
	success, err := handler.(api.HelloService).Echo(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloServiceEchoArgs() interface{} {
	return api.NewHelloServiceEchoArgs()
}

func newHelloServiceEchoResult() interface{} {
	return api.NewHelloServiceEchoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) HelloMethod(ctx context.Context, request *api.HelloReq) (r *api.HelloResp, err error) {
	var _args api.HelloServiceHelloMethodArgs
	_args.Request = request
	var _result api.HelloServiceHelloMethodResult
	if err = p.c.Call(ctx, "HelloMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Echo(ctx context.Context, request *api.EchoReq) (r *api.EchoResp, err error) {
	var _args api.HelloServiceEchoArgs
	_args.Request = request
	var _result api.HelloServiceEchoResult
	if err = p.c.Call(ctx, "echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
