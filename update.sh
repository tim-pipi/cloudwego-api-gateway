#!/usr/bin/env bash
# Updates the rpc-server and http-server with the given IDL file

cd ./rpc-server
IDL_PATH="../idl/hello_api.thrift"
MODULE_PATH="github.com/tim-pipi/cloudwego-api-gateway/rpc-server"

kitex -module "$MODULE_PATH" "$IDL_PATH"

cd ../http-server
hz update -idl "$IDL_PATH"

cd ..