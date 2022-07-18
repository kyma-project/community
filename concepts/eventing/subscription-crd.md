# Proposal for refactoring eventing subscription CRD filters

Related issue: https://github.com/kyma-project/kyma/issues/14552

# Current flow of Eventing CRD

Let suppose the following:
- JetStreamPrefix: `kyma`
- EventTypePrefix: `sap.kyma.custom`
- EventType: `sap.kyma.custom.noapp.order.created.v1`
    - Format of EventType:
        ```
        Prefix:         sap.kyma.custom
        Application:    commerce
        Event:          order.created
        version:        v1

        EventTypeFormat: <prefix>.<application>.<event>.<version>
        ```
- Subscription filter:
    ```
    filter:
        filters:
        - eventSource:
            property: source
            type: exact
            value: ""
        eventType:
            property: type
            type: exact
            value: sap.kyma.custom.noapp.order.created.v1
    ```

## Subscribing and Dispatching

- JetStream:
    - EC will internally prepend JetStreamPrefix to the EventType and will subscribe to `kyma.sap.kyma.custom.noapp.order.created.v1` on NATS Server.
        - Becasue our JetStream Stream is persisting all the events with `kyma.>`. If we set the Stream to persist all the events with `>` then it will also receive JetStream internal events.

- EventMesh:
    - EC will subscribe (Queue Subscription) to `sap.kyma.custom.noapp.order.created.v1` on EventMesh Server.


## Publishing

- Through EPP: User will publish events on `sap.kyma.custom.noapp.order.created.v1`.
    - JetStream: Internally EPP will append JetStreamPrefix before forwarding it to the NATS Server. 
        - EPP --> `kyma.sap.kyma.custom.noapp.order.created.v1`
    - EventMesh: 
        - EPP --> `sap.kyma.custom.noapp.order.created.v1` with correct `source`
- Externally (directly to eventing-backned):
    - JetStream: User will need to publish events on `kyma.sap.kyma.custom.noapp.order.created.v1` to NATS Server.
    - EventMesh: User will need to publish events on `sap.kyma.custom.noapp.order.created.v1` with correct `source` to EventMesh Server.


---

Namespace: default/sap.kyma/tunas-develop
Queue Name: default/sap.kyma/tunas-develop/ce/webhook/test-noappbfb57d975b10806c2867c74f82f2dad0a6d6be4f
Subscription Name: test-noappbfb57d975b10806c2867c74f82f2dad0a6d6be4f

topic format := <three-segment source w/o leading slash>/ce/<type with dots replaced by slashes>
The webhook subscription internally creates a queue whose lifecycle depends on the webhook subscription, that is, the queue is deleted when the subscription is deleted. The queue that gets created is named 

Queue format: <namespace of message client>/ce/webhook/<webhook subscription name>.

---


# SAP Event Specification limitations/validations ([reference](https://github.tools.sap/CentralEngineering/sap-event-specification))

Requirements for Event Context Attributes:
- source:
    - MUST NOT exceed 64 characters
    - MUST follow the schema `/<region>/<applicationNamespace>[/<instanceId>]`
        - For example: `/default/sap.kyma/tunas-develop`

- type: (or event type. Describes the type of the event related to the source the event originated in.)
    - MUST NOT exceed 127 characters
    - MUST follow the schema `<domainNamespace>.<businessObjectName>.<operation>.<version>`
```
SAP EventType format:        <domainNamespace>      . <businessObjectName>   . <operation> . <version>
Eventing EventType format:   <prefix>.<application> . <---------------event--------------> . <version>
Example:                     sap.kyma.custom.noapp  .         order          .   created   .    v1
```




# Event Mesh limitations/validations

- All queue names must have a namespace as a prefix, for example, <namespace>/myqueue
- Topics must consist of one or more segments.
- `<namespace>/<application-specific topic>` --> `<up to 63 characters>/<up to 87 characters>`

---

# Proposal for new flow of Eventing CRD

**Proposed Spec for the Subscription filters:**

```
filter:
    filters:
    - sourceType: connectedApp
    source: <application>    
    eventType: <event>.<version>
    eventMeshNamespace: <eventMeshNamespace|optional>    (e.g. /default/sap.kyma/tunas-develop)
    - sourceType: internal
    source: <application|optional>
    eventType: <event>.<version>
    - sourceType: external
    eventType: <full-event-type>      (e.g. sap.external.custom.exampleapp.order.created.v1)
    eventMeshNamespace: <eventMeshNamespace|optional>
```

- `sourceType`: Defines the type of the event of event source. It can be:
    - `connectedApp`: For events coming thourgh connected applications through Kyma Gateway and Event Publisher Proxy.
    - `internal`: For in-cluster eventing where events are published within the same Kyma cluster.
    - `external`: For subscribing to external events where events are published directly to the Eventing backend, not through Event Publisher Proxy. 

