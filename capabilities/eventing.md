---
displayName: Eventing
epicsLabels:
  - area/eventing
  
---

## Scope

Eventing helps deliver the following business use cases:

* Extend existing applications by capturing the business happenings as events and triggering business extensions developed in Kyma.
* Integrate two or more applications via business events.
* Enable application developers to define workflows to acheive a business scenario.
* Enable integrating third party messgaing systems including Cloud PubSub solutions.

## Vision

### Pluggability
* Enable customers to use a messaging middleware that fits their requirements best.
* Enable customers to run multiple messaging implementations in parallel.

_Value added_
* Decouple the workload from the management plane.
* Flexiblity
  * Allow the customers to choose the solution when it comes to volume, latency, scalability needs.
  * Allows the customers to use the solution already existing in her ecosystem.

### System metrics
* Provide observability for the event-bus to enable smooth operations and proactive actions.
* This includes _Health metrics_ and _Performance metrics_.

_Value added_
* Improved operations.

### Benchmarking
* Provide tooling to enable customers to benchmark event-bus for their volume, fanout needs.

_Value added_
* Transparency.
* Increased trust in the solution.
* Enable operators to take corrective actions such as scaling some subcomponents.

### Align with CloudEvents specification
* Event bus can consume Events aligned with CloudEvents specification](https://github.com/cloudevents/spec).

_Value added_
* Enable interoperability.
* Enable developers to leverage CloudEvents SDKs for various languages.

### Event trigger for services
* Provide a user interface for creating Event triggers for sevices deployed using Kyma.

_Value added_
* Improved developer experience.

### Event filtering
* Allow only the Events that trigger certain business workflows, to use only what you need.

_Value added_
* Reduce costs.
* Reduce unnecessary network overhead.

### Model for internal Events
* Generate Events inside a Kyma cluster and use them to trigger business workflows. 

_Value added_
* Enable asynchronous processing and workflows for exposed APIs and scheduled activities.

### Batching
* Support sending multiple Events in a single request.

_Value added_
* Performance optimization.

### Retry backoff
* Enable the subscriber to configure the backoff applied when the system retries to deliver the Events.

_Value added_
* Efficient resource usage.
* Improved stability.

### Dead letter
* Support semantics allowing to move the message to a dead letter queue if it was not processed by a lambda or a service.

_Value added_
* Ensure business processes recovery in case of mismatches.
* Unblock delivery of Events in case of processing errors.

### Configurable durability
* The Event producer can decide about the durability while sending an Event to Kyma.

_Value added_
* Improved flexibility.
* Customers can choose either performance or durability based on business requirements.

### Event attributes and selective filtering

* Possibility of assigning attributes to an Event published to Kyma Event Bus.

_Value added_
* Increased flexibility of Event consumption.
* Enable intelligent filtering.
* Enable efficient resource usage.
