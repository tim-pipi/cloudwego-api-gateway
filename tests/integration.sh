#!/bin/bash
set -eu

wait_for_url() {
	echo "Testing $1..."
	echo -e "GET $1\nHTTP 200" | hurl --retry $2 >/dev/null
	return 0
}

echo "Starting Integration Test"
docker-compose -f ./examples/hello/docker-compose.yml up -d
while ! docker network inspect hello_default >/dev/null 2>&1; do sleep 1; done

echo "Waiting for Docker instance to be ready..."
wait_for_url 'http://localhost:8080/ping' 60
sleep 5

echo "Running Hurl tests"
hurl --test tests/hello/*.hurl

echo "Stopping Docker container"
docker-compose -f ./examples/hello/docker-compose.yml down
