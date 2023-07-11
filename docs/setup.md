# Setup

This document outlines the steps required to set up the API Gateway.

## Prerequisites

- Install the latest version of [Go](https://golang.org/doc/install) and [Docker](https://docs.docker.com/get-docker/).
- Install [Kitex](https://www.cloudwego.io/docs/kitex/getting-started/) and [Hertz](https://www.cloudwego.io/docs/hertz/getting-started/)
- [thrift-gen-validator](https://github.com/cloudwego/thrift-gen-validator)
- [etcd](https://github.com/etcd-io/etcd/releases/) for Service Registry (see below for further guide)
- [Postman](https://www.postman.com/downloads/) or [Insomnia](https://insomnia.rest/download) for API testing.

### Installing etcd

- Download the latest version of `etcd` from the [**Releases**](https://github.com/etcd-io/etcd/releases/) page.
- Add the directory to your System's `PATH`. See [**this guide**](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/) for instructions.

## Getting Started

In the `http-server` directory: `go run .`

In the `rpc-server` directory: `go run .`

In your terminal, run: `etcd --advertise-client-urls http://localhost:7000 --listen-client-urls http://127.0.0.1:7000`

Test with Postman/Insomnia using the following request: `http://127.0.0.1:8080/HelloService/HelloMethod` with the following JSON body:

```json
{
  "Name": "Timothy"
}
```

You should receive the following response:

```json
{
  "RespBody": "hello, Timothy"
}
```
