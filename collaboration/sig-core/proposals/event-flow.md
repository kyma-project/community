# EventFlow Proposal

Created on 2019-03-24 by Nils Schmidt (@nachtmaar)

## Status

Not proposed yet.

## Goal

Improve provisioning of Knative Event-Mesh related components and implement deprovisioning of those components.

## Current solution

The [Application-Broker](https://kyma-project.io/docs/components/application-connector/#architecture-application-connector-components-application-broker) is a broker which implements the [Open Service Broker API](https://www.openservicebrokerapi.org/) ⁰.
It basically does two things:

1. "The AB fetches all the applications' custom resources and exposes their APIs and Events as service classes to the Service Catalog" ⁴.
2. It is responsible for *provisioning* and *deprovisioning* of a *Service Class*.


### Provisioning Flow

When Application-Broker receives a provisioning request, it returns `202` HTTP status code and triggers asynchronous provisioning of the ServiceInstance - as described in provisioning spec⁵.
The Service Catalog then polls the `Last Operation for Service Instances endpoint` to watch the operation progress ⁶.

During the asynchronous provisioning of a ServiceInstance Application-Broker performs the following tasks in the given order:
1. Create `EventActivation`
2. Create `Knative Subscription` to enable event flow between Application Channel and Broker Channel
3. Label user namespace in order to let the Knative Namespace controller create the `Knative Broker`
4. Create `Istio Policy` (allows permissive communication between Knative Broker <-> Knative Broker filter)


## Motivation/Problems

There are three open issues related to provisioning/deprovisioning of resources which are created by the Application Broker.
This proposal tries to provide ideas how to solve them.

1. [Provisioning events service instance fails on kyma 1.9 #7193](https://github.com/kyma-project/kyma/issues/7193)

TL;DR:
- Application-Broker does not implement retries then a provisioning request fails. The creation of a Knative Subscription can fail, if the Application Channel has not been created yet.
  If that happens, the Service Instance will stay in failed status forever => **No Eventing!**
- According to the _Open Service Broker API_, it is the Service Brokers responsibility ³ to implement retries. In contrast to [Kyma Environment Broker (KEB)](https://github.com/kyma-incubator/compass/blob/master/docs/kyma-environment-broker/02-01-architecture.md), Application-Broker does not implement retries.
- Platform retry behaviour may be part of OSB API v3.

2: [Deprovision Knative Broker #6342](https://github.com/kyma-project/kyma/issues/6342)

TL;DR:
- There can be a race condition between provisioning and deprovisioning of a ServiceInstance.
- Consequence could be a user namespace without Knative Broker => **No eventing!**
- Therefore, deprovisioning of a Knative Broker is not implemented at the moment.

3: [Reflect Status of Knative Broker on ServiceInstance provisioning status #7696](https://github.com/kyma-project/kyma/issues/7696)

TL;DR:
- We can check the status of the Broker, but it would block the provisioning request for some time. If the broker is not ready, we mark the provisioning request as failed. However, if the broker heals itself (e.g. reconciliation or dependent objects heal), 
  the status of the ServiceInstance will stay in failed status.

## Possible Solutions

### Custom Eventing Controller

Instead of performing steps 2-4 inside Application-Broker, the necessary steps to enable eventing for a Kyma Application should be `delegated` to a `new Kubernetes controller`.

How would we solve the above mentioned problems with this solution:

Problem 1: The `EventFlow` CR (Custom Resource) would be created by the Application-Broker.
   The chance that the creation of EventFlow CR fails due to an error in the provisioning request (inside Application-Broker - no retries) is very low.
   Therefore, retries in the Application-Broker are not required anymore.
   If the Application Channel is not ready, an Informer on the Knative Service (which reflects the status of the Application Channel as well) would trigger a new reconciliation,
   therefore implementing the required retries.

Problem 2: Whenever a ServiceInstance is provisioned by the Application-Broker, an EventFlow CR would be created in the same namespace as well.
   That means we know which applications are relying on a Knative Broker.
   When a deprovisiong request arrives in the Application-Broker, we delete the appropriate EventFlow CR.
   If there are no EventFlow CRs left, we can safely delete the Knative Broker.
   If a new provisioning request arrives in the meantime, we wait for the deprovisioning of the Knative Broker and trigger the creation of a new Knative Broker. If the broker is ready, the EventFlow CR status is set to ready.

Problem 3: We can easily reflect the status of the Knative Broker inside the EventFlow status using an Informer on the Knative Broker. We won't reflect the status of the Knative Broker in the ServiceInstance.

#### Example of EventFlow CR

Status of EventFlow could be based on the following objects:
- the HTTP Source Adapter (Knative Service) - if the HTTP Source is down, no events are entering the event-mesh
- the Knative Broker - if the Broker is down, the events will not be dispatched to  the 
- Knative Subscription

```yaml
apiVersion: eventing.kyma-project.io/v1alpha1
kind: EventFlow
metadata:
  name: event-flow-<service-instance-name> # e.g. event-flow-es-all-events-07ee5-bumpy-communication
  namespace: <user-namespace> # e.g my-user-namespace
  finalizers:
  - event-flow-controller # required to deprovision Knative Broker
spec:
  source: <application-name> # same name as Application/HTTP Source
status:
  SinkURI: http://default-broker.default.svc.cluster.local  # URI of the Knative Broker - field has to be in status because it will be set by the EventFlow controller, not Application-Broker
  conditions:
  - lastTransitionTime: "2020-03-23T09:52:47Z"
    status: "True"
    type: SourceReady # <- source is abstraction over HTTPSource, could be any sources.kyma-project.io/v1alpha1 or even duck type
  - lastTransitionTime: "2020-03-23T09:52:32Z"
    severity: Info
    status: "True"
    type: SinkReady # <- sink is abstraction over Knative Broker, could be any sink
  - lastTransitionTime: "2020-03-23T09:52:32Z"
    severity: Info
    status: "True"
    type: SubscriptionReady # Knative Subscription between Application Channel <-> Broker
  - lastTransitionTime: "2020-03-23T09:52:31Z"
    status: "True"
    type: Ready
```

#### Implementation

1. Use `knative.dev/pkg/controller` to implement the controller - similar to [HTTP Source controller](https://github.com/kyma-project/kyma/blob/bb5810fdb969035617bb0fd70f0d1d1d91bea58b/components/event-sources/reconciler/httpsource/controller.go#L63)
1. Remove steps 2-4 from Application-Broker. Instead create/delete EventFlow CR when Application-Broker receives provisioning/deprovisiong request. 
1. Add EventFlow controller component to Kyma
1. Bump Application-Broker image


#### Migration/Upgrade

In order to upgrade Kyma to a new version with EventFlow controller in place, we need to 
- recreate all ServiceInstances belong to Application-Broker (similar to [event-mesh-migration](https://github.com/kyma-project/kyma/blob/ed5c61f07e2e22172fcaf045fb55c7bc10336f55/components/event-bus/cmd/mesh-namespaces-migration/main.go#L68))
   - this will delete all objects created in steps 2-4
   - create EventFlow CR
   - EventFlow controller will create objects from steps 2-4

If we don't perform the migration, we won't have an EventFlow CR for old ServiceInstances and can't implement the deprovision correctly. 

Also, the upgrade test has to be modified to wait for the readiness of the EventFlow CR [here](https://github.com/kyma-project/kyma/blob/a60d814eb91ea1e97f7fb1516f78227b73fe1e3a/tests/end-to-end/upgrade/pkg/tests/eventmesh-migration/migrate-eventmesh/eventmesh.go#L59) and [here](https://github.com/kyma-project/kyma/blob/a60d814eb91ea1e97f7fb1516f78227b73fe1e3a/tests/end-to-end/upgrade/pkg/tests/eventmesh-migration/migrate-eventmesh/eventmesh.go#L86)


#### Backup/Restore

- The backup tests needs to be modified to wait for the creation of the EventFlow CR [here](https://github.com/kyma-project/kyma/blob/1ac01ae13f2cd2de98d8069827e69e98e778bb7d/tests/end-to-end/backup/pkg/tests/eventmesh/eventmesh.go#L201) and [here](https://github.com/kyma-project/kyma/blob/1ac01ae13f2cd2de98d8069827e69e98e778bb7d/tests/end-to-end/backup/pkg/tests/eventmesh/eventmesh.go#L185).
- Since the backup test already waits for the ServiceInstance, the EventFlow CR does not need to be included in the backup itself. It will be recreated by the ApplicationBroker if the ServiceInstance get's created.


### Retry in Application-Broker

#### Queue-based Application Broker

Kyma Environment Broker implements retries as follows, we could do the same for Application-Broker:

- Kyma Environment Broker implements a work queue, where all provisioning operations are stored.
- The queue operations are backed by a Postgres database and populated at [Broker start time](https://github.com/kyma-incubator/compass/blob/cf15310a2ed8f0f90d6ef9d739ab6fb27651a717/components/kyma-environment-broker/cmd/broker/main.go#L159)
- If a provision request fails, it will get rescheduled (each step in KEB can define the retry interval based on the error).

Disadvantage: We would need to store some state to keep track of provisioning operations.


#### Informer-Based Application Broker

Implement retries for failed ServiceInstances in Application-Broker:

- We could use an `Informer` for `ServiceInstances` and trigger an update of the failed ServiceInstances (only the ones which belong to Application-Broker). 
- OSB API mentions an [update REST endpoint](https://github.com/kyma-project/kyma/blob/f2c3b3498f91c22250ddbc7a6a4449b679a40263/components/application-broker/internal/broker/server.go#L143).
  Unfortunately, Application-Broker does not implement the endpoint yet. See code endpoints [here](https://github.com/kyma-project/kyma/blob/f2c3b3498f91c22250ddbc7a6a4449b679a40263/components/application-broker/internal/broker/server.go#L143)

Disadvantages: The provisioning/deprovisioning endpoints are called by ServiceCatalog. We would need to call the the update endpoint from Application-Broker which feels hacky.


#### Blocking provisioning operation in Application Broker

We can implement retries ourselves without keeping any state.
A simple retry loop would be enough to delay the creation of the Knative Subscription until the Application Channel exists and is ready.

Disadvantage: Every provisioning/deprovisioning request is executed in a goroutine. By waiting for the Channel we are blocking/leaking the goroutine.


# Sources

0: <https://kyma-project.io/docs/components/application-connector/#architecture-application-connector-components-application-broker>

1: <https://www.openservicebrokerapi.org/>

2: <https://tanzu.vmware.com/content/white-papers/an-inside-look-at-the-open-service-broker-api>

3: <https://github.com/kyma-project/kyma/issues/7288#issuecomment-600908082>

4: <https://github.com/kyma-project/kyma/blob/2d11b3a22b56c327aee1d2920311d9623cff5910/components/application-broker/README.md#L17-L16>

5: <https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#provisioning>

6: <https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#polling-last-operation-for-service-instances>

7: <https://github.com/kyma-project/community/pull/386/files>

8: <https://kyma-project.io/docs/components/event-bus/#details-knative-eventing-mesh-alpha>

9: <https://github.com/kyma-incubator/compass/blob/master/docs/kyma-environment-broker/03-03-add-provisioning-step.md>