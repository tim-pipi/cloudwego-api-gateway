package main

import (
	"net"
	"sync"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"

	"log"

	api "github.com/tim-pipi/cloudwego-api-gateway/rpc-server/kitex_gen/api/helloservice"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"localhost:7000"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", ":7050")
	addr2, _ := net.ResolveTCPAddr("tcp", ":7051")

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		svr0 := api.NewServer(new(HelloServiceImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "HelloService",
		}), server.WithServiceAddr(addr))
		if err := svr0.Run(); err != nil {
			log.Println(err.Error())
		}
	}()
	go func() {
		defer wg.Done()
		svr1 := api.NewServer(new(HelloServiceImpl1), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "HelloService",
		}), server.WithServiceAddr(addr2))

		if err := svr1.Run(); err != nil {
			log.Println(err.Error())
		}
	}()

	wg.Wait()

}
