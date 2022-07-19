title: Proposal for refactoring eventing subscription CRD filters

Related issue: https://github.com/kyma-project/kyma/issues/14552

# Current flow of Eventing

The current Subscription CRD is:

```
apiVersion: eventing.kyma-project.io/v1alpha1
kind: Subscription
metadata:
  name: <string>
  namespace: <string>
spec:
  sink: <HTTP URL>
  filter:
    filters:
    - eventSource:
        property: source
        type: exact
        value: ""
      eventType:
        property: type
        type: exact
        value: <prefix>.<application>.<event>.<version>
```


Let suppose the following:
- JetStreamPrefix: `kyma`
- EventTypePrefix: `sap.kyma.custom`
- EventType: `sap.kyma.custom.commercemock.order.created.v1`, where:
    ```
    Prefix:         sap.kyma.custom
    Application:    commercemock
    Event:          order.created
    version:        v1
    ```

## Subscribing

### JetStream:

Eventing Controller will internally prepend `JetStreamPrefix` to the `EventType` (i.e. `<JetStreamPrefix>.<EventType>`) before subscribing on NATS Server.
- For example: `kyma.sap.kyma.custom.commercemock.order.created.v1`

The reason for prepending `JetStreamPrefix` to the `EventType` is that our JetStream Stream is persisting all the events with "`kyma.>`". If we set the Stream to persist all the events with "`>`" then it will also receive JetStream internal events.

### EventMesh:

Eventing Controller will subscribe to the provided `EventType` without any modification on EventMesh Server.
- For example: `sap.kyma.custom.commercemock.order.created.v1`


## Publishing

### Connected Application

User will publish events on `sap.kyma.custom.commercemock.order.created.v1`.

For legacy events, the user will publish the events on Url: `http://<host>/<applicationName>/v1/events`. For example:
```
curl -v -X POST \
    -H "Content-Type: application/json" \
    -H "X-B3-Sampled: 1" \
    --data @<(<<EOF
    {
		"event-type": "order.created",
		"event-type-version": "v1",
		"event-time": "2020-09-28T14:47:16.491Z",
		"data": {"orderCode": "987"}
  	}
EOF
    ) \
    http://localhost:8081/commercemock/v1/events
```

The legacy events are converted into cloud events by Event publisher proxy. The application name is extracted from the url and prepended to the `event-type` before publishing the event to the eventing backend. Also, the header `ce-source` in Cloud Event header is added as required by the eventing-backend.

For cloud events, the user will publish the events on Url: `http://<host>/publish`. For example:
```
curl -v -X POST \
     -H "ce-specversion: 1.0" \
     -H "ce-type: sap.kyma.custom.commercemock.order.created.v1" \
     -H "ce-source: /default/sap.kyma/tunas-develop" \
     -H "ce-eventtypeversion: v1" \
     -H "ce-id: A234-1234-1234" \
     -H "content-type: application/json" \
     -d "{\"foo\":\"binary-mode\"}" \
     http://localhost:8081/publish
```
The user needs to provide the correct `ce-source` and `ce-type`.

Eventing backend specific modifications:
- JetStream: Internally EPP will append `JetStreamPrefix` before forwarding it to the NATS Server. 
    - For example: `kyma.sap.kyma.custom.commercemock.order.created.v1`
- EventMesh: Internally EPP will append `EventTypePrefix` in case of legacy events. 
    - For example: `sap.kyma.custom.commercemock.order.created.v1`

### Internal (In-cluster) Eventing

The flow of in-cluster is same as in Connected Applications, only the difference here is that we use some non-existing application name. For example, `noapp` etc.

### Externally (directly to eventing-backned)

- **JetStream:** User will need to publish events on `kyma.sap.kyma.custom.noapp.order.created.v1` to NATS Server.

- **EventMesh:** User will need to publish events on `sap.kyma.custom.noapp.order.created.v1` with correct `source` to EventMesh Server.

# SAP Event Specifications ([reference](https://github.tools.sap/CentralEngineering/sap-event-specification))

