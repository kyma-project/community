# OpenTelemetry SDK Usage for Metrics in Golang Applications

OpenTelemetry is a project to handle telemetry data (logs, metrics and traces) in a vendor-neutral way. It provides SDKs to instrument applications, APIs and processing tools ([Collector](https://opentelemetry.io/docs/collector/)).

This document describes our findings using the OpenTelemetry Golang SDK for metrics export in a microservice. To evaluate the functionality of the SDK, a Golang implementation for the reviews service of the [Istio Bookinfo](https://istio.io/latest/docs/examples/bookinfo/) example was extended.

## Metric Types

OpenTelemetry supports metrics of type `counter` (increasing value) , `observer` (gauge value), and `measure`. Measures can be combined using aggregations. A possible aggregation is a histogram. The metric types are represented by the `Int64Counter`, `Int64UpDownCounter`, and `Int64ValueRecorder` types in the Golang SDK.

## Code Instrumentation

To instrument a Golang application, first, a `Meter` object must be created. To use the use the Prometheus protocol, it has to be exposed as an http handler:

```
exporter, _ := prometheus.InstallNewPipeline(prometheus.Config{})
http.HandleFunc("/", exporter.ServeHTTP)
meter := global.Meter("my-service")
```

The exporter initialization is identical to the [tracing instrumentation](instrumentation-tracing.md) for the OTLP protocol.

Individual code snippets can be instrumented with metrics by first creating a metric:

```
requestCounter := metric.Must(meter).NewInt64Counter("http.server.requests.total", metric.WithDescription("measures total the number of HTTP requests"))
```

And modified for example in a HTTP request handler:

```
h.requestCounter.Add(r.Context(), 1)
```

### Synchronous and Asynchronous Instrumentation

The OpenTelemetry [API specification](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/metrics/api.md) describes both [synchronous](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/metrics/api.md#synchronous-instrument-details) and [asynchronous](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/metrics/api.md#asynchronous-instrument-details) instruments. Both ways allow to capture single metrics or batches. While the synchronous instrumentation has already been shown in the previous examples, the asynchronous instrumentation creates a callback function that can be bound with labels afterwards.

## Findings

* The metrics API is declared to be unstable and has no documentation yet.
* Metric types in the Golang SDK abstract from the wire protocol. For instance `measure` vs `Int64ValueRecorder`.
* Access to histograms that are derived from a `Int64ValueRecorder` metric do via Grafana does not show any values.
* Other than the Prometheus Golang client, the OpenTelemetry SDK does not provide process and runtime metrics.
* The OpenTelemetry SDK does not include automated request handler instrumentation as Prometheus does with the `promhttp` package. This can potentially be solved by [third party projects](https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation).
* The asynchronous instrumentation is described in the API specification, but not implemented in the Golang SDK yet.

## Conclusions

OpenTelemetry provides similar capabilities as the [Prometheus Golang client library](https://github.com/prometheus/client_golang) for the supported metric types and instrumentation steps. However, the early stage and expected changes of the SDK does not make it suitable for production use.
