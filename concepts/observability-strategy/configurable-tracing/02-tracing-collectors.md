# Comparison of collector technologies

## Requirements

The agent's responsibility is to collect all span data in the cluster (Istio, Serverless, Eventing, custom) and ship them in a configurable way to targets. The agent is not meant to be used as backend system, so ideally it is even stateless.

- Collection
  - Support OTLP push approach, ideally with service discovery. A lot of apps are based on the pull approach and should not be rewritten or operated with a sidecar only for transforming that approach.
  - Support push-based model using OTLP.
  - Scale-out must be possible dependent on incoming traffic
- Filter
  - Support filtering traces/spans per pipeline based on the following criteria:
     - Source by Namespace or Pod name and Kubernetes labels
  - Simple tail-based sampling based on percentage, potentially also by other criterias
- Output
  - Independent outputs (with buffer or retry management)
  - OTLP support

## Alternatives

### Jaeger Collector

The Jaeger Collector is providing support for several input like OTLP, however, the outputs are related to backends support by Jaeger like Cassandra. It is not meant to be used to integrate other systems via OTLP.

### Grafana Agent

License is problematic
### OpenCensus Collector

OpenCensus will be deprecated, OpenTelemetry should be used instead
## Decision
Because the general observability strategy is based on OTLP, support for that as input and output is mandatory. Evaluation of  potential instability of the OpenTelemetry Collector shows that the big contributors are using it in production. Because the trace protocol in OpenTelemetry Collector is stable already, it is just a matter of time till all aspects are marked as stable. Because we plan to bring our own abstraction for supported scenarios, a beta API used underneath is acceptable.
With that, the OpenTelemetry Collector will be used in the further concept.
