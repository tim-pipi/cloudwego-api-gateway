package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"

	"log"

	api "github.com/tim-pipi/cloudwego-api-gateway/rpc-server/kitex_gen/api/helloservice"
	"github.com/tim-pipi/cloudwego-api-gateway/rpc-server/middleware"
)

// Constants for testing purposes
const PORT = 7050
const NUMSERVERS = 5

func main() {
	// Create the service registry
	r, err := etcd.NewEtcdRegistry([]string{"localhost:7000"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	counter := new(Counter)

	// Creates a new RPC server for the HelloService
	createHelloServer := func() {
		defer wg.Done()
		count := counter.Increment()

		// Runs on a different port for each server
		addr, err := net.ResolveTCPAddr("tcp", ":"+fmt.Sprintf("%d", PORT+count)) 
		if err != nil {
			log.Println(err.Error())
		}
		svr := api.NewServer(
			new(HelloServiceImpl),
			server.WithRegistry(r),
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
				ServiceName: "HelloService",
			}),
			server.WithServiceAddr(addr),
			// Middleware to log which server is being called
			server.WithMiddleware(middleware.MiddleWareLogger(fmt.Sprintf("HelloService: Server %d called", count))),
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
