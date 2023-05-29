#!/bin/bash

echo "Starting Hertz server..."
cd ./http-server
go run . &

echo "Starting RPC server..."
cd ../rpc-server
go run . &
