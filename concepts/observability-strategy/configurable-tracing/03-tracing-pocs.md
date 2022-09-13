# Proof of Concepts

As foundation for the final concept, see the following collection of investigations:
## (Done) Will all system components work with w3c-tracecontext?

The goal is to use w3c-tracecontext for propagation only. For that, it must be verified that this is feasible. Are there any components not yet supporting it, which will require to run both protocols (zipkin-B3) in parallel?

Check
- Eventing/SAP Eventing solution
- Istio
- Serverless

The [w3c-tracecontext](./pocs/w3c-tracecontext/README.md) proves that Kyma Serverless and Eventing supports w3c-tracecontext natively. Istio can be enabled for it with the openCensusAgent tracer.

## (Done) Head-based sampling always on

The user must be able to control how much trace data is streamed into the backend, because that has a huge impact on costs (mainly data transfer and storage). 
An option could be to make the [Head-based sampling](https://uptrace.dev/opentelemetry/sampling.html#rate-limiting-sampling) configurable in every client that pushes trace data, mainly Istio. However, the potential for making good sampling decisions is very limited, because most attributes like errors or latencies are not known at that time.
So a better approach seems to be a tail-based sampling at a central place like the planned TracePipeline configuration.
To support sampling centrally with the pipeline in the collector config, head-based sampling must always be enabled, so that the decision can be delayed to the collector.

For that, it is crucial to understand the impact of a permanently enabled sampling strategy on the used infrastructure, mainly Istio and Kubernetes.

The [Istio trace sampling rate analysis](./pocs/sampling-rate/README.md) investigation does that in the following way:

Running a Istio performance test with high load, and compare the following settings:
- 1% sampling to otel-collector with noop exporter
- 100% sampling to otel-collector with noop exporter
- 100% sampling to kubernetes service without endpoints
- 100% sampling to non-existing URL
Comparing the resource consumption and throughput of the envoys, and checking other relevant Kubernetes components for suspicious effects like CoreDNS.

Conclusions:
- Client-based sampling should not be set to 100% by default, it must be a user-controlled configuration
- If there is no endpoint available to push the data, tracing should be deactivated

## Tail-based Sampling

Opposed to head-based sampling, tail-based sampling makes the decision at the end of the entire flow, when all the trace data has been gathered. This kind of sampling decision is made at the collector level.

With tail-based sampling, it's possible to create advanced rules to filter out traces based on any **span** property, including their attributes, duration etc. With tail-based sampling, we can collect data like unusually long operations and rare errors.

With the sampling decision delayed to the end of flow when all spans of a trace are available, the decision is better because it is based on all data from the trace.

However, deciding at the end of the trace means that the backed must buffer the entire trace data, which can increase storage overhead.

Choosing the right sample rate is difficult - it mostly depends on system requirements, the way services are built,  and the amount of traffic they have. For example, when a service is very noisy and receives a lot of traffic, a smaller percentage of sampling rate is better, to prevent the costs and the noise. But for an endpoint with less traffic (like a core service), high percentage is a better choice, because it it won't cost much and, most likely, is valuable.

At the first glance, tail-based sampling seems to be a better solution than head-based sampling. 

Policy-based sampling processor configuration offers wide capabilities to configure sampling according to application needs, but this capability brings its own complexity. 

At the time this document is written (Sept '22), tail-based sampling processor is still in a *beta* state and not fully tested.

For further information about tail-based sampling, see:

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

## (Done) Can istio/envoy report spans via OTLP already, what is with w3c-tracecontext support?

At the moment, there is no way to let Istio send trace data to a backends in OTLP protocol. The [envoy-otel](https://github.com/envoyproxy/envoy/issues/9958) integration made very good progress already and support will be provided soon.

Istio can be configured to use the w3c-tracecontext for trace propagation already. However, that cannot be achieved by a native Istio feature directly, but by using the `openCensusAgent tracer` instead of the `zipkin` tracer. That will change the data protocol from current zipkin to openCensus. As Jaeger does not support OpenCensus protocol, an otel-collector deployment as converter in the middle is required. The [w3c-tracecontext PoC](./pocs/w3c-tracecontext/README.md) provides such setup.

## How pipeline isolation can be achieved, is it feasible at all?

The goal of the TracePipeline is to push trace data to multiple destinations using a different set of processors or sampling strategies. How can an isolation of these pipelines be achieved based on the otel-collector? If one destination is down, can the other destination be continued?

## Modularization - three modules or one telemetry module

Benefits of one module
- Shared caches for the controllers possible
- less maintenance (3 images instead of one, think of security scanning, kube-builder updates)
- shared packages possible without spending time in dependency tree maintenance

Benefits of three modules
- There are 3 independent domains sharing a lot of common things, it sounds more natural to model them individually
- Feature selection (on/off) is natively supported (no sub-attributes with similar purposes)

## Modularization - Re-use the deployer library of the module-manager
Can the new library of the Module Manager be used already to manage the Otel Collector deployment? Is deletion supported? What artifact format should be used?

## Otel Collector base setup
Which processors and extensions does the base setup need? What are the configuration options?
