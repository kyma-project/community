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

Opposed to Head-based sampling, Tail sampling make the decision at the end of the entire flow, wenn whole trace data already gathered. This kind of sampling decision made at the collector level.

With Tail sampling, it's possible to create advanced rules to filter out traces based on any **span** property, include their attributes, duration etc. Tail sampling will allow us to collect data like unsually long operations and rare errors.

Basically, with Tail sampling, sampling decision delayed end of flow untill all spans of a trace are available, this enables better sampling decisions based on all data from the trace.

However, make decision at the end of the trace, the backed has to buffer entire trace data, which can increase storage overhead.

Choosing right sample rate is diffucult, mostly depend on system requirement, the way services are built and amount of traffic they have. For example when a service is verry noisy and receiving a lot of traffic, rather a small percentece of sampling rate is a better decision, to avoid the costs and the noise. But when there is a endpoint with less traffic (e.g. a core service), high percentece is a better chooice since it won't cost much but most likely be valuable.

On first look, Tail sampling seems to be better solution over Head-based sampling. Policy based sampling processor configuration offers wide capabilities to configure sampling according to application needs, but this capability brings own complexity with itself. At the time, this document written, tail based sampling processor was still in a *beta* state and not fully tested.

For further information about Tail sampling can be found in following liks:

- [OpenTelemetry: head-based and tail-based sampling, rate-limiting](https://uptrace.dev/opentelemetry/sampling.html)
- [TraceState: Probability Sampling](https://opentelemetry.io/docs/reference/specification/trace/tracestate-probability-sampling/)
- [Tail Sampling Processor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/tailsamplingprocessor)
- [Probabilistic Sampling Processor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/probabilisticsamplerprocessor)


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

Istio can be configured to use the w3c-tracecontext for trace propagation already. However, that cannot be achieved by a native Istio feature directly, but by using the `openCensusAgent tracer` instead of the `zipkin` tracer. That will change the data protocol from current zipkin to openCensus. As Jaeger does not support OpenCensus protocol, an otel-collector deployment as converter in the middle is required. The [w3c-tracecontext PoC](./pocs/w3c-tracecontext/README.md) provides such setup.

## How pipeline isolation can be achieved, is it feasible at all?

The goal of the TracePipeline is to push trace data to multiple destinations using a different set of processors or sampling strategies. How can an isolation of these pipelines be achieved based on the otel-collector? If one destination is down, can the other destination be continued?
