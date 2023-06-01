# cloudwego-api-gateway

This repository contains the code for the implementation of an API Gateway for Orbital 2023.

# Getting Started

## IDL

IDL files are stored in `/idl`.

## Hertz

Navigate to the `http-server` directory and generate the Hertz scaffolding code with the `hz new` command:

```shell
$ hz new -module "github.com/tim-pipi/cloudwego-api-gateway/http-server" -idl ../idl/[YOUR_IDL_FILE].thrift
$ go mod tidy
```

To update the code after changes in the IDL:

```shell
$ hz update -idl ../idl/[YOUR_IDL_FILE].thrift
```

Update the logic in `biz/handler/api/hello_api.go` (make the Remote Procedure Call).

## Kitex

Navigate to the `rpc-server` directory and generate the Kitex server scaffolding code with the `kitex` command:

```shell
$ kitex -module "github.com/tim-pipi/cloudwego-api-gateway/rpc-server" -service hello ../idl/[YOUR_IDL_FILE].thrift
```

Update the logic in `handler.go`.

## Service Registration and Discovery

Service registration and discovery is done using [etcd](https://etcd.io/docs/v3.5/)
and the [`registry-etcd`](https://github.com/kitex-contrib/registry-etcd) library.

### Installing etcd

- Download the latest version of `etcd` from the **Releases** [page](https://github.com/etcd-io/etcd/releases/).
- Add the directory to your System's PATH. See [this guide](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/) for instructions

## Running

In the `http-server` directory: `go run .`
In the `rpc-server` directory: `go run .`
In your terminal, run: `etcd --advertise-client-urls http://localhost:7000 --listen-client-urls http://127.0.0.1:7000`

Test by using Postman/Insomnia with the following request: `http://127.0.0.1:8080/hello` with the following JSON body

```json
{
	"Name": "Timothy"
}
```
