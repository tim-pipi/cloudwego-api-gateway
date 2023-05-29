# cloudwego-api-gateway

This repository contains the code for the implementation of an API Gateway for Orbital 2023.

## Directory
IDL files are stored in `/idl`.

## Hertz
To create the Hertz server based on the Thrift IDL, run the following commands:

```shell
$ hz new -module "github.com/tim-pipi/cloudwego-api-gateway" -idl idl/hello.thrift
$ go mod tidy
```

To update the project:

```shell
hz update -idl idl/hello.thrift
```