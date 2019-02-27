# Overview

As an operator, I want to configure a default messaging middleware that my Kyma cluster should use for eventing.

As a developer, I just want to configure a subscription for an event type from a source in Kyma and do not care about the underlying details.

## Requirement

* The pluggable Knative layer for the messaging middlewar (Knative ClusterChannelProvisioner) is available and deployable.
* Operator will provision the pluggable layer into Kyma (Knative ClusterChannelProvisioner).
  > Note: There will be some refernce examples that customers can follow. The provisioning/deployment will be an operator action as customes can choose a implementation specific to their needs.
* Kyma is updated to use Knative version 0.4+.

## The out-of-box Kyma installation

The OOB Kyma installation will have NATS Streaming set as the default messaging middleware. The required Knative configurations will be set to use that.

## When to specify the default?

### Post Installation

* Operator either creates a PubSub instance or use existing one.

  >Note: For instance creation, she can use service catalog. Ths is not a must.

* Operator deploys the cluster channel provisioner and injects the secrets.

* Use [Knative semantics](https://github.com/knative/docs/blob/master/eventing/channels/default-channels.md#setting-the-default-channel-configuration) to specify the default ClusterChannelProvisioner.

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: default-channel-webhook
  namespace: knative-eventing
data:
  default-channel-config: |
    clusterdefault:
      apiversion: eventing.knative.dev/v1alpha1
      kind: ClusterChannelProvisioner
      name: cloud-pubsub
    namespacedefaults:
      some-namespace:
        apiversion: eventing.knative.dev/v1alpha1
        kind: ClusterChannelProvisioner
        name: cloud-pubsub
```

### Why not during installation?

* There is no standard way to deploy a provisioner. It can be done via Helm Chart, plain kubernetes deployments or some other mechanism. 
  * We do not intend to impose a restriction as provisioners will be implemented in Knative open source or by customers or partners themselves.
* The injection of secrets during installation will be tricky.
* There may be scenarios to reuse some existing instance.

## What happens to NATS Streaming?

* Operator has the possibility to deprovision if she wishes to save costs. 
  >Note: The detailed deprovisioning/migration will be solved via collaboration with Knative community and will be solved in future.

**Concerns**

* It can lead to broken state if there are already some resources created with NATS Streaming.

## Can the default be changed?

### Yes, default can be changed

* The existing channels/subscriptions still stay the same with the previous PubSub. [Same approach](https://github.com/knative/docs/blob/master/eventing/channels/default-channels.md#caveats-1) as followed by Knative.

**Concerns**

* It can lead to broken state if operators deprovision the previous default.

## How to support the operator

With great power comes great responsibility. Operator will have power to leverage pluggability of messaginng middleware. This implies the responsibility to not break the system.

There should be some User Interface that can help operator to avoid mistakes. Some of the details include:

* Existing ClusterChannelProvisioner in the system.
* Any channels/resources related for the ClusterChannelProvisioner. If there is none, then operator knows that deprovisioning is safe.
* Additionally for each channel there is a drill down for channel details and its resources. This will help the operator to understand the existing set up better.
