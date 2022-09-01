# Proof of Concepts

As foundation for the final concept, see the following collection of investigations:
## Will all system components work with w3c-tracecontext?

The goal is to use w3c-tracecontext for propagation only. For that, it must be verified that this is feasible. Are there any components not yet supporting it, which will require to run both protocols (zipkin-B3) in parallel?

Check
- Eventing/SAP Eventing solution
- Istio
- Serverless

The [w3c-tracecontext](./pocs/w3c-tracecontext/README.md) proves that Kyma Serverless and Eventing supports w3c-tracecontext natively. Istio can be enabled for it with the openCensusAgent tracer.

## Head-based sampling always on

The user must be able to control how much trace data is streamed into the backend, because that has a huge impact on costs (mainly data transfer and storage). 
An option could be to make the [Head-based sampling](https://uptrace.dev/opentelemetry/sampling.html#rate-limiting-sampling) configurable in every client that pushes trace data, mainly Istio. However, the potential for making good sampling decisions is very limited, because most attributes like errors or latencies are not known at that time.
So a better approach seems to be a tail-based sampling at a central place like the planned TracePipeline configuration.
To support sampling centrally with the pipeline in the collector config, head-based sampling must always be enabled, so that the decision can be delayed to the collector.

For that, it is crucial to understand the impact of a permanently enabled sampling strategy on the used infrastructure, mainly Istio and Kubernetes.

The [Istio trace sampling rate analysis](https://github.com/kyma-project/kyma/issues/15304) investigation does that in the following way:

Running a Istio performance test with high load, and compare the following settings:
- 1% sampling to otel-collector with noop exporter
- 100% sampling to otel-collector with noop exporter
- 100% sampling to kubernetes service without endpoints
- 100% sampling to non-existing URL
Comparing the resource consumption and throughput of the envoys, and checking other relevant Kubernetes components for suspicious effects like CoreDNS.

## Tail Sampling

Understand what tail sampling strategies are available already in OTEL and understand the scaling requirements of them.
## Plugin mechanism

If the Telemetry component is disabled, there must be a way to bring your own otel-collector. How to achieve that with the planned push approach? The client's apps like Istio are preconfigured with a hardcoded URL to push traces. By default, trace data should be served by the telemetry otel-collector, but should be also servable by the custom otel-collector stack.

## AutoScaler

An otel-collector deployment must be scaled out dependent on the work it has to do. Memory and CPU are not a sufficient KPI for scaling, because the memory should be stable and the CPU is too flaky. Check which metrics are available in the otel-collector itself that indicate the amount and size of incoming OTLP messages. Because a tool like KEDA is not available yet, check how autoscaling based on these metrics can be implemented with the telemetry-operator.

## How to name the service

What should be the name for the OTLP push URL for traces? Should we use one for all signal types or dedicated ones?

## What are meaningful ways to filter traces

Is it relevant to include/exclude trace data by Namespace or container? (Results in incomplete traces?)
Is it relevant to include/exclude trace data by attributes? (Attributes might differ in the spans for one trace? Results in incomplete traces)

## Can istio/envoy report spans via OTLP already, what is with w3c-tracecontext support?

At the moment, there is no way to let Istio send trace data to a backends in OTLP protocol. The [envoy-otel](https://github.com/envoyproxy/envoy/issues/9958) integration made very good progress already and support will be provided soon.

You can enable w3c-tracecontext support already via the openCensusAgent tracer. That will change the data protocol from current zipkin to openCensus. As Jaeger does not support OpenCensus protocol, an otel-collector deployment as converter in the middle is required. The [w3c-tracecontext](./pocs/w3c-tracecontext/README.md) will outline how to enable it with an E2E scenario.

## How pipeline isolation can be achieved, is it feasible at all?

The goal of the TracePipeline is to push trace data to multiple destinations using a different set of processors or sampling strategies. How can an isolation of these pipelines be achieved based on the otel-collector? If one destination is down, can the other destination be continued?
