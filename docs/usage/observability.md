# Observability

## Setup

To obtain observability data, follow these steps:

```shell
$ cd observability
$ docker-compose up
WARN[0000] The "OTELCOL_ARGS" variable is not set. Defaulting to a blank string.
[+] Running 4/0
 ✔ Container victoriametrics                    Running                                                            0.0s
 ✔ Container observability-grafana-1            Running                                                            0.0s
 ✔ Container observability-jaeger-all-in-one-1  Running                                                            0.0s
 ✔ Container observability-otel-collector-1     Running                                                            0.0s
```

## Traces

Navigate to [http://localhost:16686](http://localhost:16686) to view the Jaeger UI.

## Metrics

Navigate to [http://localhost:3000](http://localhost:3000) to view the Grafana UI.

### Add data source

TODO

### Import dashboard

Import the `dashboard.json` file to view the dashboard default dashboard.

### Additional Resources

For more information on the metrics provided see the following repositories:

- [hertz obs-opentelemetry](https://github.com/hertz-contrib/obs-opentelemetry)
- [kitex obs-opentelemetry](https://github.com/kitex-contrib/obs-opentelemetry)