- `source`: Defines the source of the event originated from. For `connectedApp`, it is a required field. The value will be cleaned up from non-alphanumeric characters by Eventing Controller.

- `eventType`: Defines the event name for topic we need to subscribe for messages. 
    - For `sourceType: connectedApp`: It needs to define be have atleast two segments i.e. `<event>.<version>`.
    - For `sourceType: internal`: It needs to be have atleast two segments i.e. `<event>.<version>`.
    - For `sourceType: external`: It needs to provide complete event type including any prefixes.

- `eventMeshNamespace`: Optional field to override the `ce-source` header of the cloud event. It is only relevant for EventMesh. If not provided then a default value will be used by the Eventing Controller. 

## Cases:

Let suppose the following:
- EventType: `commerceMock.order.created.v1`
- JetStreamPrefix: `kyma`
- EventTypePrefix: `sap.kyma.custom`


### Case 1: Event sent by a connected application

```
apiVersion: eventing.kyma-project.io/v1alpha1
kind: Subscription
metadata:
  name: example-sub
  namespace: goldfish-example
spec:
  sink: http://test.goldfish-testing.svc.cluster.local
  filter:
    filters:
    - sourceType: connectedApp
      source: commerceMock
      eventType: order.created.v1
```

**JetStream:**
- Subscribing: 
    - EC will internally prepend `JetStreamPrefix` and `applicationName` to the EventType and will subscribe to NATS Server as `kyma.commerceMock.order.created.v1`.

- Publishing: 
    - EPP will internally prepend `JetStreamPrefix` and `applicationName` to the received event type before forwarding it to the NATS Server.
        - User --> `order.created.v1` --> EPP --> `kyma.commerceMock.order.created.v1` --> NATS


**EventMesh:**
- Subscribing: 
    - EC will internally prepend `EventTypePrefix` and `applicationName` to the EventType before subscribing on EventMesh Server.
        - For example: `sap.kyma.custom.commerceMock.order.created.v1`

- Publishing: 
    - EPP will internally prepend `JetStreamPrefix` and `applicationName` to the received event type before forwarding it to the EventMesh Server. The application name will be in the request Url.
        - User --> `order.created.v1` --> EPP --> `sap.kyma.custom.commerceMock.order.created.v1` --> EventMesh

### Case 2: Event sent from within the cluster

```
apiVersion: eventing.kyma-project.io/v1alpha1
kind: Subscription
metadata:
  name: example-sub
  namespace: goldfish-example
spec:
  sink: http://test.goldfish-testing.svc.cluster.local
  filter:
    filters:
    - sourceType: internal
      source: goldfishbot
      eventType: order.created.v1
```

**JetStream:**
- Subscribing: 
    - EC will internally prepend `JetStreamPrefix` and `applicationName` to the EventType and will subscribe to NATS Server as:
        - For example: `kyma.goldfishbot.order.created.v1`

- Publishing: 
    - Question: We don't know the `applicationName` in this case, so should we expect the user to prepend `applicationName` to eventType when publishing??? Or we can hard-code the `applicationName` to "internal" for in-cluster eventing.
    - EPP will internally prepend `JetStreamPrefix` to the received event type before forwarding it to the NATS Server.
        - User --> `goldfishbot.order.created.v1` --> EPP --> `kyma.goldfishbot.order.created.v1` --> EventMesh


**EventMesh:**
- Subscribing: 
    - EC will internally prepend `EventTypePrefix` and `applicationName` to the EventType before subscribing on EventMesh Server.
        - For example: `sap.kyma.custom.goldfishbot.order.created.v1`

- Publishing: 
    - EPP will internally prepend `JetStreamPrefix` and `applicationName` to the received event type before forwarding it to the EventMesh Server.
        - User --> `order.created.v1` --> EPP --> `sap.kyma.custom.goldfishbot.order.created.v1` --> EventMesh

### Case 3: Event sent directly to eventing-backned (not through EPP)

```
apiVersion: eventing.kyma-project.io/v1alpha1
kind: Subscription
metadata:
  name: example-sub
  namespace: goldfish-example
spec:
  sink: http://test.goldfish-testing.svc.cluster.local
  filter:
    filters:
    - sourceType: external
      eventType: sap.external.custom.exampleapp.order.created.v1
```

**JetStream:**
- Not Supported! Because NATS Servers are deployed inside the Kyma cluster and they are not exposed externally so users cannot directly publish events to NATS Servers.

**EventMesh:**
- Subscribing: 
    - EC will not prepend append anything EventType. It will subscribe to `sap.kyma.custom.commerceMock.order.created.v1` on EventMesh Server.
    - The EventType will not be cleaned by Eventing Controller from non-alphanumeric characters.
- Publishing: 
    - User will directly publish events on EventMesh Server as `sap.kyma.custom.commerceMock.order.created.v1`.