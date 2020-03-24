todo: move file to sig-and-wg/sig-core/proposals

# EventFlow Proposal


## Current Status

The [Application-Broker](https://kyma-project.io/docs/components/application-connector/#architecture-application-connector-components-application-broker) is a broker which implements the [Open Service Broker API](https://www.openservicebrokerapi.org/) ⁰.
It basically does two things:

1. "The AB fetches all the applications' custom resources and exposes their APIs and Events as service classes to the Service Catalog" ⁴.
2. It is responsible for *provisioning* and *deprovisioning* of a *ServiceInstance*.

TODO: provisioning flow diagram

### Provisioning Flow

When Application-Broker receives a provisioning request, it returns 202 HTTP status code and triggers asynchronous provisioning of the ServiceInstance - as described in provisioning spec⁵.
The platform then polls the `Last Operation for Service Instances endpoint` to watch the operation progress ⁶.

During the asynchronous provisioning of a ServiceInstance Application-Broker performs the following tasks in the following order:
1. Create `EventActivation`
2. Create `Knative Subscription` to enable event flow between Application Channel and Broker Channel
3. Label user namespace in order to let the Knative Namespace controller create the `Knative Broker`
4. Create `Istio Policy` (allows permissive communication between broker <-> broker filter)
5. TODO: authorization policy ??

### Problems

1. [Provisioning events service instance fails on kyma 1.9 #7193](https://github.com/kyma-project/kyma/issues/7193)
TL;DR:
- Application-Broker does not implement retries then a provisioning request fails. The creation of a Knative Subscription can fail, if the Application Channel has not been created yet.
  If that happens, the ServiceInstance will stay in failed status forever => **No Eventing!**
- According to the _Open Service Broker API_, it is the service brokers responsibility to implement retries. In contrast to Kyma Environment Broker (KEB), Application-Broker does not implement retries.
- Platform retry behaviour may be part of OSB API v3.

2: [Deprovision Knative Broker #6342](https://github.com/kyma-project/kyma/issues/6342)
TL;DR:
- There can be a race condition between provisioning and deprovisioning of a ServiceInstance.
- Consequence would be a user namespace without Knative Broker => **No eventing!**
- Therefore, deprovisioning is not implemented at the moment.

3: [Reflect Status of Knative Broker on ServiceInstance provisioning status #7696](https://github.com/kyma-project/kyma/issues/7696)
TL;DR:
- We can check the status of the Broker, but it would block the provisioning request for some time. If the broker is not ready, we mark the provisioning request as failed. However, if the broker heals itself (e.g. reconciliation or dependent objects heal), 
  the status of the ServiceInstance will stay in failed status.

## Possible Solutions

### Custom Eventing Controller

Instead of performing steps 2-5 inside Application-Broker, the necessary steps to enable eventing for a Kyma Application should be delegated to a new Kubernetes controller.
In contrast to a Service Broker, a Kubernetes controller supports reconciliation such that we can react on changes on Kubernetes objects.

Status of EventFlow could be based on the following objects:
- the HTTP Source Adapter (Knative Service)
- the Knative Broker
- Knative Subscription

How would we solve the above mentioned problems with this solution ?
1. The EventFlow CR (Custom Resource) would be created by the Application-Broker.
   The chance that the creation of the CR fails due to an error is very low. Therefore, retries in the Application-Broker are not required anymore.
   If the Application Channel is not ready, an Informer on the Knative Service (which reflects the status of the Application Channel as well) would trigger a new reconciliation,
   therefore implementing the required retries.
2. Whenever a ServiceInstance is provisioned by the Application-Broker, an EventFlow CR would be created in the same namespace as well. That means we know which applications are relying on a Knative Broker.
   When a deprovisiong request arrives in the Application-Broker, we delete the appropriate EventFlow CR. If there are no EventFlow CRs left, we can safely delete the Knative Broker.
   If a new provisioning request arrives in the meantime, we wait for the deprovisioning of the Knative Broker and trigger the creation of a new Knative Broker. If the broker is ready, the EventFlow CR status is set to ready.
3. We can easily reflect the status of the Knative Broker inside the EventFlow status using an Informer on the Knative Broker. We won't reflect the status of the Knative Broker in the ServiceInstance.

TODO: give example of EventFlow CR inclusive status

```yaml
apiVersion: <TODO>.kyma-project.io/v1alpha1
kind: EventFlow
metadata:
  name: event-flow-<service-instance-name> # e.g. event-flow-es-all-events-07ee5-bumpy-communication
  namespace: kyma-integration
spec:
  source: nachtmaar-test
status:
  SinkURI: http://default-broker.default.svc.cluster.local  # URI of the Knative Broker
  conditions:
  - lastTransitionTime: "2020-03-23T09:52:47Z"
    status: "True"
    type: SourceReady # <- source is abstraction over HTTPSource, could be any sources.kyma-project.io/v1alpha1 or even duck type
  - lastTransitionTime: "2020-03-23T09:52:32Z"
    severity: Info
    status: "True"
    type: SinkReady # <- sink is abstraction over Knative Broker, could be any sink
  - lastTransitionTime: "2020-03-23T09:52:31Z"
    status: "True"
    type: Ready
```


TODO: apigroup?

TODO: implementation ideas:
- knative eventing controller pkg  ... similiar to event-sources controller manager

finalizer on EventFlow



#### Migration/Upgrade

TODO: What is required to implement the change. ??

1. Remove steps 2-5 from Application-Broker. Instead create and delete EventFlow CR when Application-Broker receives provisioning/deprovisiong request. 
2. Implement the changes in EventFlow controller
3. Add EventFlow controller component to Kyma
4. Bump Application-Broker image

TODO: what is required to upgrade Kyma ?

TODO: what is required to backup Kyma ?
include EventFlow or not ?


### Retry in application-broker stateless

Application-Broker already has [informers](https://github.com/kyma-project/kyma/blob/f2c3b3498f91c22250ddbc7a6a4449b679a40263/components/application-broker/cmd/broker/main.go#L149) on:
- Applications
- Service Catalog (which CRs ??)
- ApplicationMappings
- Namespaces

We could add an `Informer` for `ServiceInstances` and trigger an update of the failed ServiceInstances which belong to Application-Broker. 
OSB API mentions an [update REST endpoint](https://github.com/kyma-project/kyma/blob/f2c3b3498f91c22250ddbc7a6a4449b679a40263/components/application-broker/internal/broker/server.go#L143).

Implementation:
- Unfortunately, Application-Broker does not implement the endpoint yet. See code endpoints [here](https://github.com/kyma-project/kyma/blob/f2c3b3498f91c22250ddbc7a6a4449b679a40263/components/application-broker/internal/broker/server.go#L143)
- Finding ServiceInstances belonging to Application-Broker is already done [here](https://github.com/kyma-project/kyma/blob/ed5c61f07e2e22172fcaf045fb55c7bc10336f55/components/event-bus/cmd/mesh-namespaces-migration/serviceinstance.go#L72)

Disadvantages: 


### Retry in application-broker stateful

Kyma Environment Broker implements retries using a database for storing state. That requires each step to be idempotent in order that each step can be repeated if one step fails ⁹.

Disadvantage: We are stateful :/


### Block provisioning call in application-broker

We can implement retries ourselves without keeping any state.
A simple retry loop would be enough to delay the creation of the Knative Subscription until the Application Channel exists and is ready.

Disadvantage: Every provisioning/deprovisiong request is executed in a goroutine. By waiting for the Channel we are blocking/leaking the goroutine.


## Problem

# Sources

0: <https://kyma-project.io/docs/components/application-connector/#architecture-application-connector-components-application-broker>

1: <https://www.openservicebrokerapi.org/>

2: <https://tanzu.vmware.com/content/white-papers/an-inside-look-at-the-open-service-broker-api>

3: <https://github.com/kyma-project/kyma/issues/7288#issuecomment-600908082>

4: <https://github.com/kyma-project/kyma/blob/2d11b3a22b56c327aee1d2920311d9623cff5910/components/application-broker/README.md#L17-L16>

5: <https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#provisioning>

6: <https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#polling-last-operation-for-service-instances>

7: <https://github.com/kyma-project/community/pull/386/files>

8: https://kyma-project.io/docs/components/event-bus/#details-knative-eventing-mesh-alpha

9: <https://github.com/kyma-incubator/compass/blob/master/docs/kyma-environment-broker/03-03-add-provisioning-step.md>