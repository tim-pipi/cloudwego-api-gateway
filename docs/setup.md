# Getting started

Welcome to the guide for getting started with the CloudWeGo API Gateway!
This page will walk you through setting up and running the API Gateway using simple examples from the `/examples` directory.
You have the flexibility to run the API Gateway locally or within a Docker container.

## Prerequisites

Before you get started, ensure that you have the following installed on your development machine:

- [Go](https://golang.org/doc/install) v1.20.0 or higher
- [Docker](https://docs.docker.com/get-docker/)
- [Kitex](https://www.cloudwego.io/docs/kitex/getting-started/)
- [Hertz](https://www.cloudwego.io/docs/hertz/getting-started/)
- [thrift-gen-validator](https://github.com/cloudwego/thrift-gen-validator)

## Docker

The easiest way to get started with the CloudWeGo API Gateway is to run one of the
examples from the `/examples` directory.

### Step 1: Clone Repository

To start, clone this repository:

```shell
$ git clone https://github.com/tim-pipi/cloudwego-api-gateway.git
$ cd cloudwego-api-gateway/examples/hello
```

### Step 2: Launch Docker

Launch the Docker container:

```shell
$ docker-compose up
[+] Running 3/3
 ✔ Container hello-etcd-1         Created                                                                          0.0s
 ✔ Container hello-rpc-server-1   Recreated                                                                        0.1s
 ✔ Container hello-http-server-1  Recreated                                                                        0.1s
# ...
```

### Step 3: Test the API Gateway

In a separate terminal window, test the API Gateway by making a `GET` request.

```shell
$ curl http://localhost:8080/ping
Hello, CloudWeGo!
```

## Running locally

When running locally, you need to ensure that the following dependencies are installed.

- [etcd](https://github.com/etcd-io/etcd/releases/) for Service Registry
- Optional: [Postman](https://www.postman.com/downloads/), [Insomnia](https://insomnia.rest/download), or [Hurl](https://hurl.dev/) for API testing

### Step 1: Start the Service Registry

Start `etcd` to initiate the service registry.

```shell
$ etcd --advertise-client-urls http://etcd:2379 --listen-client-urls http://127.0.0.1:2379
```

### Step 2: Set Environemnt Variables

Set the required environment variables for the API Gateway.

```shell
$ export IDL_DIR="$(pwd)/tests/hello/idl"
$ export LOG_PATH="$(pwd)/tests/hello/output.log"
```

### Step 3: Start the API Gateway

```shell
$ go run cmd/cloudwego/main.go start
# ...
{"level":"info","msg":"HERTZ: Using network library=netpoll","time":"2023-07-14T12:37:35+08:00"}
{"level":"info","msg":"HERTZ: HTTP server listening on address=[::]:8080","time":"2023-07-14T12:37:35+08:00"}
```

### Step 4: Start the RPC Server

```shell
$ cd examples/hello/rpc-server
$ go run .
```

### Step 5: Test the API Gateway

In a separate terminal window, test the API Gateway by making a `GET` request:

```shell
$ curl http://localhost:8080/ping
Hello, CloudWeGo!
```

Alternatively, you can use `hurl` to run tests.

```shell
$ cd tests/hello
$ hurl --test hello_service_test.hurl
hello_service_test.hurl: Running [1/1]
hello_service_test.hurl: Success (6 request(s) in 20 ms)
--------------------------------------------------------------------------------
Executed files:  1
Succeeded files: 1 (100.0%)
Failed files:    0 (0.0%)
Duration:        23 ms
```

Congratulations! You have successfully set up and run the CloudWeGo API Gateway using the provided examples.
You are now ready to explore the capabilities of the API Gateway and start building powerful APIs for your applications.
