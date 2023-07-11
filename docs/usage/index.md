# Usage

## Adding Services

Store your IDL file in the `/idl` directory.
Ensure that your IDL file follows the [Thrift IDL Annotation Standard](https://www.cloudwego.io/docs/kitex/tutorials/advanced-feature/generic-call/thrift_idl_annotation_standards/).

Notes:

- Method name in IDL file is **case-sensitive**.
- Require a type alias as the request and return type.

```thrift
struct EchoReq {
    1:required string message
}

struct EchoResp {
    1: string response
}

service EchoService {
    EchoResp echo(EchoReq) (api.get="/EchoService/echo")
}
```

### Hertz

Navigate to the `http-server` directory and generate the Hertz scaffolding code with the `hz new` command:

```shell
hz new -module "github.com/tim-pipi/cloudwego-api-gateway/http-server" -idl ../idl/[YOUR_IDL_FILE].thrift
go mod tidy
```

Update the logic in `biz/handler/api/[YOUR_IDL_FILE].go` (make the Remote Procedure Call).

### Kitex

Navigate to the `rpc-server` directory and generate the Kitex server scaffolding code with the `kitex` command:

```shell
kitex -module "github.com/tim-pipi/cloudwego-api-gateway/rpc-server" -service hello ../idl/[YOUR_IDL_FILE].thrift
```

Notes:

- The `-service` flag generates the scaffold code for creating a new client and
  server in the `rpc-server` directory.
- `-module` flag generates the `kitex_gen` directory

## Generating From Template

To generate the RPC Server scaffolding code from template, run the following command:

```shell
$ mkdir NEW_DIRECTORY
$ cd NEW_DIRECTORY
$ kitex -module "github.com/tim-pipi/cloudwego-api-gateway/NEW_DIRECTORY" --template-dir
 ../templates --thrift-plugin validator ../idl/hello_api.thrift
go: creating new go.mod: module github.com/tim-pipi/cloudwego-api-gateway/test
Adding apache/thrift@v0.13.0 to go.mod for generated code .......... Done
$ go mod tidy
```

Fill in the handler logic in `handler.go`.

## Updating Services

Run `./update.sh` in the root directory.

If you would like to manually update,

### Updating Hertz

To update the code after changes in the IDL:

```shell
hz update -idl ../idl/[YOUR_IDL_FILE].thrift
```

**Updating Behaviour**:

- No Custom Path:
  - Appends any new code to the **existing file**.
    - If you rename a method, the old method's code remains in the file.
  - Easier to handle
  - Might create duplicated code
- Custom Path
  - Guaranteed "clean code"
  - Reimplement handler logic each time
  - Confusing to keep track of directories after a while

Update the logic in `handler.go`.

### Service Registration and Discovery

Service registration and discovery is done using [etcd](https://etcd.io/docs/v3.5/)
and the [`registry-etcd`](https://github.com/kitex-contrib/registry-etcd) library.
