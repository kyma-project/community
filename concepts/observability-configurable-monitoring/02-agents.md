# Agents

## Requirements

The agents responsibility is to collect all metrics in the cluster (k8s, kyma, envoy, custom) and ship them in a configurable way to targets. It is not meant to be used as backend system, so in best case t is even stateless.

- Collection
  - Support Prometheus Format and pull approach, best with service discovery. A lot of apps are based on the pull approach and should not be re-written or operated with a sidecar only for transforming that approach
  - Support for push based model using OTLP
  - Flexible deployment model so that pull-based model might be scaled (by sharding) independent of the push-based approach
- Filter
  - Way to filter out metrics per pipeline
     - based on the source by namespace/pod name and k8s labels
     - based on metrics name (with wildcard support)
     - based on metrics labels
- Output
  - Independent outputs (with buffer/retry management)
  - OTLP support
  - Prometheus support

## Alternatives

### OpenTelemetry Collector
The [otel-collector](https://opentelemetry.io/docs/collector/) is the agent of the OpenTelemetry project, which is the vendor-neutral approach of aligning all the telemetry aspects, not focussing on any backends. Besides an own protocol OTLP, APIs with backing SDKs, it also provides a collector.

Pros
- Natively supports OTLP even as internal format
- Support for prometheus format and pull approach as receiver
- Support for prometheus as output
- Flexible way of configuring pipelines
- Big emerging community
- Part of CNCF and with that safe investment from license perspective
- Outputs do have buffer/retry management
- Pluggable and extensible
- Different deployment modes
- Emerging K8S metrics community
- Native K8S support

Cons:
- A lot of aspects are already stable, however, the collector itself is not final yet from configuration perspective. Also most of the plugins are marked as stable. However, the metrics API itself is stable.

### Prometheus agent
Prometheus is the mast famous tool for K8S when it comes to metrics. It is very mature and takes care of collection and storage. When running it in [agent mode](https://prometheus.io/blog/2021/11/16/agent/) it gets turned into a pure collector being stateless.

Pros:
- Very mature
- Support for prometheus format and pull approach as receiver
- Support for prometheus as output
- Part of CNCF and with that safe investment from license perspective
- Already in use, expertise is in the project already
- Native K8S support

Cons:
- not designed to be an agent only, there might be surprises
- No support for OTLP
- No pluggable pipeline mechanisms as only remotewrite is supported as output

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
[CollectD](https://collectd.org/) and potentially the [cassandra specific k8s port](https://docs.k8ssandra.io/components/metrics-collector/) is a well known metrics collector mainly for linux systems.

Pros:
- Performant as optimized for small devices
- Pluggable input and outputs

Cons:
- No support for OTLP (neither input/output)
- No native K8S support, only by additional casandra specific elements
- No pluggable pipeline concept

### vmAgent

The [VictoriaMetrics agent](https://docs.victoriametrics.com/vmagent.html) is part of the VictoriaMetrics project and is a tiny but powerfull collector with the main use case of collecting from multiple source and pushing them to VictoriaMetrics. It is compatible with the prometheus format

Cons:
- No OTLP support
- No pluggable pipelines

## Decision
As the general observability strategy is based on OTLP, support for that as input and output is mandatory. Even considering the potential instability of the otel-collector, it shows that the big contributors are start using it in production already. As the metrics protocol is stable already, it is just a matter of time till all aspects are being marked as stable. As we plan to bring our own abstraction for supported scenarios, a beta API used underneath is acceptable.
As the otel-collector is supporting OTLP native, the otel-collector will be used in the further concept. However, FluentBit seems to become an actual alternative and should be re-evaluated as soon as metrics input via OTLP and full prometheus scraping is supported.
