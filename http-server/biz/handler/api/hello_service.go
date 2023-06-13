// Code generated by hertz generator.
package api

import (
	"context"

	kclient "github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"

	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api "github.com/tim-pipi/cloudwego-api-gateway/http-server/biz/model/api"
)

type HelloClient struct {
	genericclient.Client
}

func NewHelloClient(idlPath string) *HelloClient {
	p, err := generic.NewThriftFileProvider(idlPath)
	if err != nil {
		klog.Fatalf("new thrift file provider failed: %v", err)
	}

	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		klog.Fatalf("new http pb thrift generic failed: %v", err)
	}

	r, err := etcd.NewEtcdResolver([]string{"localhost:7000"})
	if err != nil {
		klog.Fatalf("new etcd resolver failed: %v", err)
	}

	lb := loadbalance.NewWeightedRoundRobinBalancer()

	cli, err := genericclient.NewClient("HelloService", g,
		kclient.WithResolver(r),
		kclient.WithLoadBalancer(lb),
	)
	if err != nil {
		klog.Fatalf("new http generic client failed: %v", err)
	}

	return &HelloClient{cli}
}

func (cli *HelloClient) Call(ctx context.Context, c *app.RequestContext, method string) (err error) {
	jsonBody := string(c.Request.BodyBytes())

	// Make the Generic Call
	resp, err := cli.GenericCall(ctx, method, jsonBody)
	if err != nil {
		klog.Info("remote procedure call failed: %v", err)
		// Retries the request if error
		// This is because the connection to the RPC server has a timeout
		// if the connection is idle for a long time.
		resp, _ = cli.GenericCall(ctx, method, jsonBody)
	}

	c.JSON(consts.StatusOK, resp)

	return
}

var cli = NewHelloClient("../idl/hello_api.thrift")

// HelloMethod .
// @router /hello [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	// Validate the request body
	var err error
	var req api.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	klog.Info("req path: ", string(c.Path()))
	klog.Info("req full path: ", c.FullPath())

	cli.Call(ctx, c, "HelloMethod")
}

// Echo .
// @router /HelloService/echo [GET]
func Echo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.EchoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	
	cli.Call(ctx, c, "echo")
}
