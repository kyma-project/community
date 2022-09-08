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

## 2. Module Manager

The Module Manager installs and manages the Kyma components in a local or remote setup. For example, it remotely manages and synchronizes the installation in the SAP Kyma Runtime. The Module Manager retrieves the container image and processes the Kubernetes manifest by calling a rendering framework like Helm or Kustomize. Finally, it deploys the manifest in Kubernetes and monitors the outcome of the deployment.

## 3. Runtime Watcher

The Runtime Watcher monitors the relevant Kyma resources of the user’s Kyma runtime for any changes, such as:

* Expected changes, like a Kyma user providing a configuration update, which must lead to an update of the particular Kyma runtime.
* Unexpected changes, like an accidental deletion or modification, which must be reverted to recover the Kyma runtime to a healthy state.

If needed, the Runtime Watcher triggers a reconciliation of the user’s Kyma runtime by propagating the desired changes to the Module Manager and the Lifecycle Manager.

## 4. Lifecycle Manager

The Lifecycle Manager is responsible for managing the lifecycle of Kyma instances, like installing or deleting a Kyma instance from a cluster. It checks whether a CR is provided for every module, and deletes obsoletes module if they’re no longer required

> **TIP:** Learn how to build a module(ADD_LINK).
