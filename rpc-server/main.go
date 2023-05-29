package main

import (
	hello "github.com/tim-pipi/cloudwego-api-gateway/rpc-server/kitex_gen/hello/echo"
	"log"
)

func main() {
	svr := hello.NewServer(new(EchoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
