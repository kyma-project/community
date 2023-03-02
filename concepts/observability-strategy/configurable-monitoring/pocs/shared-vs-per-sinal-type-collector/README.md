# Shared OpenTelemetry Collector vs Individual Collectors Per Signal Type 

## Goal 
The [OpenTelemetry collector](https://opentelemetry.io/docs/collector/) is a core component of Kyma Telemetry. It receives different telemetry signals (tracing, metrics, and, in the future, logs), filters and augments them and sends them to the respective backends. The goal of the PoC is to determine the pros and cons of having a single shared collector with different pipelines per signal type versus having an individual collector per signal type.


## Ideas

### Scaling

The decision of whether to share the collector for all signal types or to keep them separate comes down to the question of scaling. In practice, each type may have different scaling needs and require different scaling strategies:

* Since we don't yet know enough about our users' telemetry data trends, having separate collectors would give us more flexibility in the future.
* Push and Pull Gateways are scaled differently: The tracing collector will most likely remain a push gateway (stateless) that receives, filters, augments, and pushes race data to the configured backend. Such a gateway can easily scale horizontally by adding more replicas and balancing the load between them. Using Prometheus, the metrics collector is a pull gateway (stateful) that scrapes the metrics from the configured targets. Adding more replicas would result in that all replicas scrape the the same targets. The scrapers do not scale horizontally and should be sharded. See more about [how to scale an OpenTelemetry collector] (https://opentelemetry.io/docs/collector/scaling/). 
* Workloads may generate different amounts of signals of different types. In this case, having separate collectors allows us to scale only what is needed.

### Implementation

When implementing Kubernetes operators, it's common practice to have a single controller per custom resource. The Telemetry Manager is no exception. It has a `LogPipelineController` and a `TracePipelineController`. There is also a `LogParserController`, but it's not a standalone thing and is more of an extension of the `LogPipelineController`. 

Each controller reconsiles it's own custom resource, translates it to a respective configuration (FluentBit or OpenTelemetry Collector), and deploys the resources. Imagine we implement a new `MetricPipelineController`, so that it shares the deployed OpenTelemetry Collector with the `TracePipelineController`. This would have some major drawbacks:
Both controllers will have to share at least the configuration rendering code that contains both metric and tracing pipelines. If there is a bug in the configuration rendering logic, it will affect both metrics and traces
* Both controllers will be reconciling the same resources (restarting the controller pod, etc.), possibly interfering with each other

As we can see, having a shared collector also presents certain implementation challenges.

### Single Endpoint

## Proposal
