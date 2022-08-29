# Proof of Concepts

A collection of investigations are the base of the final concept. Here the investigations are documented.
## Will all system components work with w3c-tracecontext?

The goal is to use w3c-tracecontext for propagation only. For that it must be verified that this is feasible. Are there any components not supporting it yet which will require to run both protocols (zipkin-B3) in parallel?

Check
- Eventing/SAP Eventing solution
- Istio
- Serverless

## Head-based sampling always on

For the user it is important to control how much trace data gets streamed into the backend as it will have a huge impact on costs, mainly data transfer and storage. An option could be to make the [Head-based sampling](https://uptrace.dev/opentelemetry/sampling.html#rate-limiting-sampling) configurable in every client pushing trace data, mainly istio. However, here the potential for making good sampling decision is very limited (as most attributes like errors or latencies) are not known at that time. So a better approach seem to be a tail-based sampling at a central place like the planned TracePipeline configuration.
For supporting a sampling centrally via the pipeline in the collector config, it will require to have the head-based sampling turned on always (so that the decision can be delayed to the collector).

For that it is crucial to understand the impact of an always enabled sampling strategy on the used infrastructure, mainly istio and kubernetes.

The [Istio trace sampling rate analysis](https://github.com/kyma-project/kyma/issues/15304) investigation is doing that by:

Running a istio performance test with high load and compare the follwoing settings:
- 1% sampling to otel-collector with noop exporter
- 100% sampling to otel-collector with noop exporter
- 100% sampling to kubernetes service without endpoints
- 100% sampling to non-existing URL
Compare the resource consumption and throughput of the envoys and check other relevant kubernetes components for suspicious effects like CoreDNS.

## Tail Sampling

Understand what tail sampling strategies are available already in OTEL and understand the scaling requirements of them.
## Plugin mechanism

If the telemetry component is disabled, there should be a way to bring your own otel-collector. How to achieve that with the planned push-approach? The clients apps like istio are pre-configured with a hardcoded URL to push traces. By default it should be served by the telemetry otel-collector, but should be also servable by the custom otel-collector stack.

## AutoScaler

An otel-collector deployment need to be scaled-out dependent on the work it has to do. Memory and CPU are not a sufficiant KPI for scalign as the memory should be stable and the CPU is too flaky. Check what metrics are available in the otel-collector itself indicating the amount and size of incoming OTLP messages. As a tool like KEDA is not available yet, check how an autoscaling on base of these metrics can be impleented via the telemetry-operator.

## How to name the service

What should be the name for the OTLP push URL for traces? Should we use one for all signal types or dedicated ones?

## What are meaningful ways to filter traces

Is it relevant to include/exclude trace data by namespace or container? (Results in incomplete traces?)
Is it relevant to include/exclude trace data by attributes? (Attributes might differ in the spans for one trace? -> Incomplete traces)
