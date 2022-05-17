# Modularization

Kyma provides Kubernetes building blocks. It should be easy to pick only those that are needed for the job and it should be easy to add new blocks to extend Kyma features. With the growing number of components, it is not possible to always install them all anymore. 

- [Independent releases of modules](#independent-releases-of-modules)
- [Dependencies between components](#dependencies-between-components)
- [Release channels](#release-channels)
- [Versioning](#versioning)
- [Manifest operator](#manifest-operator)

# Independent releases of modules
With the growing number of components, it is hard to deliver features and fixes quickly and efficiently. Changes in manifests require a new release of Kyma. Operators (reconcilers) are tightly coupled and must be released together. In most cases, new component releases don't involve any API changes and could be delivered in a few minutes. 


# Dependencies between components
Components can depend only on core Kubernetes API or API extensions introduced by other components. Component operators are responsible for checking if the required APIs are available and react properly to the missing dependencies by reducing functionality or even reporting errors. Component owners are responsible for integration tests with their dependencies (with all versions supported in official release channels).

# Release channels
Release channels let customers balance between available features and stability. The number of channels and their characteristics can be established later. Usually, projects introduce between 2 and 4 channels. Some examples:
- [GKE](https://cloud.google.com/kubernetes-engine/docs/concepts/release-channels)
- [Nextcloud](https://nextcloud.com/release-channels/)

TO DO:
- how release channels are implemented (one config file or folder in some repository / config maps)
- what component versions are included in the channel 
- governance model - requirements for promoting version to the stable channel (hotfix should be possible anyway)

# Versioning
Kyma ecosystem produces several artifacts that can be deployed in the central control plane (KEB + operators) and in the target Kubernetes cluster. Versioning strategy should address pushing changes for all these artifacts in an unambiguous way with full traceability. Identified artifacts for each component
- Operator CRD ([api/config/crd](https://github.com/kyma-project/manifest-operator/tree/main/api/config/crd))
- Operator deployment ([operator/config/manager](https://github.com/kyma-project/manifest-operator/blob/main/operator/config/manager/manager.yaml))
- Operator image (docker image in gcr)
- Component CRDs ([installation/resources/crds](https://github.com/kyma-project/kyma/tree/main/installation/resources/crds))
- Component deployment ([resources](https://github.com/kyma-project/kyma/tree/main/resources))
- Component images (docker images in gcr)

Component operators should be deployed continuously. Operators should support all versions that are currently available in all release channels. It is up to the component owner to decide how they manage different component versions inside operator (operator per version or single operator supporting multiple versions).

## Simple versioning

Simple versioning of component resources could be achieved by packaging component CRDs and charts into component operator binary (or container image). This way released operator would contain CRDs and charts of its components in the local filesystem. 
The image could be signed and we can ensure the integrity of component deployment easily. 

![](assets/modularization.drawio.svg)

**Example:**

If we migrate eventing component to the proposed structure it would look like this:
- github.com/kyma-project/eventing-operator repository
    - charts (copied from kyma/resources/eventing)
    - crds (copied from kyma/installation/resources
    - operator source code (inspired by kyma-incubator/reconciler/pkg/reconciler/instances/eventing importing kyma-project/manifest-operator-lib to do the common tasks like rendering and deploying charts)
- github.com/kyma-project/eventing-controller repository (moved from kyma/components/eventing-controller)
- github.com/kyma-project/event-publisher-proxy repository (moved from kyma/components/event-publisher-proxy)

New images of our own components (eventing-controller, event-publisher-proxy) would require changes in charts inside eventing operator. Also changes in nats/jetstream would require chart updates.


# Manifest operator
Some components do not require any custom operator to install or upgrade (e.g. api-gateway, logging, monitoring) and use base reconciler. With the operator approach, this task could be completed by a generic manifest operator. Custom resource for manifest operator would contain information about chart location and overlay values. A single operator for multiple components can have some benefits in the in-cluster mode (better resource utilization) but would introduce challenges related to independent releases of charts and the manifest operator itself. Therefore a recommendation is that all components will always provide the operator. Manifest operator can be used as a template with a placeholder for your charts and default implementation (using manifest library).
# Component submission process
To submit a new component to Kyma you need to prepare:
- Component operator custom resource definition (CRD in YAML)
- Operator deployment and component config (YAML)

Component validation should be automated by governance jobs that check different aspects like:
- component status format 
- support for release channels
- exposing metrics
- proper logging format
- ...

Governance jobs should be owned by teams (no shared responsibility)

Apart from technical quality, Kyma components should fulfill other standards (24/7 support, documenting micro-deliveries, etc)

TODO:
- verify if we can use [OCM/CNUDIE](https://github.com/gardener/component-spec) for component submission


