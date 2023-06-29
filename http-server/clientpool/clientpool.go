package clientpool

import (
	"context"
	"strings"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type ClientPool struct {
	serviceMap map[string]genericclient.Client
	mutex      sync.Mutex
}

var (
	clientPool = &ClientPool{
		serviceMap: make(map[string]genericclient.Client),
	}
)

// Performs the generic call to the RPC server
// Calls the same client for the same service name
func Call(ctx context.Context, c *app.RequestContext, idlPath string) {
	jsonBody := string(c.Request.BodyBytes())
	klog.Info("Request Body: ", jsonBody)

	cli := getClient(c, idlPath)
	serviceMethod := getServiceMethod(c)

	// Make the Generic Call
	resp, err := cli.GenericCall(ctx, serviceMethod, jsonBody)
	if err != nil {
		klog.Info("remote procedure call failed: %v", err)
		// Retries the request if error
		// This is because the connection to the RPC server has a timeout
		// if the connection is idle for a long time.
		resp, _ = cli.GenericCall(ctx, serviceMethod, jsonBody)
	}
	klog.Info("Response body: ", resp)

	respString, ok := resp.(string)
	if !ok {
		klog.Error("Response is not a string:", resp)
	}

	c.String(consts.StatusOK, respString)
	c.SetContentType("application/json")
}

// getClient returns the same client for the same service name
func getClient(c *app.RequestContext, idlPath string) genericclient.Client {
	clientPool.mutex.Lock()
	defer clientPool.mutex.Unlock()

	serviceName := getServiceName(c)
	client, ok := clientPool.serviceMap[serviceName]

	if !ok {
		client = newClient(idlPath, serviceName)
		clientPool.serviceMap[serviceName] = client
	}

	return client
}

// getServiceName returns the service name from the full path of the url
func getServiceName(c *app.RequestContext) string {
	fullPath := c.FullPath()

	service := strings.Split(fullPath, "/")
	serviceName := service[1]
	klog.Info("Service Name: ", serviceName)

	return serviceName
}

// getServiceMethod returns the service method from the full path of the url
func getServiceMethod(c *app.RequestContext) string {
	fullPath := c.FullPath()

	service := strings.Split(fullPath, "/")
	serviceMethod := service[2]
	klog.Info("Service Method: ", serviceMethod)

	return serviceMethod
}

func newClient(idlPath string, serviceName string) genericclient.Client {
	p, err := generic.NewThriftFileProvider(idlPath)
	if err != nil {
		klog.Fatalf("new thrift file provider failed: %v", err)
	}

	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		klog.Fatalf("new http pb thrift generic failed: %v", err)
	}

	r, err := etcd.NewEtcdResolver([]string{"localhost:2379"})
	if err != nil {
		klog.Fatalf("new etcd resolver failed: %v", err)
	}

	lb := loadbalance.NewWeightedRoundRobinBalancer()

	cli, err := genericclient.NewClient(serviceName, g,
		kclient.WithResolver(r),
		kclient.WithLoadBalancer(lb),
	)

	if err != nil {
		klog.Fatalf("new http generic client failed: %v", err)
	}

	return cli
}