Requirements for Event Context Attributes:
- source: (Identifies the instance the event originated in.)
    - MUST NOT exceed 64 characters.
    - MUST follow the schema `/<region>/<applicationNamespace>[/<instanceId>]`
        - For example: `/default/sap.kyma/tunas-develop`

- type: (Describes the type of the event related to the source the event originated in.)
    - MUST NOT exceed 127 characters
    - MUST follow the schema `<domainNamespace>.<businessObjectName>.<operation>.<version>`

# Event Mesh limitations/validations ([reference](https://help.sap.com/docs/SAP_EM/bf82e6b26456494cbdd197057c09979f/df532e8735eb4322b00bfc7e42f84e8d.html))

Event Mesh follows the SAP Event Specifications. For the `source` attribute, it uses the value of Event Mesh Namespace (e.g. `default/sap.kyma/tunas-develop`). Apart from that, there are some limits on queue and subscription names. For example,

- All queue names must have a namespace as a prefix, for example, <namespace>/myqueue
- Topics must consist of one or more segments.
    - `<namespace>/<application-specific topic>` --> `<up to 63 characters>/<up to 87 characters>`

If you create a webhook subscription then EventMesh internally creates a queue whose lifecycle depends on the webhook subscription, that is, the queue is deleted when the subscription is deleted. The queue that gets created is named as: 

- Queue name: `<namespace of message client>/ce/webhook/<webhook subscription name>`
- Example Queue name: `default/sap.kyma/tunas-develop/ce/webhook/test-noappbfb57d975b10806c2867c74f82f2dad0a6d6be4f`


# Mapping between Event type formats

```
SAP EventType format:        <domainNamespace>      . <businessObjectName>   . <operation> . <version>
Eventing EventType format:   <prefix>.<application> . <---------------event--------------> . <version>
Example:                     sap.kyma.custom.noapp  .         order          .   created   .    v1
```


# Proposal for new flow of Eventing CRD

Following is the proposed CRD for Kyma Subscription:

```
apiVersion: eventing.kyma-project.io/v1alpha1
kind: Subscription
metadata:
  name: <string>
  namespace: <string>
spec:
  sink: <HTTP URL>
  events:
    - sourceType: connectedApp
      source: <application>    
      type: <event>.<version>   
    - sourceType: internal
      source: <name>
      type: <event>.<version>
    - sourceType: external
      source: <string>             (e.g. /default/sap.kyma/tunas-develop)
      type: <full-event-type>      (e.g. sap.external.custom.exampleapp.order.created.v1)
```

Explanation of fields:

- `sourceType`: Defines the type of the event of event source. It can have the following values:
    - `connectedApp`: For events coming thourgh connected applications through Kyma Gateway and Event Publisher Proxy.
    - `internal`: For in-cluster eventing where events are published within the same Kyma cluster.
    - `external`: For subscribing to external events where events are published directly to the Eventing backend, not through Event Publisher Proxy.

- `source`: Defines the source of the event originated from. 
    - For `sourceType: connectedApp`: It needs to define the application name.
    - For `sourceType: internal`: It needs to be define the source name. For example, the name of the lamda function etc.
    - For `sourceType: external`: It needs to provide the correct source as required for the external source.

- `type`: Defines the event name for topic we need to subscribe for messages. 
    - For `sourceType: connectedApp|internal`: It needs to have atleast two segments i.e. `<event>.<version>`.
    - For `sourceType: external`: It needs to provide complete event type including any prefixes.


## Flow of Eventing:

Let suppose the following:
- EventType: `commercemock.order.created.v1`
- JetStreamPrefix: `kyma`
- EventTypePrefix: `sap.kyma.custom`

### Case 1: Event sent by a connected application

```
apiVersion: eventing.kyma-project.io/v1alpha1
kind: Subscription
metadata:
  name: example-sub
  namespace: goldfish
spec:
  sink: http://test.goldfish.svc.cluster.local
  events:
  - sourceType: connectedApp
    source: commercemock
    type: order.created.v1
```

