# Overview

This document aims to capture details about how Kyma eventing can adopt Knative eventing. 

It attempts to answer:

1. How can a typical Kyma installation with multiple messaging solutions look like?
2. How can Knative eventing objects be mapped to the Kyma eventing model?
3. What will be the flow of publishing and consuming events?


A typical setup with Kyma eventing post-Knative adoption is as follows:

![](assets/overview.svg)

The above setup demonstrates the following aspects:

* Kyma running with 2 cluster channel provisioners backed by 2 implementations (NATS Streaming and a Cloud PubSub chosen by the customer).
* Some events such as `order.created` and `payment.received` are configured to use Cloud PubSUb.
* Other events, such as `item.viewed` and `item.compared` are configured to use NATS Streaming.
* Each event type has a channel linked in the Knative eventing which is backed by the PubSub implementation.
* The flow of the `order.created` and `payment.received` events an external solution sends, and their delivery to a respective serverless compute.


# Provisioning a messaging solution
The provisioning of various messaging solutions such as NATS Streaming, Google PubSub, Kafka etc. is abstracted from the production and consumption of events in Kyma.

You can provision the messaging solutions during installation or while Kyma is running.


You can provision messaging system using:

* Helm Charts
* Service Catalog
* Plain K8S deployments

In the long run, we can build utils that automate the provisioning process. For now, we can start with a simple solution.

A typical provisioner looks as follows:

![](assets/provisiner.svg)

> **TBD** The deprovisioning flow. The impact on the existing event types needs to be understood first.

# Kyma event types
Before Knative adoption, there was no need to create any metadata such as topics or channels in the Kyma Event-Bus as the event types were mapped to NATS Streaming subjects while publishing an event.

With Knative adoption, this model can no longer be applied due to the following:

* Knative eventing requires to create channels as a pre-step to publish. The channels are heavy objects with cascading resources. 
*The Kyma eventing solution needs to be generic and in-sync with Knative eventing to enable plugging in various messaging solutions.

## Option 1
* `1 Kyma event-type` is mapped to a `1 Knative channel`.

**Pros**

* Simple approach.
* Aligned with Knative eventing concepts and underlying bus implementations.
  ![](assets/mapping-refer.svg)
* Some reference implementations such as Kafka are mapping one `Topic` to a `Knative Channel`
* **This enables us to build a thin abstraction layer on top of Knative**.
  * Kyma does not want to run heavy workloads on the cluster when the customer is using `Cloud PubSub`.
* The Knative subscription object is mapped to channel where a single channel can have many subscriptions.
  * Following this approach we can keep the subscription model simple.

**Cons**

* A Knative channel is a heavy object with many cascading resources. Having many channels can increase the load on Istio Service Mesh. 

## Option 2

* Map multiple `Kyma event-types` to a `single channel` or do some kind of grouping.

**Pros**

* Create a minimal number of Knative channels thus creating less cascading resources.

**Cons**

* The extra complexity of maintaining the mapping.
* **Kyma eventing will be a thick layer with potential heavy workloads** running on cluster despite customer using the `Cloud PubSub`.
* The subscription management will become complex.
  * A single Knative channel will have subscriptions for multiple event types, which will generate unnecessary network traffic and workload.

  * The `dispatcher` will receive events for all event types belonging to a channel. Then it has to discard and only deliver those for which the Kyma Subscription has been created.
  * This will lead to fat topics in the underlying PubSub.

## Option 3 (Future alternative)

There have been some discussions in the Knative community about the challenges and how to map business event types to a channel.
One idea is to create **one channel per PubSub**. This will create one `K8S Service` and `Istio Virtual Service`. 
To make the dispatcher aware of the multiple diverse event types, use the **URI** part of the K8S service mapped to a channel.

![](assets/event-types-future-option3.svg)

* Knative eventing should ensure that a unified approach is used across dispatchers.
* The `Knative subscription <--> Knative Channel` model needs to evolve to support such semantics. Currently all subscriptions belong to a channel.
* There might be impact on the `Knative source` model.

## Decision

The event types in Kyma will map to Channels in Knative eventing. (Option 1).

A detailed mapping would look like below:

![](assets/event-types.svg)

## Choosing an implementation for event type

For each event type, there needs to be a pre-configuration to decide which implementation should be used. 
For events which are not critical, one can decide `local PubSub` based on NATS Streaming, while for critical events, one can decide to use battle-tested services such as `Google PubSub`.

This can be achieved either:

* during event registration,
* as separate action from the user in the Event catalog.

## When to create a channel

* Create or get a channel when a subscription is created
  * Only create the channel when someone actually wants to consume events.
  * This can be extended in a way that we discard the events at an early stage when there is no consumer configured.
  * The **missing piece of puzzle** is:
	* how/when to enable the user to specify which PubSub to use?
	* How to handle the use case when the event needs to trigger a compute in the cloud such as GCP via publishing events. In this case, there will not be any Kyma subscription. *Perhaps a operator action to manually create a channel could be applied*

**Known Challenges**

* K8S and Istio service explosion. Imagine having 1000 event types thus creating 1000 services. This will put the load on the service discovery.
  >**Note**: This has been discussed with Knative community and they are aware of this issue.

# Publishing and consumption

Knative eventing interfaces need to be abstracted to:

* Accommodate the flux in the Knative eventing APIs.
* Provide an event publish and subscribe model that is consistent with Kyma concepts.
  * Publish
  
   | Kyma                                                             | Knative                                                                                                                              |
   |------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------|
   | A well defined API with a consistent URL                         | No defined API and URL is dynamic and based on channel                                                                               |
   | `http://core-publish.kyma-system:8080/v1/events`                 | `http://order.created.ec-prod-channel.kyma-system`. This is an evolving concept and might change in the future.                      |
   | Errors are handled and translated to simplified user infomation. | Errors such as `No channel` will need to be inferred.                                                                                |
   | User always deals with the `event types` in Kyma.                | This puts the burden on the user to understand the channel aspect and translate from the event type.                                 |
   | A consistent experience for external and internal events.        | Even if we do the tranlation at the gateway for events from external, the internal events would still need to deal with Knative APIs |
   | API ensures that the event paylaod is a valid json               | No validation of event payload being a json. This can lead to difficult to troubleshoot errors especially during consumption.        |

   - Eventually Publish should evolve into a meagre proxy that is only mapping some headers or performing an http rewrite. The first step in this direction is to evolve the API to align with CE specification. 

  * Consume

	- There is not much difference for consume apart from the extra knowledge of channels and translating them to Kyma concpets of `event types` and `application identifier`.
	- Consume needs to implement `Event Activation` as Knative eventing has no such or similiar concept. Post discussion with Knative community, they do not want to introduce such constraints and expect applications to build them.
Refer [examples](./kyma-knative-eventing-examples.md)

## Publish

### Evolving the API

As per the current API, the event metadata and payload are part of the HTTP body. This was in sync with the initial cloud event specification.
The CE specification has expanded and now supports HTTP transport binding with event metadata being passed in headers. See [example](https://github.com/cloudevents/spec/blob/v0.2/http-transport-binding.md#314-examples).

The publish API will be evolved to support cloud events specification.
This could enable us in future to remove the translation being done and directly call the Knative API to publish an event.

![](assets/publish-api.svg)

### Consumption

A serverless deployed and running in Kyma should be configure events as triggers. 
For example, a Lambda configured with the trigger for the `order.created` event.

![](assets/consume.svg)


> **TBD** How can Knative sources be applied to Kyma eventing model?
