# Migration plan

Migration plan will be executed several times until all modules are moved from reconciler to Lifecycle Manager. Is it possible to migrate several modules at once but it is not recommended. The migration plan uses `serverless` module as an example.

The prerequisite for migration is that Lifecycle Manager is already activated in KCP for all the clusters (Kyma CRs are created in the KCP and SKR).

## Enabling module for all clusters at once

In contradiction to previous plans it is unnecessary to separate migration process into new and old clusters. The flow will be identical regardless whether the cluster was created before, during or after the migration phase.

The Reconciler is already prepared with the [code](https://github.com/tobiscr/control-plane/blob/d43ad59ec47b9815efa1683c1f3b467b2ae3a5a1/resources/kcp/charts/mothership-reconciler/values.yaml#L153) that will disable synchronisation when a specific CRD is present in the cluster. This allows for a migration without worrying about specific timing and sequences.

With the tooling and functionalities we have recently implemented the whole migration process got quite simple. In the end it narrows down to few simple steps:
- Apply the Module Template to the KCP cluster
- Migrate the cluster (create CR and enable the Module)
- Release a new Kyma version without the Module

This covers both existing and new cluster cases. The last step is required to permanently disable the old Reconciler sync, if this step is skipped Reconciler would step in as soon as the Module is disabled (and CRD is removed with it).

## Validation scenarios

The migration process should be resilient and validated against possible user actions that could cause problems, thus we need to consider several scenarios in which user manually enables the module:
- in the new clusters - OK, clusters will be created without the Module, users can enable it via Kyma CR
- in the old (not migrated) cluster before migration script is executed - OK, default CR will be applied, the Module will no longer be reconciled by the Reconciler and Lifecycle Manager will step in
- in the old cluster during migration - OK, Reconciler will pass over the reconciliation to the Lifecycle Manager, default CR will be applied
- in the old cluster after migration - OK, Reconciler will pass over the reconciliation to the Lifecycle Manager

There is only one small concern - if the user somehow enables the module before or after the migration they will get the default configuration that could possibly override their existing config, there will be no disruption in the module availability though.

# Migration example

> **NOTE:** As we want to replicate the default, "managed" behaviour of the environment, this example bases on the dual-cluster setup for Lifecycle Manager, in order to set up your own testing grounds please refer to the Lifecycle Manager documentation available in the [KLM repository](https://github.com/kyma-project/lifecycle-manager).

Prerequisites:
- The new ModuleTemplate should be applied in the KCP cluster and synced to the SKRs
- Reconciler is prepared to [ignore your Module if the CRD is detected](https://github.com/tobiscr/control-plane/blob/d43ad59ec47b9815efa1683c1f3b467b2ae3a5a1/resources/kcp/charts/mothership-reconciler/values.yaml#L153)

Steps:
1. Enable the Module in Kyma CR:
    - If the default Module CR applies to all scenarios and there is no need to make a configuration based on existing in-cluster setup:
      ```
      kyma alpha enable module {MODULE NAME} -n {KYMA_NAMESPACE} -k {KYMA_NAME} -c {MODULE_CHANNEL}
      ```
    - If script needs to be run in order to adjust a configuration before the operator picks the module up:
      ```
      kyma alpha enable module {MODULE NAME} -n {KYMA_NAMESPACE} -k {KYMA_NAME} -c {MODULE_CHANNEL} -p Ignore
      ```
      **This will not create a Module CR and it should be handled by your automation script** (read the config, prepare Module CR based on the existing setup, apply the CR).

2. Update Kyma to a version without your Module.