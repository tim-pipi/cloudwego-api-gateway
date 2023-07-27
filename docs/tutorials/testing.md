# Testing

## API Testing

Validate your API endpoints to ensure they respond as expected.
You can use either [Postman](https://www.postman.com/downloads/) or [Insomnia](https://insomnia.rest/download) for manual testing.

### Hurl

For automated API testing, utilise [Hurl](https://hurl.dev/docs/installation.html) by executing the following command:

```shell
$ cd tests
$ hurl --test hello_service_test.hurl
hello_service_test.hurl: Running [1/1]
hello_service_test.hurl: Success (6 request(s) in 12 ms)
--------------------------------------------------------------------------------
Executed files:  1
Succeeded files: 1 (100.0%)
Failed files:    0 (0.0%)
Duration:        15 ms
```

## Load Tests

Ensure your API Gateway's performance under high traffic by conducting load tests.
First, install [K6](https://k6.io/docs/get-started/installation/) and then, run the load tests using the following command:

```shell
$ cd tests
$ k6 run load.js

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: load.js
     output: -

  scenarios: (100.00%) 1 scenario, 100 max VUs, 1m0s max duration (incl. graceful stop):
           * default: 100 looping VUs for 30s (gracefulStop: 30s)


running (0m03.6s), 100/100 VUs, 300 complete and 0 interrupted iterations
default   [===>----------------------------------] 100 VUs  03.6s/30s
```

## Integration Tests

Automate your integration tests using the workflows provided in this repository.

Here's a sample integration test script that you can add to your project:

```bash
#!/bin/bash
set -eu

wait_for_url() {
	echo "Testing $1..."
	echo -e "GET $1\nHTTP 200" | hurl --retry $2 >/dev/null
	return 0
}

echo "Starting Integration Test"
docker-compose -f ./examples/hello/docker-compose.yml up -d
docker-compose -f ./examples/hello/docker-compose.yml logs
while ! docker network inspect hello_default >/dev/null 2>&1; do sleep 1; done

echo "Waiting for Docker instance to be ready..."
wait_for_url 'http://localhost:8080/ping' 60
sleep 5

echo "Running Hurl tests"
hurl --test tests/hello/*.hurl

echo "Stopping Docker container"
docker-compose -f ./examples/hello/docker-compose.yml down

```

Additionally, add the following workflow to your GitHub Actions configuration to run the integration tests on push and pull request events targeting the default branch and the `testing` branch:

```yml
name: ci
on:
  push:
    branches: [$default-branch, testing]
  pull_request:
    branches: [$default-branch, testing]
permissions:
  contents: write
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: "1.20"

      - name: Build
        run: go build -v ./...

      - name: Unit Test
        run: go test -v ./...

      - name: Integration Test
        run: |
          # Install Hurl
          curl --location --remote-name https://github.com/Orange-OpenSource/hurl/releases/download/4.0.0/hurl_4.0.0_amd64.deb
          sudo dpkg -i hurl_4.0.0_amd64.deb
          tests/integration.sh
```

With these integration tests and the CI workflow, you can ensure that your CloudWeGo API Gateway functions correctly and reliably, catching any potential issues early in the development process.
