package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"

	"log"

	api "github.com/tim-pipi/cloudwego-api-gateway/rpc-server/kitex_gen/api/helloservice"
	"github.com/tim-pipi/cloudwego-api-gateway/rpc-server/middleware"

	"github.com/tim-pipi/cloudwego-api-gateway/rpc-server/pkg/utils"

	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

// Constants for testing purposes
const NUMSERVERS = 5

func main() {
	// Create the service registry
	r, err := etcd.NewEtcdRegistry([]string{"localhost:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	serviceName := "HelloService"

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	var wg sync.WaitGroup
	counter := new(utils.Counter)

	// Creates a new RPC server for the HelloService
	createHelloServer := func() {
		defer wg.Done()
		count := counter.Increment()

		addr, err := utils.FindAvailablePort()
		if err != nil {
			log.Println(err.Error())
			panic(err)
		}

		svr := api.NewServer(
			new(HelloServiceImpl),
			server.WithRegistry(r),
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
				ServiceName: "HelloService",
			}),
			server.WithServiceAddr(addr),
			server.WithMiddleware(middleware.MiddleWareLogger(fmt.Sprintf("HelloService: Server %d called", count))),
			server.WithMiddleware(middleware.ValidatorMW),
			server.WithSuite(tracing.NewServerSuite()),
		)

		if err := svr.Run(); err != nil {
			log.Println(err.Error())
		}
	}

	wg.Add(NUMSERVERS)
	for i := 0; i < NUMSERVERS; i++ {
		go createHelloServer()
	}

	wg.Wait()
}
