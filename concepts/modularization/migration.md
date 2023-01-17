# Migration plan

Migration plan will be executed several times until all modules are moved from reconciler to lifecycle-manager. Is it possible to migrate several modules at once but it is not recommended. The migration plan uses `serverless` module as an example.

## Prerequisite: enable lifecycle-manager
- KEB creates Kyma resources in the control plane. Lifecycle-manager is deployed in KCP and enabled for all clusters with empty module list.

## Enable module for new clusters (first)

First step is to enable module installation with lifecycle manager for new clusters. Existing clusters are still managed by reconciler.
- Serverless operator(manager) is ready and module template is available in the fast/regular channel. 
- New Kyma version is enabled in KEB (without serverless module). New clusters do not have serverless until customer enabled it in Kyma resource.

## Update existing clusters

Migration can be stopped at any moment if any problem occurs. Not migrated clusters will still use serverless reconciler (old Kyma version with serverless)

1. Upgrade the clusters to the new Kyma version (without serverless)

2. Create CRD and default resource for serverless manager in the SKR (we can have different CRs for different plans/profiles if needed)


![migration](assets/migration.drawio.svg)

# Validation scenarios

The migration process should be resilient and validated against possible user actions that could cause problems. It is mainly about turning on and off the module that is being migrated:
- in the new cluster - OK (reconciler is already disabled for serverless)
- in the old (not migrated) cluster before migration script is executed - almost OK (if user adds serverless module manually it will work in parallel with reconciler)
- in the old cluster during migration - OK (reconciler is already disabled for serverless)
- in the old cluster after migration - OK (reconciler is already disabled for serverless)


# Challenges

Migration has one challenge: it can become unstable if the user will add and remove module using Kyma CR in SKR when module template is available but before migration has been started. To minimize the risk we can disable module template propagation for old clusters until migration is over.
