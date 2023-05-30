# cloudwego-api-gateway

This repository contains the code for the implementation of an API Gateway for Orbital 2023.

## Directory
IDL files are stored in `/idl`.

## Hertz
Navigate to the `http-server` directory and generate the Hertz scaffolding code with the `hz new` command:

```shell
$ hz new -module "github.com/tim-pipi/cloudwego-api-gateway/http-server" -idl ../idl/hello_api.thrift
$ go mod tidy
```

To update the code:

```shell
$ hz update -idl ../idl/hello_api.thrift
```

Generate Kitex client scaffolding code with `kitex` command:
```shell
$ kitex -module "github.com/tim-pipi/cloudwego-api-gateway/http-server" ../idl/echo.thrift
```

Update the logic in `biz/handler/api/hello_api.go` (make the Remote Procedure Call).

## Kitex
Navigate to the `rpc-server` directory and generate the Kitex server scaffolding code with the `kitex` command:
```shell
$ kitex -module "github.com/tim-pipi/cloudwego-api-gateway/rpc-server" -service hello ../idl/hello_api.thrift
```

Update the logic in `handler.go`.

## Running
In the `http-server` directory: `go run .`
In the `rpc-server` directory: `go run .`
Test by using Postman/Insomnia with the following request: `http://127.0.0.1:8080/hello?name=tim`