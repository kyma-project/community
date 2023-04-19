---
title: How to enable and disable a Kyma module
---

## Overview

Your cluster comes with the Kyma custom resource (CR) already installed. It collects all metadata about the cluster, such as enabled modules, their statuses, or synchronization, using Lifecycle Manager. Lifecycle Manager uses `moduletemplate.yaml` to enable or disable modules on your cluster. 

## Procedure

<div tabs name="steps" group="enable-module">
  <details>
  <summary label="cli">
  Kyma CLI
  </summary>

Use kubectl to check which modules are available on your cluster. Run: 
   ```bash
   kubectl get ModuleTemplates -A
   ```

Use Kyma CLI to enable a module on your cluster in the release channel of your choice. Run: 

   ```bash
   kyma alpha enable module {MODULE_NAME} --channel {CHANNEL_NAME} --wait
   ```

Similarly, to disable a module, run: 

   ```bash
   kyma alpha disable module {MODULE_NAME}
   ``` 

</details>
<details>
<summary label= Kyma Dashboard>
Kyma Dashboard
</summary>

1. Go to the `kyma-system` Namespace.
2. In the **Kyma** section, choose the **Kyma** resource.
3. Click on the name of your Kyma instance (`default-kyma`) and click **Edit**.
4. In the **Modules** section, click **Add**.
5. Choose the name of your module.
6. Choose the available channel.
7. Click **Update** and then **Force update**.
The operation was successful if the module status changed to `READY`.
</details>
</div>

To configure your module, use the module CR. For the configuration options, check the module configuration page {LINK}. 