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

- Creating multiple backends of different types (for example, `nats`, `beb`).
- Creating multiple backends of the same type (for example, `beb-instace-1`, `beb-instace-2`).
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
  conditions:
  - lastTransitionTime: "2022-05-06T08:35:11Z"
    reason: Eventing Backend Ready
    status: "True"
    type: Eventing Backend Ready
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

> **Note:** As a result of having a flexible design for the backend configuration in the CRD, we decided not to mention which configuration needs to move from the Eventing controller charts to the backend CR. It is left for the implementation phase.

## Metadata

You can set `namespaces` and `name`. This allows you to run multiple different instances of the same backend type.

```yaml
metadata:
  name: any
  namespace: any
  labels:
    kyma-project.io/eventing: backend
```

## Spec

Each backend instance is defined with its own `spec`. The `spec` contains the objects `type` and `config`. While `type` accepts the string of the backend name (`"nats"`, `"beb"`), the `config` accepts the key-value pairs of the actual configuration for the specified backend.

You can find instructions how to configure `BEB` as the Eventing backend in the following example:

```yaml
spec:
  type: "beb"
  config:
  - secretRef: "some-secret"
```

## Validation

- Spec config validation can be done by writing a custom `ValidatingAdmissionWebhook`.
- Defaulting the required but not-set spec config can be done by writing a custom `MutatingAdmissionWebhook`.

The following scenarios must return a `validationError` and result in the backend not being created.
- Spec config keys that do not belong to the specified backend.
- Spec config values set to forbidden data.
- Spec config values using the wrong data type for the given key.
  Similar, changing the specs of an existing backend in an invalid way must be rejected and return a `validationError`.

## Status

`status` contains the following fields:
- `conditions` contains the readiness of all the components which the backend depends on.
- `eventingReady` is evaluated by `ANDing` all the `status.conditions`.

Example of a backend CR with no validation errors:

```yaml
status:
  conditions:
  - lastTransitionTime: "2022-05-06T08:35:11Z"
    reason: Eventing Backend Ready
    status: "True"
    type: Eventing Backend Ready
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

> **Note:** The `Eventing Backend Ready` condition indicates if the configured backend type in the spec is ready for use. It can be as simple as checking the underlying connection status with the backend. This decision is left for the implementation phase.

## Discovery

Eventing users must have the ability to discover the supported Eventing backends and their corresponding configurations.

There are multiple options to provide Eventing backends discovery:

1. Document the supported Eventing Backends and their corresponding configurations.
2. Provide discovery endpoints:
  - GET `/backend` returns the supported Eventing backends.
  - GET `/backend/${TYPE}` returns more details about the given backend type (e.g. configuration details).
3. Provide an `openAPIV3Schema` schema validation in the Eventing backend CRD.

> **Note:** For the time being we decided to go with the first option and in the future we can support the necessary discovery automation if needed.
