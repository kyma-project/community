# Comparison of collector agents

## Requirements

**Vendor-neutral on input and output**
The goal of the new agent layer is to separate the telemetry data collection from the actual backend technology. It should be possible to integrate with a lot of different vendor specific backend providers. Furthermore, the way of providing the input to the collector should again be as much as vendor-neutral and open as possible, allowing different kinds of instrumentation in any programming language as you like.

**Rich ecosystem**
To support a variety of backends and also to increase the likelihood that the user knows how to integrate, the agent should provide an active and feature-rich ecosystem.

**Mature on K8S**
It should be a battle-tested component which can scale-out and is already in use for big setups. It should have built-in kubernetes support and a lightweight footprint.

## Candidates

**OTEL-Collector**
The most prominent technology nowadays is the `otel-collector` of the `OpenTelemetry` project. It is part of the CNCF and aims to cover all the three kind of data: logs, traces, metrics.
It is quite new but already has a big community behind as it is covering the community of OpenCensus and OpenTracing. With that the ecosystem is already good enough and the majority of vendor-specific backends are supported (seeing the available exporters on the contrib repo) being based on a new vendor-neutral protocol.
It is most probably the only real vendor-neutral project covering really all three aspects of logs, traces and metrics.
Bigger companies are actively planning to use it at least for the trace and metrics aspects. Here, the configuration specification is still not in a stable state. For logging, the whole feature is still in a beta stadium.
See also [more detailed here](../observability-opentelemtry/README.md)

**Logs - Fluent-Bit**
It is the traditional log collector used in Kyma already, being part of CNCF and widely adopted on K8S. It is vendor-neutral and supports a variety of generic outputs being capable of integrating with quite some backends. If a backend is not supported, you can use the fluent specific Forward protocol to connect a FluentD providing a even more richer ecosystem. It will scale-out perfect by running on every node with a minimalistic resource footprint.

**Logs - FluentD**
Written in Ruby and being more resource extensive then fluent-bit, FluentD presents a valid alternative to fluent-bit. However, it is usually recommended as the central application not running on every node but doing the central aggregation and shipment to the backend.

**Logs - rsyslog**
It also supports a variety of inputs and outputs. Here mainly the built-in kubernetes support is missing.

**Logs - Logstash**
Popular project with a rich ecosystem with built-in kubernetes support. Developed under the Elastic umbrella and a clear addiction to use ElasticSearch as a backend.

**Logs - Promtail**
Is not vendor-neutral on the output side, mainly designed for the Loki backend

**Logs - Flume**
Seems to have no active community anymore

**Traces - OpenCensus/Jaeger Agent**
Both agents are deprecated in favor of using otel-collector.

**Metrics - Prometheus**
Prometheus is a collector and backend at the same time. While you could reduce the backend settings to a minimum and mainly use the metrics forwarding to act as a collector/forwarder, the solution will still be not flexible. Only prometheus compatible backends will be supported.

**Metrics - Telegraf**
Has a broad community, supports various input/outputs and supports kubernetes
## Conclusion

Seeing the wide adoption of OpenTelemetry and the grow of the Community, seeing that it will cover all three data kind in a consistent way, seeing it fully vendor-neutral, seeing it written in a resource-efficient and scalable way, then the `otel-collector` is the way to go. It covers all requirements and brings the consistency on top.
However, only in the trace dimension it already gets classified as mature enough, see also recent [announcement](https://www.jaegertracing.io/docs/1.21/opentelemetry/) of jaeger. In the metrics and log dimension it might be good to still stay on some alternative solution, being ready to switch to `otel-collector` when ready.

For logs, the only real vendor-neutral solution which is supporting kubernetes natively is `fluent-bit` (with optional combination with `FluentD`). As Kyma has good experience with the project already, we should stay with it till the otel-collector is ready.

For metrics, telegraf sounds like a good intermediate candidate however the transition from prometheus to it and then forward to the otel-collector sounds cumbersome. Maybe we should stay for now on prometheus.