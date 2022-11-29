# Concept

## Architecture Overview

![b](./assets/tracing-future.drawio.svg)

## Simple Auto-Scaling - Tail-based sampling

Auto-scaling should be possible from the beginning. A simple Deployment of the Otel Collector is the best option to achieve scaling. We only have to evaluate the relevant scaling criteria based on incoming traffic to scale the ReplicaSet up and down. A single Deployment of the Otel Collector requires that sampling decisions are made on the data that is constant and available across all spans of a trace, mainly the traceID.
More advanced sampling like respecting some error status would require to batch all spans for a trace at one instance in order to make a decision. Having a scalable setup would require to introduce another StatefulSet of the Otel Collector and forward the spans for a specific trace always to the same instance of the StatefulSet (sticky). That advanced architecture can be adopted at a later time.
A simple sampling approach is acceptable for now.

## Temporary OpenCensus support

When introducing the configurable tracing, w3c-tracecontext should be used from the beginning. That requires to use the OpenCensus tracer in the Istio configuration. Till OTLP support is available, that protocol must be supported as a receiver.

## Custom base image

Unused plugins for the Otel Collector should not be installed in the used image to reduce the attack vector. With that, a custom bundling of the Otel Collector image will be used

## Operator manages the Otel Collector Deployment

There should be no resource consumption if there is no pipeline defined. To achieve this, the Telemetry Operator will fully manage the Otel Collector; meaning the operator will deploy, configure, and potentially also delete the Deployment.

### Ensuring a Single TracePipeline

The initial release of configurable tracing is designed to support only a single pipeline. While this limitation keeps the Otel Collector setup simple, the restriction has to be enforced to ensure a correct Otel Collector configuration.

We intend to utilize the optimistic concurrency control property of the Kubernetes API through creating an arbitrary API resource that acts as a lock. Even though, the user would be able to create multiple trace pipelines, only one of them must be reconciled and applied to the Otel Collector configuration. All other trace pipelines must wait in a `Pending` state until the active trace pipeline is deleted.

The following steps have to be taken while reconciling to ensure only one active trace pipeline:

* The controller's reconcile method creates the "lock" resource before performing any other reconciliation steps. This can be for instance a ConfigMap or a dedicated CRD that can hold additional state. The reconciled lock pipeline becomes owner of the "lock" resource ([owner reference](https://kubernetes.io/docs/concepts/overview/working-with-objects/owners-dependents/)).

* If the creation of the "lock" fails, we already have an active trace pipeline. We compare the owner reference and continue to reconcile if the reconciled trace pipeline is the owner. If the "lock" has a different owner, the current trace pipeline stays in `Pending` state.

* When then active trace pipeline is deleted, the "lock" resource is also deleted through the [cascading deletion](https://kubernetes.io/docs/concepts/architecture/garbage-collection/) property of the owner reference. Another trace pipeline can now take the "lock" and become the active one.

We discussed the following alternatives to this approach, which are not suitable to guarantee a single active trace pipeline:

* Rejecting the creation of a second trace pipeline with a validating webhook: Creating a second trace pipeline cannot be prevented for sure since the validation and creation are two different steps. The second pipeline could be successfully pass the validation before the first one has been created. A webhook can be used additionally to the described approach to give early feedback to the user on a best-effort base.

* Observing the existence of other trace pipelines and reconciling the current pipeline only if no other one is active: Concurrent creation of multiple trace pipelines can again bypass this check. Depending on the implementation, situations might occur where either no pipeline or multiple pipelines may become active.

* Utilizing Kubernetes' [Lease API](https://kubernetes.io/docs/reference/kubernetes-api/cluster-resources/lease-v1/): Very similar to the approach that is described above. However, a lease requires regular renewal. This causes the risk to loose the state if the operator is (temporarily) not running.

## Action Items

The following work packages are identified and will be transformed into stories:
- Build a custom Otel Collector image with a receiver for Istio traces and tail-based sampling processor.
- If a pipeline is defined, the Telemetry Operator deploys Otel Collector.
- When the last pipeline is deleted,  the Telemetry Operator removes the Deployment.
- Define basic settings for the Otel Collector Deployment.
- Configure Otel Collector with tracePipeline settings.
- Limit the setup to a singular pipeline only. To keep it simple, there can be only 0 or 1 pipeline.
- Define auto-scaling vs vertical scaling options.
- Expose relevant Otel Collector metrics for troubleshooting.
- Support Secret rotation.
