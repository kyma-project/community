# Concept

## Architecture Overview

![b](./assets/tracing-future.drawio.svg)

## Simple Auto-Scaling - Tail-based sampling

An auto-scaling should be possible from the beginning. Hereby, a simple Deployment of the otel-collector will be the best option to achieve a scaling. Only relevant scaling criterias based on incoming traffic needs to be evaluated to scale up/down the replicaset. A single Deployment of the otel-collector requires that sampling decisions are made on the data which is constant and available across all spans of a trace, mainly the traceID. Any other sampling would require to batch trace data and do that batching sticky on dedicated instances. That would require an advanced architecture which could be switched to at a later point in time. A simple sampling approach is acceptable.

## Temporary OpenCensus support

When introducing the configurable tracing, w3c-tracecontext should be used from the beginning. That requires to use the OpenCensus tracer in the istio configuration. That protocol needs to be supported in the beginning as a receiver till OTLP support is available.

## Custom base image

Not used plugins for the otel-collector should not be installed in the used image to reduce the attack vector. With that a custom bundling of the otel-collector image will be used

## Operator manages the Otel-Collector Deployment
There should be no resource consumption if there is no pipeline defined. That will be achieved by managing the otel-collector fully by the telemetry-operator. So the operator will deploy and configure and potentially also delete the Deployment

## Action Items

Following work packages are identified and will be transformed into stories:
- build a custom ote-collector image (with receiver for istio traces and tail-based sampling processor)
- operator deploys otel-collector if a pipeline is defined
- operator removes deployment if last pipeline gets deleted
- Have base settings in place for otel-collector deployment
- Configure otel-collector with tracePipeline settings
- Limit setup to a singular pipeline only (there can be only 0 or 1 pipeline to keep it simple)
- auto-scaling vs vertical scaling options
- Expose relevant otel-collector metrics for troubleshooting
- Secret rotation support
