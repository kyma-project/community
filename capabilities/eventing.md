---
displayName: Eventing
epicsLabels:
  - area/eventing
  
---

## Scope

* Enable asynchronous processing
* Integrate external solutions with Kyma using an event-driven architecture.
* Enable loose coupling of applications deployed on Kyma.


## Vision

### Pluggability
* Enable customers to use a underlying messaging implementation that best fit their requirements.

![](assets/pluggable.png)

**Value Addition**
* Decouple workload from the management plane
* Flexiblity
  * Enable customers to choose the solution when it comes to volume, latency, scalability needs
  * Allows a customer to use the solution already existing in her ecosystem.

### System metrics

**Health metrics**

* Availability of event bus components
* Connectivity to the messaging solution
* Active / Inactive subscriptions

**Performance metrics**

* Request Latency
* Request Rate
* Failure Rate
* Delivery failure rate
* Throughput
* End-to-end latency

**Value Addition**
* Improved operations.

### Benchmarking
* Implement a mechanism to benchmark eventing
* Factors such as number of event types, fan-out, payload size configurable

**Value Addition**
* Transparency
* Increased trust in the solution

### Event Trigger for Services
* Provide a user interface for creating event triggers for sevices deployed in Kyma.

**Value Addition**
* Improved developer experience.

### Event Filtering
* Filter out events for which there are no subscribers
* Filter out close to the source

**Value Addition**
* Cost reduction
* Reduce unnecessary network overhead

### Model for internal events

* As a developer, I should be able to define and generate internal events to trigger business workflows.

**Value Addition**

* Enable asynchronous processing and workflows for:
  * Exposed APIs
  * Scheduled activities

**Dependencies**

It is possible to generate events inside a Kyma cluster and use them to trigger business workflows. 

The eventing in Kyma is tightly coupled with the aspect of events being sent by an external solution. e.g. `source-id` which identifies the external solution is used as a part of publishing an event in Kyma.

For events generated inside Kyma, the following open questions or desing details needs to be answered:

* How the orign of the internal event is identified? (events originated from the Kyma environments)
* How the event schema will be published?
* How the event schema will be made available to the developers?
* Do we need concepts such as event activation for internal events?

### Configurable durability
* The event producer should have possiblity to decide about the durability while sending an event to Kyma.
* When sending an event to Kyma event-bus, a developer can specify the durability level.

**Value Addition**
* Improved flexibility.
* Enable customers to choose performance vs durability based on the business requirements.

### Batching
* Enable support to send multiple events in a single request.

**Value Addition**
* Performance optimization.

### Retry back-off
* Enable the subscriber to configure the back-off that should be applied when retrying delivery of events.
* The backoff strategies could evolve from a constant pause to a exponential backoff.

**Value Addition**
* Efficient resource usage.
* Improved stability.

### Dead letter
* Support semantics where a message is moved to dead letter queue after failure to being processed by a lambda or a service.
* The decision parameters to move an event to a dead letter are configurable. e.g. After x number of event delivery attempts.

**Value Addition**
* Ensure business processing recovery in case of mismatches.
* Unblock delivery of events in case of processing errors. 

### Event attributes + Selective filtering

* It should be possible to assign attributes to a event published to Kyma event-bus.
* Event attributes are `key:value` pairs of metatdata that the event publisher can set while publishing the events.

**Value Addition**

* Increased flexbility for event consumption.
* Enable intelligent filtering.
* Enable efficient resource usage.
