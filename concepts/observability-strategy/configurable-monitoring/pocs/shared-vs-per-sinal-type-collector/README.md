# Shared OpenTelemetry Collector vs Individual Collectors Per Signal Type 

## Goal 
The [OpenTelemetry Collector](https://opentelemetry.io/docs/collector/) is a core component of Kyma Telemetry. It receives different telemetry signals (tracing, metrics, and, in the future, logs), filters and augments them and sends them to the respective backends. The goal of the PoC is to determine the pros and cons of having a single shared collector with different pipelines per signal type versus having an indvidual collector per signal type.

## Ideas

### Scaling


### Implementation

When implementing Kubernetes operators, it's common practice to have a single controller per custom resource. The Telemetry Manager is no exception. It has a `LogPipelineController` and a `TracePipelineController`. There is also a `LogParserController`, but it's not a standalone thing and is more of an extension of the `LogPipelineController`. 

Each controller reconsiles it's own custom resource, translates it to a respective configuration (FluentBit or OpenTelemetry Collector), and deploys the resources. Imagine we implement a new `MetricPipelineController`, so that it shares the deployed OpenTelemetry Collector with the `TracePipelineController`. This would have some major drawbacks:
* Both controllers will have to share at least the configuration rendering code that contains both metric and tracing pipelines. If there iss a bug in the configuration rendering logic, it will affect both metrics and traces
* Both controllers will be reconciling the same resources (restarting the controller pod, etc.), possibly interfering with each other

As we can see, having a shared collector also presents certain implementation challenges.

## Proposal
