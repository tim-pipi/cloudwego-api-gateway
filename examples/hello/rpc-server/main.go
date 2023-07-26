package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"

	klog "github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"

	api "github.com/tim-pipi/cloudwego-api-gateway/examples/hello/rpc-server/kitex_gen/api/helloservice"
	"github.com/tim-pipi/cloudwego-api-gateway/examples/hello/rpc-server/middleware"
)

func main() {
	// Logging
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelDebug)

	logfile := os.Getenv("LOGFILE")
	if logfile == "" {
		logfile = "./output.log"
	}

	f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	mw := io.MultiWriter(os.Stderr, f)
	klog.SetOutput(mw)

	// Service Registry
	etcdURL := os.Getenv("ETCD_URL")
	if etcdURL == "" {
		etcdURL = "localhost:2379"
	}
	klog.Info("ETCD_URL: ", etcdURL)
	r, err := etcd.NewEtcdRegistry([]string{etcdURL}) // r should not be reused.
	if err != nil {
		klog.Fatal(err)
	}

	// Observability
	serviceName := "HelloService"
	allowMetrics := os.Getenv("ALLOW_METRICS")
	if allowMetrics == "1" {
		p := provider.NewOpenTelemetryProvider(
			provider.WithServiceName(serviceName),
			provider.WithExportEndpoint(":4317"),
			provider.WithInsecure(),
		)
		defer p.Shutdown(context.Background())
	}

	// Server
	container := os.Getenv("CONTAINER")
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:8888", container))
	if err != nil {
		klog.Fatal(err)
	}

	svr := api.NewServer(
		new(HelloServiceImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: serviceName,
		}),
		server.WithServiceAddr(addr),
		server.WithMiddleware(middleware.MiddleWareLogger(fmt.Sprintf("%s called", serviceName))),
		server.WithMiddleware(middleware.ValidatorMW),
		server.WithSuite(tracing.NewServerSuite()),
	)

	if err := svr.Run(); err != nil {
		klog.Fatal(err)
	}
}
