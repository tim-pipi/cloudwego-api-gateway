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

func NewHelloClient() genericclient.Client {
	idlPath := "../idl/hello_api.thrift"
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

	lb := loadbalance.NewWeightedBalancer()

	cli, err := genericclient.NewClient("HelloService", g, kclient.WithResolver(r), kclient.WithLoadBalancer(lb))
	if err != nil {
		klog.Fatalf("new http generic client failed: %v", err)
	}
	return cli
}

// HelloMethod .
// @router /hello [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	jsonBody := string(c.Request.BodyBytes())

	// Make the Generic Call
	cli := NewHelloClient()
	resp, err := cli.GenericCall(ctx, "HelloMethod", jsonBody)
	if err != nil {
		klog.Fatalf("remote procedure call failed: %v", err)
		return
	}

	c.JSON(consts.StatusOK, resp)
}
