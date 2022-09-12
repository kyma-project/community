---
title: Modularization in Kyma
---

With Kyma’s modular approach, you can install exactly the Kyma components you want, making the installation very lightweight and adjusted to your business needs. Kyma provides three components to support the modular approach: The Module Manager, the Lifecycle Manager, and the Runtime Watcher.

## Kyma’s Modular Approach

![](assets/Modular_Approach.svg)

## 1. Kyma Module

A Kyma module contains all information required to install and run the associated components in a Kyma runtime:

* A renderable Kubernetes manifest that includes the Kubernetes resources of the module
* A default configuration for module chart resource installation
* Optionally, further information

All these assets are bundled into a single container image using the OCI image specification.

> **TIP:** Learn how to build a module(ADD_LINK).

## 2. Module Manager

The Module Manager installs, uninstalls, and manages the Kyma modules in a local (single cluster mode) or remote setup. For example, it remotely manages and synchronizes Kyma module manifest resources in the SAP Kyma Runtime.

The Module Manager retrieves the image layers from the specified image registry and processes the manifest resources by calling a rendering framework like Helm or Kustomize.

Finally, it deploys this resource along with a custom resource to track state changes. It ensures consistency both with a time-based reconciliation and by watching state changes of the deployed custom resource in Kyma runtime.

## 3. Runtime Watcher

The Runtime Watcher monitors the relevant Kyma resources of the user’s Kyma runtime for configured changes, specified by the operator in the control plane, such as:

* Expected changes, like a Kyma user providing a configuration update, which must lead to an update of the particular Kyma runtime.
* Unexpected changes, like an accidental creation, deletion, or modification, which must be reverted to recover the Kyma runtime to a healthy state.

If needed, the Runtime Watcher triggers a reconciliation of the user’s Kyma runtime by propagating the desired changes to the Module Manager and the Lifecycle Manager.

## 4. Lifecycle Manager

The Lifecycle Manager is responsible for orchestrating Kyma module operators to process their respective resources on the Kyma runtime. This is done by generating the required Kyma module custom resources, which bootstrap Kyma module processing. Furthermore, the Lifecycle Manager reconciles a Kyma custom resource for each Kyma runtime, indicating the consolidated status of all modules configured for that Kyma cluster.
