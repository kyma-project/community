# OpenTelemetry Collector Usage for Distributed Tracing

OpenTelemetry is a project to handle telemetry data (logs, metrics and traces) in a vendor-neutral way. It provides SDKs to instrument applications, APIs and processing tools (collector).

This document describes different ways to deploy the OpenTelemetry collector to handle traces and findings from a PoC. To showcase the functionality of OpenTelemetry, we deployed the [Istio Bookinfo](https://istio.io/latest/docs/examples/bookinfo/) example and forwarded the Istio generated traces to OpenTelemetry.

## Collector Services

The collector consists of telemetry receivers, processors and exporters, which can be combined to processing pipelines in services. E.g., the following configuration can receive traces via the Zipkin and OTLP API, batch them and forwards the batches to another Zipkin service and writing to STDOUT.

```
receivers:
  zipkin:
  otlp:
    protocols:
      grpc:
      http:

processors:
  batch:

exporters:
  zipkin:
    endpoint: "http://zipkin:9411/api/v2/spans"
  logging:
    loglevel: debug

extensions:
  health_check:

service:
  extensions: [health_check]
  pipelines:
    traces:
      receivers: [otlp, zipkin]
      exporters: [zipkin, logging]
```

### Supported APIs

The collector can receive traces from different APIs. Supported are for traces

* Jaeger (grpc, thrift_binary, thrift_compact, and thrift_http)
* Kafka
* OpenCensus
* OTLP (grpc and http)
* Zipkin

Traces can be exported using the same APIs as well as writing to STDOUT and files.

### Collector Deployment

The collector can be deployed as a central service, a DaemonSet or sidecar container to each pod. Telemetry data can be forwarded by an exporter to another collector or an external service (e.g., Jaeger).

We build an OpenTelemetry collector deployment for this PoC where each pod has a sidecar collector that can receive traces via the Zipkin and OTLP API. The sidecar collector forwards the traces to a central collector. The central collector sends all traces in batches to the Zipkin endpoint of the Jaeger service. The used configuration can be found in [deployment](deployment/). To integrate this setup with Istio, one can set `defaultConfig.tracing.zipkin.address` to `localhost` in Istio's mesh configuration.

### Findings

* Traces that are forwarded to Jaeger via a OpenTelemetry collector do not have a root span.
* Receiving traces from Istio by a central collector did not work. This might be caused by some interference with the Istio sidecar.

## OpenTelemetry Collector Operator

The OpenTelemetry project provides an [operator](https://github.com/open-telemetry/opentelemetry-operator) to deploy the Collector to Kubernetes. The operator supports the deployment of collectors as single instance, DaemonSets or sidecar containers. Sidecar containers are injected to pods by a MutatingWebhook.

### Findings

* Istio sidecar injection has to be disabled for the operator to make its mutating webhook working.
* Modifications to `opentelemetrycollector` custom resources are not applied to existing deployments. Resources have to be deleted and recreated.
* Cert-manager is required to issue a certificate for the mutating webhook.
* Using a sidecar container configuration in multiple namespaces did not work.

## Conclusions

The OpenTelemetry Collector's capabilities can fulfill the vision of a pluggable observability backend for traces that is described in [kyma-project/kyma#10119](https://github.com/kyma-project/kyma/issues/10119). However, the project is in a very early stage is seems to be not mature for production use yet (see described findings).
