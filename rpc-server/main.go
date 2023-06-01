package main

import (
    "github.com/cloudwego/kitex/pkg/rpcinfo"
    "github.com/cloudwego/kitex/server"
    etcd "github.com/kitex-contrib/registry-etcd"

	api "github.com/tim-pipi/cloudwego-api-gateway/rpc-server/kitex_gen/api/helloservice"
	"log"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"localhost:7000"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	svr := api.NewServer(new(HelloServiceImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "HelloService",
	}))

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
