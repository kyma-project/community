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

Use Kyma CLI to check which modules are available on your cluster. Run: 
   ```bash
   kyma alpha list module
   ```

Use Kyma CLI to enable a module on your cluster in the release channel of your choice. Run: 

   ```bash
   kyma alpha enable module {MODULE_NAME} --channel {CHANNEL_NAME} --wait
   ```

You should see the following message:

```bash
   - Successfully connected to cluster
   - Modules patched!
   ```

Similarly, to disable a module, run: 

   ```bash
   kyma alpha disable module {MODULE_NAME}
   ``` 
You should see the same message as the one displayed when you enable a module.

</details>
<details>
<summary label= Kyma Dashboard>
Kyma Dashboard
</summary>

Follow these steps to enable a Kyma module in Kyma Dashboard:
1. Go to the `kyma-system` Namespace.
2. Choose the **Kyma** resource from the **Kyma** section.
3. Click on the name of your Kyma instance (`default-kyma`) and click **Edit**.
4. Click **Add** in the **Modules** section.
5. Choose the name of your module from the dropdown menu.
6. Choose the available channel.
7. Click **Update** and then **Force update**.
The operation was successful if the module Status changed to `READY`.

To disable a module, edit your Kyma instance and click on the trash icon next to your module, then force update the changes. Your module should disappear from the Module list.
</details>
</div>

To configure your module, use the module CR. For the configuration options, check the module configuration page {LINK}. 