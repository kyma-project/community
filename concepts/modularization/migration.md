# Migration plan

Migration plan will be executed several times until all modules are moved from reconciler to lifecycle-manager. Is it possible to migrate several modules at once but it is not recommended. The migration plan uses `serverless` module as an example.

## Prerequisite: enable lifecycle-manager
- KEB creates Kyma resources in the control plane. Lifecycle-manager is deployed in KCP and enabled for all clusters with empty module list.

## Enable module for new clusters (first)

First step is to enable module installation with lifecycle manager for new clusters. Existing clusters are still managed by reconciler.
- Serverless operator(manager) is ready and module template is available in the fast/regular channel. 
- New Kyma version is enabled in KEB (without serverless module). New clusters do not have serverless until customer enabled it in Kyma resource.


## Update existing clusters

Migration can be stopped at any moment if any problem occurs. Not migrated clusters will still use serverless reconciler (module not enabled in KCP)

1. Configure reconciler to ignore serverless component is the module is enabled in Kyma CR (KCP or SKR).

2. Migration script will be executed for existing clusters (in several groups - not all cluster at once). The migration script actions:
- create CRD and default resource for serverless manager in the SKR (we can have different CRs for different plans/profiles if needed)
- add module to kyma resource in KCP

3. When all clusters are migrated run migration finalization script:
- add module to Kyma CR in SKR
- remove serverless module from the Kyma CRs in the KCP 


![migration](assets/migration.drawio.svg)


## Finalization - post migration steps

- Upgrade all the clusters to the new Kyma version (without serverless)
- Remove serverless reconciler from the code.

# Validation scenarios

The migration process should be resilient and validated against possible user actions that could cause problems. It is mainly about turning on and off the module that is being migrated:
- in the new cluster - OK (only operator works as module is not included in the KEB)
- in the old (not migrated) cluster before migration script is executed - almost OK (turning on is OK, but turning off will cause reconciler running in parallel with uninstallation from lifecycle manager)
- in the old cluster during migration - OK (no effect as module is also enabled in SKR)
- in the old cluster after migration - OK (new version without serverless module is configured in the reconciler)


# Challenges

Migration has one challenge: it can become unstable if the user will add and remove module using Kyma CR in SKR during migration. The proposed mitigation to enable the module in the Kyma resource in KCP doesn't eliminate the problem as users can still turn on and off module before their cluster migration starts.

Proposal: 
As the module has to be added to Kyma CR in KCP (before migration) then removed and added to Kyma CR in SKR instead (after migration) we can simplify the process by adding it only once in SKR. To avoid switching back from lifecycle-manager to reconciler it can check if the module CRD exists to skip reconciliation. Additionally lifecycle-manager could be modified to not delete those CRDs created by the migration process (maybe some label).