**Subscribing**

Eventing Controller will internally prepend `source` to the `type` (i.e. `<source>.<type>`) irrespective of the eventing backend.

Eventing backend specific modifications:

- **JetStream:**
    - Eventing Controller will internally prepend `JetStreamPrefix` to the event type before subscribing on NATS Server. Therefore the final topic for JetStream will be `<JetStreamPrefix>.<source>.<type>`.
        - For example: `kyma.commercemock.order.created.v1`

- **EventMesh:**
    - Eventing Controller will internally prepend `EventTypePrefix` to the event type before subscribing on EventMesh Server.
        - For example: `sap.kyma.custom.commercemock.order.created.v1`

**Publishing**

User will publish events on `order.created.v1`.

For legacy events, the user will publish the events on url: `http://<host>/<applicationName>/v1/events`. For example:
```
curl -v -X POST \
    -H "Content-Type: application/json" \
    -H "X-B3-Sampled: 1" \
    --data @<(<<EOF
    {
		"event-type": "order.created",
		"event-type-version": "v1",
		"event-time": "2020-09-28T14:47:16.491Z",
		"data": {"orderCode": "987"}
  	}
EOF
    ) \
    http://localhost:8081/commercemock/v1/events
```

The **legacy events** are converted into cloud events by Event publisher proxy. The application name is extracted from the url and prepended to the `event-type` before publishing the event to the eventing backend. Also, the header `ce-source` in Cloud Event header will be added as required by the eventing-backend.

For **cloud events**, the user will publish the events on Url: `http://<host>/publish`. For example:
```
curl -v -X POST \
     -H "ce-specversion: 1.0" \
     -H "ce-type: order.created.v1" \
     -H "ce-source: commercemock" \
     -H "ce-eventtypeversion: v1" \
     -H "ce-id: A234-1234-1234" \
     -H "content-type: application/json" \
     -d "{\"foo\":\"binary-mode\"}" \
     http://localhost:8081/publish
```

The user needs to provide the application name in the `ce-source`.

Eventing backend specific modifications:

- **JetStream:** Internally EPP will append `JetStreamPrefix` before forwarding it to the NATS Server. 
    - For example: `kyma.commercemock.order.created.v1`
    - User --> `order.created.v1` --> EPP --> `kyma.commercemock.order.created.v1` --> NATS
- **EventMesh:** Internally EPP will append `EventTypePrefix` in case of legacy events. 
    - For example: `sap.kyma.custom.commercemock.order.created.v1`
    - User --> `order.created.v1` --> EPP --> `sap.kyma.custom.commercemock.order.created.v1` --> EventMesh

### Case 2: Event sent from within the cluster

```
apiVersion: eventing.kyma-project.io/v1alpha1
kind: Subscription
metadata:
  name: example-sub
  namespace: goldfish
spec:
  sink: http://test.goldfish.svc.cluster.local
  events:
  - sourceType: internal
    source: goldfish
    type: order.created.v1
```

The flow of in-cluster is same as in Connected Applications.

### Case 3: Event sent directly to eventing-backned (not through EPP)

```
apiVersion: eventing.kyma-project.io/v1alpha1
kind: Subscription
metadata:
  name: example-sub
  namespace: goldfish
spec:
  sink: http://test.goldfish.svc.cluster.local
  events:
  - sourceType: external
    source: /default/sap.kyma/tunas-develop
    type: sap.external.custom.exampleapp.order.created.v1
```

**JetStream:**
- Not Supported! Because NATS Servers are deployed inside the Kyma cluster and they are not exposed externally so users cannot directly publish events to NATS Servers.

**EventMesh:**
- Subscribing: 
    - EC will not prepend append anything EventType. It will subscribe to `sap.kyma.custom.commercemock.order.created.v1` on EventMesh Server.
    - The type will not be cleaned by Eventing Controller from non-alphanumeric characters.
- Publishing: 
    - User will directly publish events on EventMesh Server as `sap.kyma.custom.commercemock.order.created.v1`.