package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/cloudwego/kitex/client"
	"github.com/tim-pipi/cloudwego-api-gateway/kitex_gen/api"
	"github.com/tim-pipi/cloudwego-api-gateway/kitex_gen/api/echo"
)

func main() {
	// HTTP server
	h := server.Default()

	h.GET("/hello", func(c context.Context, ctx *app.RequestContext) {
		// ctx.String(consts.StatusOK, "Hello, World!")
		ctx.JSON(consts.StatusOK, utils.H{"message": "Hello, World!"})
	})

	h.Spin()


	// RPC client
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
