# Observability

## Setup

To enable observability and gain valuable insights into the API Gateway's behavior and performance, follow these steps.

Open a terminal and navigate to the observability directory in your project:

```shell
$ cd observability
```

Start the necessary services using Docker Compose:

```shell
$ docker-compose up
WARN[0000] The "OTELCOL_ARGS" variable is not set. Defaulting to a blank string.
[+] Running 4/0
 ✔ Container victoriametrics                    Running                                                            0.0s
 ✔ Container observability-grafana-1            Running                                                            0.0s
 ✔ Container observability-jaeger-all-in-one-1  Running                                                            0.0s
 ✔ Container observability-otel-collector-1     Running                                                            0.0s
```

## Traces

To view the traces generated by API requests, open your web browser and navigate to [http://localhost:16686](http://localhost:16686).
This URL will take you to the [Jaeger UI](https://www.jaegertracing.io/), where you can explore distributed traces and identify potential bottlenecks or latency issues.

## Metrics

To visualize and analyze various metrics, including request rates, latencies, and error rates, access the [Grafana](https://grafana.com) UI.
Open your web browser and navigate to [http://localhost:3000](http://localhost:3000).

## Add data source

Follow the `kitex-example` [guide](https://github.com/cloudwego/kitex-examples/blob/main/opentelemetry/README.md#add-datasource)
on adding a data source.

## Import dashboard

After setting up the data source, you can import the provided `dashboard.json` file to visualize the API Gateway's key metrics.
This dashboard basic overview of system performance and resource utilisation.

## Additional Resources

For further insights into the metrics provided and detailed information on observability in the CloudWeGo API Gateway,
refer to the following repositories:

- [hertz obs-opentelemetry](https://github.com/hertz-contrib/obs-opentelemetry)
- [kitex obs-opentelemetry](https://github.com/kitex-contrib/obs-opentelemetry)