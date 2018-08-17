# ADR 018: Knative-based Kyma Eventing 

Created on 2018-08-19 by Ahmed Abdalla (@Abd4llA).

## Context

Based on the decision of building Kyma on top of Knative and along the effort to reuse as much of knative functionality as possible, the eventing area shines as a good area for alignment. That is why a decision was made to have Kyma leverage as much of knative eventing as possible.This is a an architecture decision on how to achieve that target.

## Assumptions
1. Kyma is to be based on knative.
2. Favor a smooth, gradual migration for Kyma components to Knative Eventing
3. Consider areas of contribution to knative eventing
4. Event Bus public integration: Publish RESTful API, EventActivation and Subscriptions CRDs are still needed
5. Leveraging knative eventing is a requirement

## Decision

### Kyma Event Bus 

The decisions is to provide a solution that abstracts Knative Eventing concepts from the rest of Kyma as a short to med-term solution. That way the transition to knative eventing is transparent to the rest of Kyma components.

#### Architecture 


                                +------------------+     +-----------------+
                                |                  |     |                 |
                                |  EventActivation |     |  Subscription   <---------+
                                |                  |     |                 |         |
                                +---------+--------+     +-----------+-----+         |
                                          |                          |               |
                                          |                          |               |
                +---------------+         |    +---------------+     |               |
                |               |     Read|    |               |     |Read/Update    |
                |               |         |    |               |     |               |
                |               |         |    |               |     |               |
                |    Publish    |         +---->  Events Bus   <-----+               |
                |               |              |  Controller   |                     |Create/Delete
                |               |              |               |                     |
                +--------+------+              +--+-------+----+                     |
                         |                        |       |Create/Delete             |
                         |Publish   Create/Delete |       |                          |
                         |                        |       +-----+                    |
                     +---v------------------------v----+        |             +------+-----+
                     |                                 |        |             |            |
                     |                                 +--------v-------------+  Subscriber|
                     |         KN-Channel              |    KN-Subscription   |            |
                     |                                 +----------------------+            |
                     |                                 |                      +------------+
              +------+---------------------------------+--------+
              |                                                 |
              |                                                 |
              |                 Knative Bus                     |
              |                                                 |
              +-------------------------------------------------+

### NATS Streaming

Since knative provides an abstraction of "The Bus" which is a pluggable component of knative eventing that have different  implementations backed up by a certain messaging store "e.g. Kafka Bus, GCP PubSub Bus, etc..".

### Decision

1. Implementing a production ready `NATS Streaming` knative bus implementation.
2. Provide `NATS Streaming` operator that provides `NATS Streaming` clusters which satisfies all production needs.
3. Implementing an `Azure Service Bus Messaging` knative bus implementation as a plan B.
4. Contribute both implementations to `knative eventing`.

## Status

In Progress

## Consequences

- Refactor `Event Bus Publish` application
- Merge `Event Bus Push` & `Sub Validator` apps into a single `Event Bus Controller` app.
