
- [Motivation](#motivation)
- [Dependencies between components](#dependencies-between-components)
- [Release channels](#release-channels)
- [Component packaging and versioning](#component-packaging-and-versioning)
  - [Example](#example)
- [Module manager](#module-manager)
- [Component descriptor](#component-descriptor)
  - [OCM](#ocm)
  - [Operator bundle from Operator Lifecycle Manager (OLM)](#operator-bundle-from-operator-lifecycle-manager-olm)
  - [Own solution](#own-solution)
- [Component submission process](#component-submission-process)
- [Operator-based component management](#operator-based-component-management)
  - [Regular (local) component operators](#regular-local-component-operators)
  - [Central component operators](#central-component-operators)
  - [Central vs local operator](#central-vs-local-operator)
- [Example for local and central operators](#example-for-local-and-central-operators)
- [FAQ](#faq)
  - [Do we still release Kyma? What is a Kyma release?](#do-we-still-release-kyma-what-is-a-kyma-release)
  - [Can I still use the `kyma deploy` command to install Kyma in my cluster?](#can-i-still-use-the-kyma-deploy-command-to-install-kyma-in-my-cluster)
  - [I have a simple component with a Helm chart. Why do I need an operator?](#i-have-a-simple-component-with-a-helm-chart-why-do-i-need-an-operator)
  - [I don't know how to write the operator. Can I use some generic operator for installing my chart?](#i-dont-know-how-to-write-the-operator-can-i-use-some-generic-operator-for-installing-my-chart)
  - [Why should I provide a central operator?](#why-should-i-provide-a-central-operator)
  - [How to roll out a new module version in phases?](#how-to-roll-out-a-new-module-version-in-phases)
  - [Can I run multiple versions of central operator](#can-i-run-multiple-versions-of-central-operator)
  - [How do we migrate all the modules to the new concept?](#how-do-we-migrate-all-the-modules-to-the-new-concept)



# Motivation
Kyma provides Kubernetes building blocks. It should be easy to pick only those that are needed for the job and it should be easy to add new blocks to extend Kyma features. With the growing number of components, it is not possible to always install them all anymore. 

With the growing number of components, it is hard to deliver features and fixes quickly and efficiently. Changes in manifests require a new release of Kyma. Operators (reconcilers) are tightly coupled and must be released together. In most cases, new component releases don't involve any API changes and could be delivered in a few minutes. 

# Dependencies between components
Components can depend only on core Kubernetes API, or on API extensions introduced by other components. Component operators must check whether the required APIs are available and react properly to the missing dependencies by reducing functionality or even reporting errors. Component owners are responsible for integration tests with their dependencies (with all versions supported in official release channels).

**Example:**

If the API you need (like a core Kubernetes API or Istio virtual service) is not available, you should fail. If your component can work without the API, but some features are not available (for example, service monitor from monitoring), you should just skip it and continue to deploy other component resources. 

# Release channels
Release channels let customers try new modules and features early, and decide when the updates should be applied. 

![](assets/release-channels.drawio.svg)

The first use case will be modeled as the alpha channel. Modules available in the alpha channel are developed with all quality measures in place (functional correctness, security, etc.), but they might still have unstable API or be changed without keeping backward compatibility. When you use a module from the alpha channel, you won't get full SLA guarantees for that module or other modules that are affected (directly or indirectly).

The second use case (deciding when updates should be applied) will require 2 production-grade channels with a different update schedule. The fast channel will get updates as soon as they are released and have passed all quality gates. The regular channel will get updates a few days later. Customers can switch the entire cluster or a particular component to the fast channel to check if the upstream changes do not cause any issues with their workload. Changing back to the regular channel is possible, but the module version won't be downgraded - the next version has to reach the channel to be applied.

Hotfixes will be delivered to all channels immediately (TODO: how to apply a hotfix for the release that is not available in the current channel).



# Component packaging and versioning
Kyma ecosystem produces several artefacts that can be deployed in the central control plane (KEB + operators) and in the target Kubernetes cluster. Versioning strategy should address pushing changes for all these artefacts in an unambiguous way with full traceability. 

For each component, we identified the following artefacts:
- Operator CRD (contains mainly overrides that can be set by customer or SRE for component installation)
- Operator deployment (YAML or Helm to deploy component operator)
- Operator image (Docker image in GCR)
- Component CRDs ([installation/resources/crds](https://github.com/kyma-project/kyma/tree/main/installation/resources/crds))
- Component deployment ([resources](https://github.com/kyma-project/kyma/tree/main/resources))
- Component images (docker images in GCR)

Versioning of component resources could be achieved by packaging component CRDs and charts into the component operator binary (or container image). This way, the released operator would contain the CRDs and charts of its components in the local filesystem. 
The image could be signed and we can ensure the integrity of component deployment easily. 

![](assets/modularization.drawio.svg)

## Example

If we migrate the Eventing component to the proposed structure, it would look like this:
- `github.com/kyma-project/eventing-operator` repository
    - charts (copied from `kyma/resources/eventing`)
    - CRDs (copied from `kyma/installation/resources`)
    - operator source code (inspired by `kyma-incubator/reconciler/pkg/reconciler/instances/eventing` importing `kyma-project/module-manager-lib` to do the common tasks like rendering and deploying charts)
- `github.com/kyma-project/eventing-controller` repository (moved from `kyma/components/eventing-controller`)
- `github.com/kyma-project/event-publisher-proxy` repository (moved from `kyma/components/event-publisher-proxy`)

New images of our own components (`eventing-controller`, `event-publisher-proxy`) would require changes in charts inside Eventing operator. Also, changes in `nats/jetstream` would require chart updates.


# Module manager
Some components do not need any custom operator to install or upgrade (for example, API Gateway, Logging, Monitoring) and use a base reconciler. With the operator approach, this task could be completed by a generic `module-manager`. Custom resource for the `module-manager` would contain information about chart location and overlay values. A single operator for multiple components can have some benefits in the in-cluster mode (better resource utilization), but would introduce challenges related to independent releases of charts and the `module-manager` itself. Therefore a recommendation is that **all components will always provide the operator**. The `module-manager` can be used as a template with a placeholder for your charts and default implementation (using the module-manager library).

# Component descriptor
## OCM

[OCM/CNUDIE](https://github.com/gardener/component-spec) stands for Open Component Descriptor and is used by Gardener. OCM intends to solve the problem of addressing, identifying, and accessing artefacts for software components, relative to an arbitrary component repository. By that, it also enables the transport of software components between component repositories. 

## Operator bundle from Operator Lifecycle Manager (OLM)

[Operator bundle](https://olm.operatorframework.io/docs/tasks/creating-operator-bundle/#operator-bundle) is a container image that stores Kubernetes manifests and metadata associated with an operator. A bundle is meant to represent a specific version of an operator on cluster.
Operator bundle contains the ClusterServiceVersion resource that describes the operator version and installation descriptor:
`https://olm.operatorframework.io/docs/tasks/creating-operator-manifests/#writing-your-operator-manifests`

OLM also provides abstraction for the operator catalog with release channels, which can solve the problem how users discover modules.

## Own solution
OCM is a simple descriptor and requires additional structures to apply it for installing, configuring, and updating Kyma modules (operators). OLM addresses release channels and operator catalogs, but we must adjust it to the model with operators running centrally and remote subscriptions. We must decide if we build our own solution or reuse existing software. 
We need the representation of the component in the control plane cluster. The idea is to create our own custom resource that can have an OCM component descriptor embedded. The ModuleTemplate custom resource would contain the following:
- release channel
- CRD for the module (the one managed by module operator)
- deployment of the operator (Kubernetes manifests or Helm chart)


# Component submission process
To submit a new component to Kyma, you must prepare a component descriptor with all related resources. 

Component validation should be automated by governance jobs that check different aspects like:
- operator CRD validation (for example, status format)
- exposing metrics
- proper logging format
- ...

Governance jobs should be owned by teams (no shared responsibility).

Apart from technical quality, Kyma components must fullfil other standards, such as 24/7 support, documenting micro-deliveries.

# Operator-based component management

`lifecycle-manager` is a meta operator responsible for installing and updating component operators. It is similar to [Operator Lifecycle Manger](https://olm.operatorframework.io/), but it can work in two modes: 
- central mode (managing thousands of clusters) 
- single cluster mode (local installation)

## Regular (local) component operators
Each Kyma component must provide the operator that manages the component lifecycle (installing, updating, configuring, deleting). The component provider (team) must prepare one custom resource (ComponentDescriptor) that specifies how to install the operator (`deployment.yaml`) and how to configure it (operator CRD, default values). 
The complexity of managing installation in many clusters or using configured release channels is handled by `lifecycle-manager`. Component providers can ship regular Kubernetes operators. Such regular operators are installed in the target cluster where components should be installed. 

## Central component operators
`lifecycle-manager` can also work with other component operators deployed centrally. In this case, `lifecycle-manager` creates resources for component operators in the central cluster. The reconciliation logic of the central component operator can find a remote cluster kubeconfig by convention or reference (for example, annotation). 

## Central vs local operator

With Kyma 2.0, we moved the component installation (`kyma-installer`) from the cluster to a central place (reconciler). Each component was installed by the dedicated reconciler from the Kyma control plane. For some components, it perfectly makes sense because they interact mainly with other central systems to enable some integrations. But for some components that interact only with internal cluster resources, such a setup is suboptimal. The local Kubernetes operator designed and implemented according to the best practices should have minimal resource requirements (few MB of memory and few mCPU). Therefore, for the cases where you don't need access to central services or external resources, it is better to use a regular operator than a central one. A regular operator is much easier to develop and maintain for the component team, because the operator lifecycle management is still handled centrally by `lifecycle-manager` (together with `module-manager`).


# Example for local and central operators
In the following example, two components are defined:
- central operator for Compass integration
- local operator for Eventing

![](assets/operators-example.drawio.svg)

# FAQ

## Do we still release Kyma? What is a Kyma release?
 
Every Kyma release contains the `lifecycle-manager` in the given version and component descriptors for release channels. As we can continuously release components and upgrade their versions in the release channels, the Kyma release is not that important anymore, but it can be useful for the open source community.

## Can I still use the `kyma deploy` command to install Kyma in my cluster?

Yes, but under the hood, `lifecycle-manager` will be used to install component operators. It is not decided yet whether Kyma CLI will install `lifecycle-manager` in the cluster or will contain `lifecycle-manager` code and run it locally.

## I have a simple component with a Helm chart. Why do I need an operator?

With the operator, you can fully control your component lifecycle and ensure that your component are reconciled to the desired state (watch component config and actual state). Each operator comes with a custom resource that describes the module configuration and represents the module installation status. It is a way to enable users with providing chart overrides in a controlled way.

## I don't know how to write the operator. Can I use some generic operator for installing my chart?

Yes. You can use [Operator SDK - Helm](https://sdk.operatorframework.io/docs/building-operators/helm/) to generate it from your charts. You can create a Helm-based operator in a few minutes. If you want more control over the operator logic, use [Operator SDK - Go](https://sdk.operatorframework.io/docs/building-operators/golang/) or [Kubebuilder](https://book.kubebuilder.io/)

## Why should I provide a central operator?

Consider providing central operator in the following cases:
- You deal with resources outside the Kyma cluster.
- You need access to external systems or resources with powerful credentials that cannot be stored in the Kyma cluster.

## How to roll out a new module version in phases?

Use release channels to push the new version in the rapid channel first. After some time, you can push that version to the stable channel. Release channels are flexible; and if you need to test the new version only on one cluster, you can create a new release channel and assign only one cluster to that channel. 

## Can I run multiple versions of central operator

Yes. But to avoid concurrent updates and unpredictable outcomes, you must ensure that each module instance described by your module custom resource is reconciled (managed) by a single operator. You can achieve that by marking your custom resources with labels pointing to release channels in your module template. Then you can deploy one central module per release channel and update them independently.

## How do we migrate all the modules to the new concept?

Read about the initial plan in the [transition document](transition.md).