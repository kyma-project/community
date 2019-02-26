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

* Enable customers to plug in messaging middleware that fits their requirements best and run multiple messaging implementations in parallel. 
* Provide _Health metrics_ and _Performance metrics_ for the event-bus to enable smooth operations and proactive actions.
* Provide tooling to enable customers to benchmark event-bus for their volume, fanout needs.
* Align with [CloudEvents specification](https://github.com/cloudevents/spec).
* Provide a user interface for creating Event triggers for sevices deployed using Kyma.
* Filter events and transfer only those with existing subscriptions (triggers)
* Generate Events inside a Kyma cluster and use them to enable asynchronous processing.
* Support sending multiple Events in a single request.
* Enable the subscriber to configure the backoff applied when the Event-Bus retries to deliver the Events.
* Support semantics allowing to move the message to a dead letter queue if it was not processed by a lambda or a service.
* Enable possibility to assign event attributes and specify event durability.
