# Shared OpenTelemetry Collector vs Individual Collectors Per Signal Type 

## Goal 
The [OpenTelemetry collector](https://opentelemetry.io/docs/collector/) is a core component of Kyma Telemetry. It receives different telemetry signals (tracing, metrics, and, in the future, logs), filters and augments them and sends them to the respective backends. The goal of the PoC is to determine the pros and cons of having a single shared collector with different pipelines per signal type versus having an individual collector per signal type.


## Ideas

### Scaling

The decision of whether to share the collector for all signal types or to keep them separate comes down to the question of scaling. In practice, each type may have different scaling needs and require different scaling strategies:

* Since we don't yet know enough about our users' telemetry data trends, having separate collectors would give us more flexibility in the future.
* Push and pull gateways are scaled differently. The tracing collector will most likely remain a push gateway (stateless) that receives, filters, augments, and pushes race data to the configured backend. Such a gateway can easily scale horizontally by adding more replicas and balancing the load between them. Metrics collection is a bit more complicated, as we need to combine a push gateway with an OTLP receiver (stateless) and Prometheus scrapers (statefull). The scrapers do not scale horizontally and should be sharded. See more about [how to scale an OpenTelemetry collector] (https://opentelemetry.io/docs/collector/scaling/). 
* Workloads may generate different amounts of signals of different types. In this case, having separate collectors allows us to scale only what is needed.

### Implementation

When implementing Kubernetes operators, it's common practice to have a single controller per custom resource. The Telemetry Manager is no exception. It has a `LogPipelineController` and a `TracePipelineController`. There is also a `LogParserController`, but it's not a standalone thing and is more of an extension of the `LogPipelineController`. 

Each controller reconsiles it's own custom resource, translates it to a respective configuration (FluentBit or OpenTelemetry Collector), and deploys the resources. Imagine we implement a new `MetricPipelineController`, so that it shares the deployed OpenTelemetry Collector with the `TracePipelineController`. This would have some major drawbacks:
Both controllers will have to share at least the configuration rendering code that contains both metric and tracing pipelines. If there is a bug in the configuration rendering logic, it will affect both metrics and traces
* Both controllers will be reconciling the same resources (restarting the controller pod, etc.), possibly interfering with each other

As we can see, having a shared collector also presents certain implementation challenges.

### Single Endpoint

From the usablity point of view it would be handy to provide a single endpoint to push telemetry data in the OTLP format. It's possible to do it using URL rewriting, so e.g. requests targeting "telemetry.kyma-system.svc.cluster.local/v1/metrics" are forwarded to the metrics gateway and "telemetry.kyma-system.svc.cluster.local/v1/traces"
to the tracing gateway. There are a few possibilities to implement it:
*  [Istio Virtual Service](https://istio.io/latest/docs/reference/config/networking/virtual-service/) using the default `mesh` gateway
* An HTTP proxy with a hard-coded configuration to rewrite the request URLs

The first approach is the easiest one, the only drawback is a hard dependency on Istio.

### Signal Translation

It is sometimes necessary to convert signals from one signal type to another. Signal translations may be direct one-to-one operations or include derivative signals such as counts, aggregations, summarizations, etc. It is possible to, let's say, convert traces to metrics in the trace gateway and send them to the metric gateway, which in turn would send them to the backend. However, a shared collector is advantageous in this case because it can do it in-process without performing additional network calls.

## Proposal

It becomes clear that having an individual collector per signal type has a lot of advantages and gives us flexibility to come up with a good scaling strategy later.
