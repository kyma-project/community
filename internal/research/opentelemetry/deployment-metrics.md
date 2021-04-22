# OpenTelemetry Collector Usage for Metrics

OpenTelemetry is a project to handle telemetry data (logs, metrics and traces) in a vendor-neutral way. It provides SDKs to instrument applications, APIs and processing tools (collector).

This document describes different ways to deploy the OpenTelemetry collector to handle metrics and findings from a PoC. To demonstrate the functionality of OpenTelemetry, we deployed an application that exports metrics via the Prometheus `/metrics` HTTP path.

## Collector Services

The collector consists of telemetry receivers, processors and exporters, which can be combined to processing pipelines in services. For example, the following configuration receives metrics by scraping a Prometheus compatible HTTP endpoint. Then, it forwards the metrics to another central collector via the OTLP protocol. Scraping Prometheus metrics requires a [Prometheus scrape config](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#scrape_config) for the receiver.

```
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: "reviews"
          scrape_interval: 5s
          static_configs:
            - targets: ["localhost:9080"]

exporters:
  otlp:
    endpoint: kyma-collector.kyma-system:55680
    insecure: true

service:
  pipelines:
    metrics:
      receivers: [prometheus]
      exporters: [otlp]
```

### Supported APIs

The collector can receive metrics from the following APIs:

* Host metrics (similar to [Prometheus Node Exporter](https://github.com/prometheus/node_exporter))
* OpenCensus
* OTLP
* Prometheus

Metrics can be exported using the same APIs, the Prometheus remote write interface, and can be written to STDOUT and files.

### Collector Deployment

The collector must be deployed as a central service, a DaemonSet or sidecar container to each pod. Telemetry data can be either forwarded by an exporter to another collector or an external service (e.g., Prometheus or Cortex). The Collector can use the full set of Prometheus configuration options for scraping. For instance the Kubernetes service discovery.

We built an OpenTelemetry collector deployment for this PoC, where a pod that exports Prometheus metrics gets an OpenTelemetry Collector sidecar container injected to scrapes them. The sidecar collector forwards the traces to a central collector. The central Collector sends all metrics to a [Cortex](https://github.com/cortexproject/cortex) service via the Prometheus remote write endpoint.
For details about the configuration, see [deployment](deployment/).

### Findings

* The Prometheus receiver requires a plain scraping configuration and does not provide the convenience of the [prometheus-operator](https://github.com/prometheus-operator/prometheus-operator) with its `ServiceMonitor` and `PodMonitor` resources.
* Scraping Prometheus metrics from a remote pod does not work. This is probably caused by the Istio sidecar injection.
* The Prometheus remote write exporter does not support authentication. Writing to an endpoint that requires authentication can be solved by manually setting headers for the exporter.
