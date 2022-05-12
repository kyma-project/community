# Design Proposal - Eventing Backend CRD

## Table of Contents

- [Concept](#concept)
- [Metadata](#metadata)
- [Spec](#spec)
- [Validation](#validation)
- [Status](#status)
- [Discovery](#discovery)

## Concept

The goal of this proposal is to design the backend CRD in a flexible way to support:

- Creating multiple backends of different types (e.g. `nats`, `beb` ...etc).
- Creating multiple backends of the same type (e.g. `beb-instace-1`, `beb-instace-2`).
- Providing backend specific configurations in a generic way.
- Discovering the supported backends by Kyma Eventing.
- Discovering backend specific configurations.

```yaml
apiVersion: eventing.kyma-project.io/v1alpha1
kind: EventingBackend
metadata:
  name: any
  namespace: any
  labels:
    kyma-project.io/eventing: backend
spec:
  type: "nats"
  config:
  - maxInFlightMessages: 10
status:
  backendType: nats
  conditions:
  - lastTransitionTime: "2022-05-06T08:35:11Z"
    reason: Backend ready
    status: "True"
    type: Backend Ready
  - lastTransitionTime: "2022-05-06T08:35:11Z"
    reason: Publisher proxy deployment ready
    status: "True"
    type: Publisher Proxy Ready
  - lastTransitionTime: "2022-05-06T08:35:11Z"
    reason: Subscription controller started
    status: "True"
    type: Subscription Controller Ready
  eventingReady: true
```

> **Note:** As a result of having a flexible design for the backend configuration in the CRD, we decided not to mention which configuration needs to move from the Eventing controller charts to the backend CR and left it to the implementation phase.
## Metadata

The `namespaces` and `name` can be set by users. This allows to run multiple different instances of the same backend type.

```yaml
metadata:
  name: any
  namespace: any
  labels:
    kyma-project.io/eventing: backend
```

## Spec

Each backend instance is defined with its own `spec`. The `spec` contains the objects `type` and `config`. While `type` accepts the string of the name of the backend (`"nats"`, `"beb"`, ...etc), the `config` accepts the key/value pairs of the actual configuration for the specified backend.


Example how to configure `BEB` as the Eventing backend:

```yaml
spec:
  type: "beb"
  config:
  - secretRef: "some-secret"
```

## Validation

- Spec config validation can be done by writing a custom `ValidatingAdmissionWebhook`.
- Defaulting the required but not-set spec config can be done by writing a custom `MutatingAdmissionWebhook`.

In addition, the following should result in adding an error entry under the `status.validationErrors`:
- Spec config keys that do not belong to the specified backend.
- Spec config values set to forbidden data.
- Spec config values using the wrong data type for the given key.

## Status

The `status` contains the following fields:
- `backendType` contains the type of the backend which is provisioned or used.
- `validationErrors` contains the the errors that appear during the spec validation phase.
- `conditions` contains the readiness of all the components which the backend depends on.
- `eventingReady` is evaluated by `ANDing` all the `status.conditions`.

Example of a backend CR with no validation errors:

```yaml
status:
  backendType: nats
  conditions:
  - lastTransitionTime: "2022-05-06T08:35:11Z"
    reason: Backend ready
    status: "True"
    type: Backend Ready
  - lastTransitionTime: "2022-05-06T08:35:11Z"
    reason: Publisher proxy deployment ready
    status: "True"
    type: Publisher Proxy Ready
  - lastTransitionTime: "2022-05-06T08:35:11Z"
    reason: Subscription controller started
    status: "True"
    type: Subscription Controller Ready
  eventingReady: true
```

Example of a backend CR with validation errors:

```yaml
status:
  backendType: nats
  validationErrors:
  - type: "invalid-data-type-error"
    config: "maxInFlightMessages"
    value: "ten"
    error: "should be of type int"
  - type: "invalid-value-error"
    config: "foobar"
    value: "100"
    error: "must be set between 1000 and 9999"
  - type: "unsupported-config-error"
    config: "foo"
    value: "bar"
    error: "config does not exist for backend type nats"
  conditions:
  - lastTransitionTime: "2022-05-06T08:35:11Z"
    reason: The provided spec.config has validation errors
    status: "False"
    type: Backend ready
  - lastTransitionTime: "2022-05-06T08:35:11Z"
    reason: Publisher proxy deployment ready
    status: "True"
    type: Publisher Proxy Ready
  - lastTransitionTime: "2022-05-06T08:35:11Z"
    reason: Subscription controller started
    status: "True"
    type: Subscription Controller Ready
  eventingReady: false
```

> **Note:** The `Backend Ready` condition should indicate if the configured backend type in the spec is ready for use or not. It can be as simple as checking the underlying connection status with the backend. But that decision is left to the implementation phase.
## Discovery

Eventing users have the ability to discover the supported backends and their available configurations by doing get requests to the following endpoints:

- GET `http:/eventing-controller.kyma-system/backend`. The response should look like:
  - Content-Type: `application/json`.
  - Body:
  ```json
  ["nats","beb", "..."]
  ```
- GET `http:/eventing-controller.kyma-system/backend/${TYPE}`. The response should look like:
  - Content-Type: `application/json`.
  - Body:
  ```json
  {
     "type":"nats",
     "config":[
        {
           "description":"Some description",
           "dataType":"Some data type",
           "defaultValue":"Some default value"
        },
        {
           "description":"Some description",
           "dataType":"Some data type",
           "defaultValue":"Some default value"
        }
     ]
  }
  ```
