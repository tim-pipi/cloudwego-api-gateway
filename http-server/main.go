package main

import (
	"context"
	"log"
	"time"

	"github.com/tim-pipi/cloudwego-api-gateway/kitex_gen/api"
	"github.com/tim-pipi/cloudwego-api-gateway/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := echo.NewClient("echo", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	// Run indefinitely
	for {
		req := &api.Request{Message: "Hello, World!"}
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp.Message)
		time.Sleep(1 * time.Second)
	}
}