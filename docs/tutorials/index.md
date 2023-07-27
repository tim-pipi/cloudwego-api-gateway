# Adding Services

Welcome to CloudWeGo API Gateway!

In this guide, we'll walk you through the process of creating your first service with our API Gateway.
The CloudWeGo API Gateway operates using the Thrift IDL format, which allows you to define the structure of the request and response objects, as well as the service interface.

Before we begin, ensure that you have followed the instructions on the [Setup](../setup.md) page.

## Adding Services

In this section, we will guide you on how to add services to the CloudWeGo API Gateway and properly organize your Thrift IDL files.

### Setting up the IDL directory

Before adding services, it is crucial to have a designated directory to store all your Thrift IDL files.
To create this directory and set it as an environment variable, follow these steps in your terminal:

```shell
$ mkdir idl
$ export IDL_DIR="$(pwd)/idl"
```

!!! note

    Your IDL files have to follow the [Thrift IDL Annotation Standard](https://www.cloudwego.io/docs/kitex/tutorials/advanced-feature/generic-call/thrift_idl_annotation_standards/) for Kitex.

### Creating your first service: Echo Service

To demonstrate the process of adding a service, we will create a simple service named "EchoService."
Copy and paste the following code into a new file named /idl/echo.thrift:

```thrift
namespace go api

struct EchoReq {
    1:required string message
}

struct EchoResp {
    1: string response
}

service EchoService {
    EchoResp echo(1: EchoReq request) (api.get="/EchoService/echo")
}
```

### Understanding the IDL file

The /idl/echo.thrift file defines the "EchoService" with its request and response structures.
Let's go through the code:

| Code                            | Explanation                                                                                      |
| ------------------------------- | ------------------------------------------------------------------------------------------------ |
| `struct EchoReq...`             | Defines the structure representing the request object for the "Echo" operation.                  |
| `struct EchoResp...`            | Defines the structure representing the response object for the "Echo" operation.                 |
| `service EchoService...`        | Declares the service interface named "EchoService" that includes various operations.             |
| `(api.get="/EchoService/echo")` | Specifies the path for the HTTP GET request to access the "echo" operation of the "EchoService". |

## Generating the RPC Server using Kitex

To generate the scaffolding code for the Kitex RPC server, follow these steps:

Create a new directory named `/rpc-server`.

Run the `cwgo gen` command with the `-i` flag to specify the path to the Thrift IDL file, and the `-m` flag to set the module name:

```shell
$ cwgo gen -i idl/echo.thrift -m {MODULE_NAME}
Adding apache/thrift@v0.13.0 to go.mod for generated code .......... Done
```

After running the command, the /rpc-server directory will contain the following files and subdirectories:

```shell
├── Dockerfile
├── README.md
├── handler.go
├── idl
│   └── echo.thrift
├── kitex-template
│   ├── handler_tpl.yaml
│   ├── main_tpl.yaml
│   ├── middleware_tpl.yaml
│   └── readme_tpl.yaml
├── kitex_gen
│   └── api
│       ├── echo.go
│       ├── echo_validator.go
│       ├── echoservice
│       │   ├── client.go
│       │   ├── echoservice.go
│       │   ├── invoker.go
│       │   └── server.go
│       ├── k-consts.go
│       └── k-echo.go
├── main.go
└── middleware
    └── middleware.go
```

!!! note

    The `cwgo` command-line tool is a wrapper for the `kitex` command-line tool.
    It is used to generate the scaffolding code for the Kitex RPC servers using
    the a custom template for the API Gateway.

The `kitex` tool will generate the following `handler.go` file:

```go
// EchoServiceImpl implements the last service interface defined in the IDL.
type EchoServiceImpl struct{}

// Echo implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) Echo(ctx context.Context, request *api.EchoReq) (resp *api.EchoResp, err error) {
	// TODO: Your code here...
	return
}
```

### Updating handlers

In the `handler.go` file, fill in the logic for the Echo method:

```go
// EchoServiceImpl implements the last service interface defined in the IDL.
type EchoServiceImpl struct{}

// Echo implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) Echo(ctx context.Context, request *api.EchoReq) (resp *api.EchoResp, err error) {
	resp = &api.EchoResp{
		Response: request.Message,
	}
	return
}
```

## Starting the API Gateway

To quickly get the API Gateway up and running, we recommend using Docker, which streamlines the deployment process.

Copy and paste the following `docker-compose.yml` file and start the API Gateway:

```yaml
version: "3.9" # Docker version
services:
  # etcd is used for service registration and discovery
  etcd:
    image: quay.io/coreos/etcd:v3.5.0
    command:
      [
        "etcd",
        "--advertise-client-urls",
        "http://etcd:2379",
        "--listen-client-urls",
        "http://0.0.0.0:2379",
      ]
    ports:
      - "2379:2379"
  # http-server is the API Gateway which is pulled as a Docker image from Docker Hub
  http-server:
    image: czsheng/cloudwego-api-gateway
    volumes:
      - idl/:/etc/idl
      - gateway-logs:/logs
    environment:
      - ETCD_ADDR=http://etcd:2379
      - LOGFILE=/logs/http-server.log
    # Port forwarding configuration for API Gateway
    ports:
      - "8080:8080"
    depends_on:
      - etcd
  # Code for RPC server
  rpc-server:
    build: rpc-server
    environment:
      - ETCD_URL=http://etcd:2379
      - CONTAINER=rpc-server
      - LOGFILE=/logs/rpc-server.log
    volumes:
      - gateway-logs:/logs
    ports:
      - "8888:8888"
    depends_on:
      - etcd
# Save logs to a persistent volume
volumes:
  gateway-logs:
```

### Configuration

To customise the configuration, see the [Configuration](configuration.md) page.

## Testing the API Gateway

To test your API Gateway, you can use various HTTP client tools such as Insomnia, Postman, or Hurl.
Follow the steps below to test the API Gateway with a simple RPC server.

### Step 1: Sending a GET Request

1. Start the API Gateway and the RPC server by running `docker-compose up`.

2. Open your preferred HTTP client tool (e.g., Insomnia, Postman, or Hurl).

3. Set the method to `GET`.

4. Enter the following URL as the request endpoint: `http://localhost:8080/EchoService/echo`.

### Step 2: JSON Request Body

5. In the request body section of your HTTP client, provide the following JSON body:

```json
{
  "message": "Hello!"
}
```

### Step 3: Receiving the Response

6. Send the GET request by clicking the "Send" or "Run" button in your HTTP client.

7. You should receive the following JSON response:

```json
{
  "response": "Hello!"
}
```

Congratulations! You have successfully set up the CloudWeGo API Gateway with a simple RPC server,
and the test confirms that it is functioning correctly.
If you encounter any issues or have questions, refer to the documentation or seek help from the community. Happy building!
