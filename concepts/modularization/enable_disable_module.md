---
title: How to enable and disable a Kyma module
---

## Overview

Your cluster comes with the Kyma custom resource (CR) already installed. It collects all metadata about the cluster, such as enabled modules, their statuses, or synchronization, using Lifecycle Manager. Lifecycle Manager uses `moduletemplate.yaml` to enable or disable modules on your cluster. 

## Procedure

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

To configure your module, use the module CR. For the configuration options, check the module configuration page {LINK}. 