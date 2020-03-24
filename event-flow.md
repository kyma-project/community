# EventFlow Proposal


## Current Status

The [Application-Broker](https://kyma-project.io/docs/components/application-connector/#architecture-application-connector-components-application-broker) is a broker which implements the [Open Service Broker API](https://www.openservicebrokerapi.org/) ⁰.
It basically does two things:

1. "The AB fetches all the applications' custom resources and exposes their APIs and Events as service classes to the Service Catalog" ⁴.
2. It is responsible for *provisioning* and *deprovisioning* of a *ServiceInstance*

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
In contrast to a ServiceBroker, a Kubernetes controller supports reconciliation such that we can react on changes on 
- the HTTP Source Adapter (Knative Service)
- the Knative Broker
- 

#### Migration/Upgrade

TODO: What is required to implement the change. ??

### Retry in application-broker stateless

### Retry in application-broker stateful

### Block provisioning call in application-broker


## Problem

# Sources

0: <https://kyma-project.io/docs/components/application-connector/#architecture-application-connector-components-application-broker>

1: <https://www.openservicebrokerapi.org/>

2: <https://tanzu.vmware.com/content/white-papers/an-inside-look-at-the-open-service-broker-api>

3: <https://github.com/kyma-project/kyma/issues/7288#issuecomment-600908082>

4: <https://github.com/kyma-project/kyma/blob/2d11b3a22b56c327aee1d2920311d9623cff5910/components/application-broker/README.md#L17-L16>

5: <https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#provisioning>

6: <https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#polling-last-operation-for-service-instances>