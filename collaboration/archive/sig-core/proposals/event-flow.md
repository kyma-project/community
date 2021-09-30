# EventFlow Proposal

Created on 2019-03-24 by Nils Schmidt (@nachtmaar)

## Status

Not proposed yet.

## Goal

Improve the provisioning of Knative Eventing Mesh-related components and implement deprovisioning of those components.

## Current solution

The [Application Broker](https://kyma-project.io/docs/components/application-connector/#architecture-application-connector-components-application-broker) is a broker which implements the [Open Service Broker API](https://www.openservicebrokerapi.org/) ⁰.
It basically does two things:

1. "The AB fetches all the applications' custom resources and exposes their APIs and Events as service classes to the Service Catalog" ⁴.
2. It is responsible for provisioning and deprovisioning of a ServiceClass.


### Provisioning Flow

When the Application Broker receives a provisioning request, it returns `202` HTTP status code and triggers asynchronous provisioning of the Service Instance - as described in provisioning spec⁵.
The Service Catalog then polls the `Last Operation for Service Instances endpoint` to watch the operation progress ⁶.

During the asynchronous provisioning of a Service Instance, the Application Broker performs the following tasks in the given order:
1. Creates `EventActivation`.
2. Creates `Knative Subscription` to enable event flow between Application Channel and Broker Channel.
3. Labels user namespace in order to let the Knative Namespace controller create the `Knative Broker`.
4. Creates `Istio Policy` (allows Prometheus - which is outside the Service Mesh - to scrape metrics from the Knative Broker Pods).

For more details on how Application Broker is involved in Knative Eventing Mesh, see [this document](https://github.com/kyma-project/kyma/blob/release-1.11/docs/knative-eventing-mesh/02-01-knative-eventing-mesh.md#component-dependencies).


## Motivation/Problems

There are three open issues related to provisioning/deprovisioning of resources which are created by the Application Broker.
This proposal tries to provide ideas how to solve them.

1: [Provisioning events service instance fails on kyma 1.9 #7193](https://github.com/kyma-project/kyma/issues/7193)

TL;DR:
- Application Broker does not implement retries for failed provisioning requests. The creation of a Knative Subscription can fail, if the Application Channel has not been created yet.
  If that happens, the Service Instance will stay in failed status forever => **No Eventing!**
- According to the _Open Service Broker API_, it is the Service Broker's responsibility ³ to implement retries. In contrast to Kyma Environment Broker (KEB), Application Broker does not implement retries.
- Platform retry behavior may be part of OSB API v3.

2: [Deprovision Knative Broker #6342](https://github.com/kyma-project/kyma/issues/6342)

TL;DR:
- There can be a race condition between provisioning and deprovisioning of a Service Instance.
- Consequence could be a user namespace without Knative Broker => **No eventing!**
- Therefore, deprovisioning of a Knative Broker is **not implemented** at the moment.
- (*) The only way I can think of solving this with the current architecture is having a **mutex** in provisioner/deprovisioner:
    - In the provisioning step we need to check that the Knative Broker exists. If not, label the namespace and wait for successful creation of the Broker.
    - In the deprovisioning step, we need a way to check if the Knative Broker is in use, e.g. listing all successfully provisioned Service Instances belonging to Application-Broker.
    - If the Broker is not in use, we can delete it. No provisioning can happen meanwhile (mutex), therefore no race.

3: [Reflect Status of Knative Broker on ServiceInstance provisioning status #7696](https://github.com/kyma-project/kyma/issues/7696)

TL;DR:
- We can check the status of the Broker, but it would block the provisioning request for some time. If the broker is not ready, we mark the provisioning request as failed. However, if the broker heals itself (e.g. reconciliation or dependent objects heal), the status of the Service Instance will stay in failed status.

## Possible Solutions

### Custom Eventing Controller

Instead of performing steps 2-4 inside Application Broker, the necessary steps to enable eventing for a Kyma Application should be `delegated` to a `new Kubernetes controller`.

How would we solve the above mentioned problems with this solution:

*Problem 1*: The `EventFlow` CR (Custom Resource) would be created by the Application Broker.
   The chance that the creation of EventFlow CR fails due to an error in the provisioning request (inside Application Broker - no retries) is very low.
   Therefore, retries in the Application Broker are not required anymore.
   If the Application Channel is not ready, an Informer on the Knative Service (which reflects the status of the Application Channel as well) would trigger a new reconciliation,
   therefore implementing the required retries.
   If the EventFlow CR is successfully created, but the EventFlow status is not ready, then the user will think that eventing is enabled (because UI uses ServiceInstance status only), however, there was an error.
   In order to provide the user better feedback, the UI needs to query the Service Instance as well as the EventFlow status. See implementation for more details.

*Problem 2*: Whenever a Service Instance is provisioned by the Application Broker, an EventFlow CR would be created in the same namespace as well.
   That means we know which applications are relying on a Knative Broker.
   When a deprovisiong request arrives in the Application Broker, we delete the appropriate EventFlow CR.
   If there are no EventFlow CRs left, we can safely delete the Knative Broker.
   If a new provisioning request arrives in the meantime, we wait for the deprovisioning of the Knative Broker and trigger the creation of a new Knative Broker. If the broker is ready, the EventFlow CR status is set to ready.

*Problem 3*: We can easily reflect the status of the Knative Broker inside the EventFlow status using an Informer on the Knative Broker.
    We won't reflect the status of the Knative Broker in the Service Instance.


#### Example of EventFlow CR

Status of EventFlow could be based on the following objects:
- the HTTP Source Adapter (Knative Service) - if the HTTP Source is down, no events are entering the event-mesh
- the Knative Broker - if the Broker is down, the events will not be dispatched to the final sink (e.g. a Lambda)
- Knative Subscription - if the subscription is unready, the events won't be forwarded to the Broker

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
  SinkURI: http://default-broker.default.svc.cluster.local  # URI of the Knative Broker - field has to be in status because it will be set by the EventFlow controller, not Application Broker
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
1. Remove steps 2-4 from Application Broker. Instead create/delete EventFlow CR when Application Broker receives provisioning/deprovisiong request.
1. Implement combined status of Service Instance & EventFlow in UI
1. Add EventFlow controller component to Kyma resources folder, therefore creating a new Kyma installer image
1. Bump Application Broker image
1. Bump UI image


*UI Changes*:

Currently the status of the ServiceInstance can be one of the following:
1. `Provisioning`
1. `Running`
1. `Failed`
1. `Deprovisioning`

We have two options to display the status of both resources in the UI:
1. Combined status of both resources in one string.
   The following table illustrates how we can create the combined status based on the status of the ServiceInstance and the EventFlow:

   | Combined Status | Service Instance           | EventFlow      |
   |-----------------|----------------------------|----------------|
   | Provisoning     | Status.CurrentOperation && | any            |
   | Running         | Status ready and           | Status ready   |
   | Failed          | Status unready or          | Status unready |
   | Deprovisioning  | Status.CurrentOperation and| any            |

   `Provisioning`  : is the case then the `Status.CurrentOperation` field of the ServiceInstance is equal to `Provision`
   `Running`       : is the case then the status of the ServiceInstance and the EventFlow is ready
   `Failed`        : is the case then the status of the ServiceInstance or the EventFlow is unready
   `Deprovisioning`: is the case then the `Status.CurrentOperation` field of the ServiceInstance is equal to `Deprovision`

    Advantage: Easier to understand for the user

    Disadvantage: More changes in UI required

1. Show status of both resources instead of a single status

   Advantage: Easier to implement

   Disadvantage: Harder for the user to understand

In both cases the tooltips for the status need to be revisited. Currently it is based on the ServiceInstance status conditions only.
If the EventFlow CR is in unready state, the tooltip has to be taken from the EventFlow instead.

#### Migration/Upgrade

In order to upgrade Kyma to a new version with EventFlow controller in place, we need to
- recreate all Service Instances belong to Application Broker (similar to [event-mesh-migration](https://github.com/kyma-project/kyma/blob/ed5c61f07e2e22172fcaf045fb55c7bc10336f55/components/event-bus/cmd/mesh-namespaces-migration/main.go#L68))
   - this will delete all objects created in steps 2-4
   - create EventFlow CR
   - EventFlow controller will create objects from steps 2-4

If we don't perform the migration, we won't have an EventFlow CR for old Service Instances and can't implement the deprovision correctly.

Also, the upgrade test has to be modified to wait for the readiness of the EventFlow CR [here](https://github.com/kyma-project/kyma/blob/a60d814eb91ea1e97f7fb1516f78227b73fe1e3a/tests/end-to-end/upgrade/pkg/tests/eventmesh-migration/migrate-eventmesh/eventmesh.go#L59) and [here](https://github.com/kyma-project/kyma/blob/a60d814eb91ea1e97f7fb1516f78227b73fe1e3a/tests/end-to-end/upgrade/pkg/tests/eventmesh-migration/migrate-eventmesh/eventmesh.go#L86)


#### Backup/Restore

- The backup tests needs to be modified to wait for the creation of the EventFlow CR [here](https://github.com/kyma-project/kyma/blob/1ac01ae13f2cd2de98d8069827e69e98e778bb7d/tests/end-to-end/backup/pkg/tests/eventmesh/eventmesh.go#L201) and [here](https://github.com/kyma-project/kyma/blob/1ac01ae13f2cd2de98d8069827e69e98e778bb7d/tests/end-to-end/backup/pkg/tests/eventmesh/eventmesh.go#L185).
- Since the backup test already waits for the Service Instance, the EventFlow CR does not need to be included in the backup itself. It will be recreated by the ApplicationBroker if the Service Instance gets created.


### Retry in Application Broker

#### Queue-based Application Broker

Kyma Environment Broker implements retries as follows, we could do the same for the Application Broker:

- Kyma Environment Broker implements a work queue, where all provisioning operations are stored.
- The queue operations are backed by a Postgres database and populated at [Broker start time](https://github.com/kyma-incubator/compass/blob/cf15310a2ed8f0f90d6ef9d739ab6fb27651a717/components/kyma-environment-broker/cmd/broker/main.go#L159)
- If a provision request fails, it will get rescheduled (each step in KEB can define the retry interval based on the error).

*Problem 1*: Fixes the problem.

*Problem 2*: In theory possible to solve, see (*).
             We could use the database to remember which Service Instance uses the Knative Broker.

*Problem 3*: Feature can be implemented, but when the status of the Broker changes, it won't be reflected in the Service Instance unless we have an Informer for the Broker.

Disadvantage: We would need to store some state to keep track of provisioning operations.


#### Informer-Based Application Broker

Implement retries for failed Service Instances in Application Broker:

- We could use an `Informer` for `Service Instances` and trigger an update of the failed Service Instances (only the ones which belong to Application Broker).
- OSB API mentions an [update REST endpoint](https://github.com/kyma-project/kyma/blob/f2c3b3498f91c22250ddbc7a6a4449b679a40263/components/application-broker/internal/broker/server.go#L143).
  Unfortunately, Application Broker does not implement the endpoint yet. See code endpoints [here](https://github.com/kyma-project/kyma/blob/f2c3b3498f91c22250ddbc7a6a4449b679a40263/components/application-broker/internal/broker/server.go#L143)

*Problem 1*: By having an Informer on Service Instance we can react on any change on the object. We still need an in-memory queue and scheduler to retry the operation (e.g. exponential backoff).

*Problem 2*: In theory possible to solve, see (*).
             Informer doesn't help. A Lister is required to know that a Service Instance is using the Knative Broker.

*Problem 3*: Feature can be implemented, but when the status of the Broker changes, it won't be reflected in the Service Instance unless we have an Informer for the Broker.

Disadvantages: The provisioning/deprovisioning endpoints are called by ServiceCatalog. We would need to call the the update endpoint from Application Broker which feels hacky.


#### Blocking provisioning operation in Application Broker


We can implement retries ourselves without keeping any state.
A simple retry loop would be enough to delay the creation of the Knative Subscription until the Application Channel exists and is ready.

*Problem 1*: Improves the situation by adding at least retries for a reasonable time.
             But this solution will need to have a good compromise on the timeout for the retry.
             If timeout is too low, it might not solve the problem, if timeout is too high we block goroutines too long.

*Problem 2*: In theory possible to solve, see (*).

*Problem 3*: Feature can be implemented, but when the status of the Broker changes, it won't be reflected in the Service Instance unless we have an Informer for the Broker.

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
