# Migration plan

Migration plan will be executed several times until all modules are moved from reconciler to lifecycle-manager. Is it possible to migrate several modules at once but it is not recommended. The migration plan uses `serverless` module as an example.

The prerequisite for migration is that lifecycle-manager is already activated in KCP for all the clusters (Kyma CRs are created in the KCP and SKR). Migration is split into 2 phases: enabling module for the new clusters and update existing clusters.

![migration](assets/migration.drawio.svg)

## Enable module for new clusters (first)

First step is to enable module installation with lifecycle manager for new clusters. Existing clusters are still managed by reconciler.
- Disable module template replication for existing clusters (moduleCatalog=false). This prevents users to add serverless module before it is disabled in the reconciler.
- New Kyma version is enabled in KEB (without serverless module). New clusters do not have serverless until customer enabled it in Kyma resource.
- Serverless operator(manager) is ready and module template is available in the fast/regular channel. 

## Update existing clusters

Migration can be stopped at any moment if any problem occurs. Not migrated clusters will still use serverless reconciler (old Kyma version with serverless)

- Upgrade the clusters to the new Kyma version (without serverless) - from this point reconciler doesn't touch serverless resources.
- Add module to Kyma CR in SKR with option to not copy default CR. Lifecycle manager will install serverless manager without starting module installation. Default CR creation is omitted for the case when we need different setting for different clusters (eg. trial / production). If default CR can be applied for all the clusters then module copy can stay enabled and the next step is not required.
- Create CRD  and default CR for serverless manager in the SKR. CRD creation is optional - it can be skipped, but then migration script has to wait until liefecycle-manager installs it. Entire step can be skipped if the default module CR can be applied for all the clusters.

# Validation scenarios

The migration process should be resilient and validated against possible user actions that could cause problems. It is mainly about turning on and off the module that is being migrated:
- in the new cluster - OK (reconciler is already disabled for serverless)
- in the old (not migrated) cluster before migration script is executed - almost OK (if user adds serverless module manually it will work in parallel with reconciler)
- in the old cluster during migration - OK (reconciler is already disabled for serverless)
- in the old cluster after migration - OK (reconciler is already disabled for serverless)

Migration validation discovers only one challenge: it can become unstable if the user will add and remove module using Kyma CR in SKR before reconciler is disabled for that cluster. It is mitigated by disabling module template propagation for old clusters until migration is over.

# Migration example

> **NOTE:** As we want to replicate the default, "managed" behaviour of the environment, this example bases on the dual-cluster setup for Lifecycle Manager, in order to set up your own testing grounds please refer to the Lifecycle Manager documentation available in the [KLM repository](https://github.com/kyma-project/lifecycle-manager).

1. **KCP:** Disable the module catalog sync for the cluster - this will block any updates to the ModuleTemplate list present in the SKR cluster. The user will not see any new modules appearing to be installed in their cluster, the rest of the modules that were already there will still behave as normal.

    > **WARNING:** This does not mean that the module sync will be disabled, if the user somehow has the name of the module before or during the migration and enables it in Kyma CR manually it will be installed and may influence the migration process.

    ```shell
    kubectl patch kyma -n {KYMA_NAMESPACE} {KYMA_NAME} -p '{"spec":{"sync":{"moduleCatalog":false}}}' --type="merge"
    ```

2. **KCP:** Create the ModuleTemplate in the cluster.

    > **NOTE:** All the Kymas with `moduleCatalog` set to `true` will get this ModuleTemplate, make sure that every Kyma CR has this field set to proper value.

    ```shell
    cat <<EOF | kubectl apply -f -
    kubectl apply -f https://github.com/kyma-project/keda-manager/releases/download/v0.0.3/moduletemplate.yaml
    EOF
    ```

    also, we need to patch the `target` to `remote` as it is set by default to the `control-plane` value.

    ```shell
    k patch -n kcp-system moduletemplates.operator.kyma-project.io moduletemplate-keda -p '{"spec":{"target":"remote"}}' --type="merge"
    ```

3. **SKR:** Apply the CRD on the cluster, we can use the [keda](https://github.com/kyma-project/keda-manager) module.

    > **Note:** This will disable the reconcilation for the component in the old Reconciler, until the module is enabled in the Kyma CR it will stay in an unmanaged state.

4. **SKR:** Create the default Module CR in the cluster. Make sure that the metadata (name, namespace) is aligned with the Module Template and the namespaces matches the one in the ModuleTemplate or in the Kyma CR.

    ```shell
    cat <<EOF | kubectl apply -f -
    apiVersion: operator.kyma-project.io/v1alpha1
    kind: Keda
    metadata:
      name: keda-sample
      namespace: kcp-system
    spec:
      logging:
        operator:
          level: info
      resources:
        metricServer:
          limits:
            cpu: 500m
            memory: 500Mi
          requests:
            cpu: 500m
            memory: 500Mi
        operator:
          limits:
            cpu: 500m
            memory: 500Mi
          requests:
            cpu: 500m
            memory: 500Mi
    EOF
    ```

5. **SKR:** Enable the module in Kyma, this will pick up the CR that was created in the previous step and use it as the configuration for the module. This will install the module operator in the cluster which should pick up the resources that Reconciler managed previously.
    ```shell
    kyma alpha enable module keda -n kcp-system -k kyma-sample -c beta
    ```

6. **KCP:** Reenable the Module Catalog sync for Kyma.
    ```shell
    kubectl patch kyma -n {KYMA_NAMESPACE} {KYMA_NAME} -p '{"spec":{"sync":{"moduleCatalog":true}}}' --type="merge"
    ```

After the last step everything should be back to the previous state with the exception that the migrated module will be reconciled by the module operator and not the Reconciler.