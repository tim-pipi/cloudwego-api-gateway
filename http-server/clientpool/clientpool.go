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
	"github.com/tim-pipi/cloudwego-api-gateway/http-server/internal/pkg/service"

	kitextracing "github.com/kitex-contrib/obs-opentelemetry/tracing"
)

type ClientPool struct {
	serviceMap map[string]genericclient.Client
	mutex      sync.Mutex
}

func NewClientPool(idlDir string) *ClientPool {
	clientPool := &ClientPool{
		serviceMap: make(map[string]genericclient.Client),
	}

	services, err := service.GetServicesFromIDLDir(idlDir)

	if err != nil {
		klog.Fatalf("Error getting service map from IDL directory: %v", err)
	}

	for _, svc := range services {
		clientPool.serviceMap[svc.Name] = newClient(svc.Path, svc.Name)
	}

	return clientPool
}

// Performs the generic call to the RPC server
// Calls the same client for the same service name
func (cp *ClientPool) Call(ctx context.Context, c *app.RequestContext) {
	jsonBody := string(c.Request.BodyBytes())
	klog.Info("Request Body: ", jsonBody)

	cli := cp.getClient(c)
	serviceMethod := c.Param("ServiceMethod")
	klog.Info("Service Method: ", serviceMethod)

	// Kitex RPC server will return error
	resp, err := cli.GenericCall(ctx, serviceMethod, jsonBody)
	if err != nil {
		klog.Error("Remote procedure call failed: ", err)

		// Append to errors
		c.Error(err)
		errorJSON := map[string]interface{}{
			"error": c.Errors.Errors(),
		}

		if strings.Contains(err.Error(), "missing method") {
			c.JSON(consts.StatusNotFound, errorJSON)
			return
		}
		c.JSON(consts.StatusBadRequest, errorJSON)

		return
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
func (cp *ClientPool) getClient(c *app.RequestContext) genericclient.Client {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()

	serviceName := c.Param("ServiceName")
	klog.Info("Service Name: ", serviceName)
	client, ok := cp.serviceMap[serviceName]

	// TODD: Return error for getClient
	if !ok {
		return nil
	}

	return client
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
		kclient.WithSuite(kitextracing.NewClientSuite()),
	)

	if err != nil {
		klog.Fatalf("new http generic client failed: %v", err)
	}

	klog.Info("Created new client for service: ", serviceName)
	return cli
}
