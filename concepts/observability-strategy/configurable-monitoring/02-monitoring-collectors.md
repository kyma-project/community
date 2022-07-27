# Comparison of collector technologies

## Requirements

The agent's responsibility is to collect all metrics in the cluster (Kubernetes, Kyma, envoy, custom) and ship them in a configurable way to targets. The agent is not meant to be used as backend system, so ideally it is even stateless.

- Collection
  - Support Prometheus format and pull approach, ideally with service discovery. A lot of apps are based on the pull approach and should not be rewritten or operated with a sidecar only for transforming that approach.
  - Support push-based model using OTLP.
  - Support flexible deployment model so that the pull-based model might be scaled (by sharding) independently of the push-based approach.
- Filter
  - Support filtering metrics per pipeline based on the following criteria:
     - Source by Namespace or Pod name and Kubernetes labels
     - Metrics name (with wildcard support)
     - Metrics labels
- Output
  - Independent outputs (with buffer or retry management)
  - OTLP support
  - Prometheus support

## Alternatives

### OpenTelemetry Collector
The [otel-collector](https://opentelemetry.io/docs/collector/) is the agent of the OpenTelemetry project, which is the vendor-neutral approach of aligning all the telemetry aspects, not focussing on any backends. Besides an own protocol OTLP, APIs with backing SDKs, it also provides a collector.

Pros
- Natively supports OTLP even as internal format
- Support for Prometheus format and pull approach as receiver
- Support for Prometheus as output
- Flexible way of configuring pipelines
- Big emerging community
- Part of CNCF and with that, safe investment from license perspective
- Outputs with buffer and retry management
- Pluggable and extensible
- Different deployment modes
- Emerging Kubernetes metrics community
- Native Kubernetes support

Cons:
- A lot of aspects are already stable, however, the collector itself is not final yet from a configuration perspective. Also, most of the plugins are marked as stable, and the metrics API itself is stable.

### Prometheus agent
Prometheus is the most famous tool for Kubernetes when it comes to metrics. It is very mature and takes care of collection and storage. When running it in [agent mode](https://prometheus.io/blog/2021/11/16/agent/), it turns into a pure collector being stateless.

Pros:
- Very mature
- Support for Prometheus format and pull approach as receiver
- Support for Prometheus as output
- Part of CNCF and with that, safe investment from license perspective
- Already in use at Kyma, expertise is in the project already
- Native Kubernetes support

Cons:
- Not designed to be an agent only, there might be surprises
- No support for OTLP
- No pluggable pipeline mechanisms because only remote write is supported as output

### Fluent Bit
[FluentBit](https://fluentbit.io/) is a famous log collector being part of the CNCF. It has a pluggable design with a very low resource footprint and it is used in kyma already as a collector for logs. It nowadays gets extended more and more to support metrics as well.

Pros:
- Performant and lightweight
- Pluggable input and outputs
- Used already as log collector
- OTEL output support (input to come soon)
- Basic metric inputs, even k8s based like node exporter

Cons:
- Not designed to be a metric collector
- Very initial support for prometheus input and output


### CollectD
[CollectD](https://collectd.org/) - and potentially the [cassandra-specific Kubernetes port](https://docs.k8ssandra.io/components/metrics-collector/) - is a well-known metrics collector mainly for Linux systems.

Pros:
- Performant, because it's optimized for small devices
- Pluggable input and outputs

Cons:
- No support for OTLP (neither input nor output)
- No native Kubernetes support, only with additional cassandra-specific elements
- No pluggable pipeline concept

### vmAgent

The [VictoriaMetrics agent](https://docs.victoriametrics.com/vmagent.html) is part of the VictoriaMetrics project and is a tiny but powerful collector. Its main use case is collecting metrics from multiple sources and pushing them to VictoriaMetrics. It is compatible with the Prometheus format.

Cons:
- No OTLP support
- No pluggable pipelines

## Decision
As the general observability strategy is based on OTLP, support for that as input and output is mandatory. Even considering the potential instability of the OpenTelemetry Collector, it shows that the big contributors are using it in production already. As the metrics protocol in OpenTelemetry Collector is stable already, it is just a matter of time till all aspects are marked as stable. Because we plan to bring our own abstraction for supported scenarios, a beta API used underneath is acceptable.
With that, the OpenTelemetry Collector will be used in the further concept. However, FluentBit seems to become an actual alternative and should be re-evaluated as soon as metrics input via OTLP and full prometheus scraping is supported.
