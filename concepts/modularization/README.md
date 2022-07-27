
- [Motivation](#motivation)
- [Dependencies between components](#dependencies-between-components)
- [Release channels](#release-channels)
- [Component packaging and versioning](#component-packaging-and-versioning)
  - [Example](#example)
- [Manifest operator](#manifest-operator)
- [Component descriptor](#component-descriptor)
  - [OCM](#ocm)
  - [Operator bundle from Operator Lifecycle Manager (OLM)](#operator-bundle-from-operator-lifecycle-manager-olm)
  - [Own solution](#own-solution)
- [Component submission process](#component-submission-process)
- [Operator based component management](#operator-based-component-management)
  - [Regular (local) component operators](#regular-local-component-operators)
  - [Central component operators](#central-component-operators)
  - [Central vs local operator](#central-vs-local-operator)
- [Example local and central operators](#example-local-and-central-operators)
- [FAQ](#faq)
  - [Do we still release Kyma? What is a Kyma release?](#do-we-still-release-kyma-what-is-a-kyma-release)
  - [Can I still use `kyma deploy` command to install Kyma in my cluster?](#can-i-still-use-kyma-deploy-command-to-install-kyma-in-my-cluster)
  - [I have a simple component with helm chart. Why do I need an operator?](#i-have-a-simple-component-with-helm-chart-why-do-i-need-an-operator)
  - [I don't know how to write the operator. Can I use some generic operator for installing my chart?](#i-dont-know-how-to-write-the-operator-can-i-use-some-generic-operator-for-installing-my-chart)
  - [Why should I provide a central operator](#why-should-i-provide-a-central-operator)
  - [How to roll out a new module version in phases?](#how-to-roll-out-a-new-module-version-in-phases)
  - [Can I run multiple versions of central operator](#can-i-run-multiple-versions-of-central-operator)



# Motivation
Kyma provides Kubernetes building blocks. It should be easy to pick only those that are needed for the job and it should be easy to add new blocks to extend Kyma features. With the growing number of components, it is not possible to always install them all anymore. 

With the growing number of components, it is hard to deliver features and fixes quickly and efficiently. Changes in manifests require a new release of Kyma. Operators (reconcilers) are tightly coupled and must be released together. In most cases, new component releases don't involve any API changes and could be delivered in a few minutes. 

# Dependencies between components
Components can depend only on core Kubernetes API or API extensions introduced by other components. Component operators are responsible for checking if the required APIs are available and react properly to the missing dependencies by reducing functionality or even reporting errors. Component owners are responsible for integration tests with their dependencies (with all versions supported in official release channels).

**Example:**

If the API you require is not available you should fail (e.g. core kubernetes API or istio virtual service). If your component can work without the API, but some features are not available (e.g. service monitor from monitoring) you should just skip it and continue to deploy other component resources. 

# Release channels
Release channels let customers balance between available features and stability. The number of channels and their characteristics can be established later. Usually, projects introduce between 2 and 4 channels. Some examples:
- [GKE](https://cloud.google.com/kubernetes-engine/docs/concepts/release-channels)
- [Nextcloud](https://nextcloud.com/release-channels/)

TO DO:
- how release channels are implemented (one config file or folder in some repository / config maps)
- what component versions are included in the channel 
- governance model - requirements for promoting version to the stable channel (hotfix should be possible anyway)

# Component packaging and versioning
Kyma ecosystem produces several artifacts that can be deployed in the central control plane (KEB + operators) and in the target Kubernetes cluster. Versioning strategy should address pushing changes for all these artifacts in an unambiguous way with full traceability. Identified artifacts for each component
- Operator CRD (contains mainly overrides that can be set by customer or SRE for component installation)
- Operator deployment (yaml/helm to deploy component operator)
- Operator image (docker image in gcr)
- Component CRDs ([installation/resources/crds](https://github.com/kyma-project/kyma/tree/main/installation/resources/crds))
- Component deployment ([resources](https://github.com/kyma-project/kyma/tree/main/resources))
- Component images (docker images in gcr)

Versioning of component resources could be achieved by packaging component CRDs and charts into component operator binary (or container image). This way released operator would contain CRDs and charts of its components in the local filesystem. 
The image could be signed and we can ensure the integrity of component deployment easily. 

![](assets/modularization.drawio.svg)

## Example

If we migrate the eventing component to the proposed structure it would look like this:
- github.com/kyma-project/eventing-operator repository
    - charts (copied from kyma/resources/eventing)
    - crds (copied from kyma/installation/resources
    - operator source code (inspired by kyma-incubator/reconciler/pkg/reconciler/instances/eventing importing kyma-project/manifest-operator-lib to do the common tasks like rendering and deploying charts)
- github.com/kyma-project/eventing-controller repository (moved from kyma/components/eventing-controller)
- github.com/kyma-project/event-publisher-proxy repository (moved from kyma/components/event-publisher-proxy)

New images of our own components (eventing-controller, event-publisher-proxy) would require changes in charts inside eventing operator. Also changes in nats/jetstream would require chart updates.


# Manifest operator
Some components do not require any custom operator to install or upgrade (e.g. api-gateway, logging, monitoring) and use a base reconciler. With the operator approach, this task could be completed by a generic manifest operator. Custom resource for the manifest operator would contain information about chart location and overlay values. A single operator for multiple components can have some benefits in the in-cluster mode (better resource utilization) but would introduce challenges related to independent releases of charts and the manifest operator itself. Therefore a recommendation is that **all components will always provide the operator**. The manifest operator can be used as a template with a placeholder for your charts and default implementation (using the manifest library).

# Component descriptor
## OCM

[OCM/CNUDIE](https://github.com/gardener/component-spec) stands for Open Component Descriptor and is used by Gardener. OCM intends to solve the problem of addressing, identifying, and accessing artefacts for software components, relative to an arbitrary component repository. By that, it also enables the transport of software components between component repositories. 

## Operator bundle from Operator Lifecycle Manager (OLM)

[Operator bundle](https://olm.operatorframework.io/docs/tasks/creating-operator-bundle/#operator-bundle) is a container image that stores Kubernetes manifests and metadata associated with an operator. A bundle is meant to represent a specific version of an operator on cluster.
Operator bundle contains ClusterServiceVersion resource that describes operator version and installation descriptor:
https://olm.operatorframework.io/docs/tasks/creating-operator-manifests/#writing-your-operator-manifests

OLM provides also abstraction for operator catalog with release channels. What can solve the problem with modules discovery for users.

## Own solution
OCM is a simple descriptor and requires additional structures to apply it for installing, configuring and updating Kyma modules (operators). OLM addresses release channels and operator catalogs but we need to adjust it to the model with operators running centrally and remote subscriptions. We need to decide if we build our own solution or reuse existing software. 
We need the representation of the component in the control plane cluster. The idea is to create own custom resource that can have OCM component descriptor embedded. The ModuleTemplate custom resource would contain:
- release channel
- CRD for the module (the one managed by module operator)
- deployment of the operator (Kubernetes manifests / helm chart)


# Component submission process
To submit a new component to Kyma you need to prepare a component descriptor (with all related resources). 

Component validation should be automated by governance jobs that check different aspects like:
- operator CRD validation (e.g. status format)
- exposing metrics
- proper logging format
- ...

Governance jobs should be owned by teams (no shared responsibility)

Apart from technical quality, Kyma components should fulfill other standards (24/7 support, documenting micro-deliveries, etc)

# Operator based component management

`kyma-operator` is a meta operator responsible for installing and updating component operators. It is similar to [Operator Lifecycle Manger](https://olm.operatorframework.io/), but it can work in 2 modes: 
- central mode (managing thousands of clusters) 
- single cluster mode (local installation)

## Regular (local) component operators
Each Kyma component has to provide the operator that manages the component lifecycle (installing, updating, configuring, deleting). The component provider (team) has to prepare one custom resource (ComponentDescriptor) that specifies how to install the operator (deployment yaml) and how to configure it (operator CRD, default values). 
The complexity of managing installation in many clusters or using configured release channels is handled by kyma-operator and component providers can ship regular kubernetes operators. Such regular operators are installed in the target cluster (where components should be installed). 

## Central component operators
`kyma-operator` can also work with other component operators deployed centrally. In this case, `kyma-operator` creates resources for component operators in the central cluster. The Reconciliation logic of the central component operator can find a remote cluster kubeconfig by convention or reference (e.g. annotation). 

## Central vs local operator

With Kyma 2.0 we moved component installation (`kyma-installer`) from the cluster to the central place (reconciler). Each component was installed by the dedicated reconciler from the Kyma control plane. For some components, it perfectly makes sense (they interact mainly with other central systems to enable some integrations), but for some components that interact only with internal cluster resources, such a setup is suboptimal. The local Kubernetes operator designed and implemented accordingly to the best practices should have minimal resource requirements (few MB of memory and few mCPU). Therefore for the cases where you don't need access to central services or external resources it is recommended to use a regular operator (not central). Such an operator is much easier to develop and maintain for the component team, as an operator lifecycle management is still handled centrally by `kyma-operator` (together with `manifest-operator`).


# Example local and central operators
In this example 2 components are defined:
- central operator for compass integration
- local operator for eventing

![](assets/operators-example.drawio.svg)

# FAQ

## Do we still release Kyma? What is a Kyma release?
 
Kyma release contains `kyma-operator` in the given version and component descriptors for release channels. As we can continuously release components and upgrade their versions in the release channels the Kyma release is not so important anymore, but it can be useful for the open-source community.

## Can I still use `kyma deploy` command to install Kyma in my cluster?

Yes, but under the hood, `kyma-operator` will be used to install component operators. It is not decided yet if kyma CLI will install `kyma-operator` in the cluster or will contain `kyma-operator` code and run it locally.

## I have a simple component with helm chart. Why do I need an operator?

With the operator you can fully control your component lifecycle and ensure that your component will be reconciled  to the desired state (watch component config and actual state).

## I don't know how to write the operator. Can I use some generic operator for installing my chart?

Yes. You can use [Operator SDK - Helm](https://sdk.operatorframework.io/docs/building-operators/helm/) to generate it from your charts. You can create Helm based operator in a few minutes. If you want to better control the operator logic using [Operator SDK - Go](https://sdk.operatorframework.io/docs/building-operators/golang/) or [Kubebuilder](https://book.kubebuilder.io/)

## Why should I provide a central operator

Consider providing central operator when:
- you deal with resources outside of Kyma cluster
- you need access to external systems/resources with powerful credentials (that cannot be stored in the Kyma cluster)

## How to roll out a new module version in phases?

Use release channels to push the new version in the rapid channel first and after some time you can push that version to the stable channel. Release channels are flexible and if you need to test the new version only on 1 cluster you can create a new release channel and assign only one cluster to that channel. 

## Can I run multiple versions of central operator

Yes. But you have to ensure that each module instance described by your module custom resource is reconciled (managed) by a single operator to avoid concurrent updates and unpredictable outcomes. You can achieve that by marking your custom resources with labels pointing to release channels in your module template. Then you can deploy one central module per release channel and update them independently.
