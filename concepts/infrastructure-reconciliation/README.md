# Motivation
Reconciler is a framework to install, update and repair Kyma components in managed Kyma Runtimes. The reconciliation runs in the loop to ensure that all components are properly configured, up and running. This principle can be extended to cover also infrastructure components like virtual networks or kubernetes clusters, but requires some adjustments in the reconciler architecture. 

# New use cases for reconciler

- create VNET for Kyma Runtime 
- create/reconcile shoot cluster for Kyma Runtime 
- produce kubeconfig for provisioned clusters (can be rotated)
- support bring your own cluster models (BYOC) - cluster is ready and we got cluster-admin kubeconfig 
- trigger reconciliation on configuration or cluster state change

# Proposed changes

## Cluster Inventory as a separate service

Cluster Inventory is the internal component responsible for reconciliation persistence. It keeps cluster information, configuration and reconciliation status. Component reconcilers do not have access to the inventory and cannot modify it directly. One option would be to extend action context passed to the reconciler to include more data, but the better option would be to pass only information how to access cluster inventory entry (url/id) and allow components reconciler to query and update it.

## Watches and triggers
Reconciliation is currently time based (triggered in regular intervals). It is known limitation. In some cases reconciliation should be triggered immediately after some events, like:
- configuration change (e.g. network reconciler can add vnet name to the cluster inventory what should trigger shoot reconciler)
- resource in the cluster modified (e.g. connectivity service instance created in the cluster should trigger connectivity proxy reconciler)


## No sequences or dependency management
Reconciler tries to manage dependencies in the simple way. The prerequisites group installed first and then other components. Introducing infrastructure reconciler will make dependency management more complex. Moreover, dependencies are valid mainly for the first installation phase, and later on are not so important (in the second reconciliation eventing reconciler can be invoked before network reconciler). It should be responsibility of component reconciler to figure out if all dependencies exist and the reconciliation can be completed. If dependencies are not ready component reconciler should return the information that it is not ready and wait for the next trigger (time or event).

# High Level Architecture

![](assets/infrastructure-reconciliation.svg)

Mothership reconciler should provide common functionality related to watching changes and triggering component reconcilers in response to events or with time based schedule. Such aspects like queueing reconciliation requests, consolidating them, updating reconciliation status, retries and back-off strategy should be implemented once (in the mothership). 

Worker pool implementation has to be revisited. Mothership should not manage dependencies - component reconciler should be written in the way that can recognize missing dependencies and can fail fast. The next reconciliation invocation should happen when the configuration (cluster inventory) is updated (dependency)

Cluster Inventory should be extracted from the mothership and should have more flexible content. Component reconcilers should be able to read and update the content (TODO - access control). 