# Testing

## Unit Tests

To run the unit tests, run the following command:

```shell
$ cd http-server/internal/pkg
$ go test ./...
ok      github.com/tim-pipi/cloudwego-api-gateway/http-server/internal/pkg/config       0.003s
ok      github.com/tim-pipi/cloudwego-api-gateway/http-server/internal/pkg/service      0.003s
```

## API Testing

To test API endpoints, install [Postman](https://www.postman.com/downloads/) or [Insomnia](https://insomnia.rest/download).

### Hurl

To test API endpoints programmatically, install [Hurl](https://hurl.dev/docs/installation.html) and run the following command:

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

To run the load tests, install [K6](https://k6.io/docs/get-started/installation/) and run
the following command:

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
