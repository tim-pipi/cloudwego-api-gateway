# Setup

You can run the API Gateway locally or in a Docker container.

## Docker

The easiest way to run the CloudWeGo API Gateway is from a Docker container.

To get started, clone this repository.

```shell
$ cd ./examples/hello
$ docker-compose up
[+] Running 3/3
 ✔ Container hello-etcd-1         Created                                                                          0.0s
 ✔ Container hello-rpc-server-1   Recreated                                                                        0.1s
 ✔ Container hello-http-server-1  Recreated                                                                        0.1s
# ...
```

In a separate terminal window, run the following command to test the API gateway:

```shell
$ curl http://localhost:8080/hello
Hello Hertz!
```

## Running locally

### Prerequisites

To run the API Gateway locally, you will need to install the following dependencies:

- Install the latest version of [Go](https://golang.org/doc/install) and [Docker](https://docs.docker.com/get-docker/).
- Install [Kitex](https://www.cloudwego.io/docs/kitex/getting-started/) and [Hertz](https://www.cloudwego.io/docs/hertz/getting-started/)
- [thrift-gen-validator](https://github.com/cloudwego/thrift-gen-validator)
- [etcd](https://github.com/etcd-io/etcd/releases/) for Service Registry (see below for further guide)
- [Postman](https://www.postman.com/downloads/) or [Insomnia](https://insomnia.rest/download) for API testing.

#### Installing etcd

- Download the latest version of `etcd` from the [**Releases**](https://github.com/etcd-io/etcd/releases/) page.
- Add the directory to your System's `PATH`. See [**this guide**](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/) for instructions.

### Running the API Gateway

Run the following commands:

```shell
$ etcd --advertise-client-urls http://etcd:2379 --listen-client-urls http://127.0.0.1:2379
$ export IDL_DIR="$(pwd)/examples/hello/http-server/idl"
$ export LOG_PATH="$(pwd)/examples/hello/http-server/output.log"
$ cd ./cmd/cloudwego
$ go run main.go start
# ...
{"level":"info","msg":"HERTZ: Using network library=netpoll","time":"2023-07-14T12:37:35+08:00"}
{"level":"info","msg":"HERTZ: HTTP server listening on address=[::]:8080","time":"2023-07-14T12:37:35+08:00"}
```

In a separate terminal window, run the following command to test the API gateway:

```shell
$ curl http://localhost:8080/hello
Hello Hertz!
```

### Setting up the RPC Server

TODO

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
