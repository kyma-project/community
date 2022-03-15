# Reconciler infrastructure changes proposal

## Motivation
Reconciler is a framework to install, update, and repair Kyma components in managed Kyma Runtimes. The reconciliation runs in the loop to ensure that all components are up and running and properly configured. This principle can be extended to cover also infrastructure components, such as virtual networks (VNet) or Kubernetes clusters, but requires some adjustments in the Reconciler architecture. 

## New use cases for Reconciler

- Create VNet for Kyma Runtime 
- Create/reconcile a shoot cluster for Kyma Runtime 
- Create kubeconfig for provisioned clusters (can be rotated)
- Support bring your own cluster models (BYOC) - a cluster is ready and we get cluster-admin kubeconfig 
- Trigger reconciliation on a configuration or on a cluster state change
- Register Kyma Runtime within the CMP

## Proposed changes

### Cluster Inventory as a separate service

Cluster Inventory is the internal component responsible for reconciliation persistence. It keeps cluster information, configuration, and reconciliation status. Component Reconcilers do not have access to the inventory and cannot modify it directly. One option would be to extend the action context passed to the Reconciler to include more data, but the better option would be to pass only information on how to access Cluster Inventory entry (URL/ID) and allow Components Reconciler to query and update it.

### Reconciliation based on watchers and triggers
Reconciliation is currently time-based (triggered in regular intervals). It is a known limitation. In some cases, the reconciliation should be triggered immediately after some events, such as:
- Configuration change (e.g. Network Reconciler can add a VNet name to the Cluster Inventory, which would trigger Shoot Reconciler)
- Resource in the cluster modified (e.g. Connectivity Service instance created in the cluster would trigger Connectivity Proxy Reconciler)

Restoring managed resources accidentally (or intentionally) deleted by a user is not the main goal of watchers. They should be rather used to monitor resources such as service instances that should trigger some integration tasks in the Kyma Control Plane.

Watchers and triggers can be more selective to limit the number of unnecessary invocations. Component Reconcilers can specify which changes in the Cluster Inventory fields they would be triggered by (i.e. Istio Reconciler subscribing to the kubeconfig field). 

### No sequences or dependency management
Reconciler tries to manage dependencies in a simple way. The prerequisites group is installed first, before other components. Introducing Infrastructure Reconciler will make dependency management more complex. Moreover, dependencies are valid mainly for the first installation phase and are not so important later on (in the second reconciliation, Eventing Reconciler can be invoked before Network Reconciler). It should be the responsibility of the Component Reconciler to figure out if all dependencies exist and the reconciliation can be completed. If dependencies are not ready, Component Reconciler should return the information that it is not ready and wait for the next trigger (time or event).

## High level architecture

![](assets/infrastructure-reconciliation.svg)

Mothership Reconciler should provide common functionality related to watching changes and triggering Component Reconcilers in response to events or with the time-based schedule. Such aspects as queueing reconciliation requests, consolidating them, updating reconciliation status, retries and back-off strategy should be implemented once (in the Mothership). 

Worker pool implementation has to be revisited. Mothership should not manage dependencies - Component Reconciler should be written in a way that can recognize missing dependencies and fail fast. The next reconciliation invocation should happen when the configuration (Cluster Inventory) is updated (dependency). 

Cluster Inventory should be extracted from the Mothership and should have more flexible content. Component Reconcilers should be able to read and update the content. Access control should be also established as we would need access to Cluster Inventory also from external components (e.g. Central Management Plane).