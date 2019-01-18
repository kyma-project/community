# Component folder consolidation

## Introduction
The [components](https://github.com/kyma-project/kyma/tree/master/components) folder of the _kyma-project_ contain all Kyma components which are not related to the Kyma console. A Kyma component is a project based on source code or scripts, always resulting in docker images which are then referenced in Kyma modules/charts.
All subfolders in the _components_ directory define one component.

## Problem

Every subfolder/component uses a different naming style, for example _ui-api-layer_ vs _metadata-service_.

With that:
- Overall it is not easy to decide which convention to follow when creating a new component. What is the difference between a layer and a service?
- Not always it is possible to derive the related kyma module or domain, think of _configurations-generator_. Can you derive the domain or the usage in a module?
- Not always it is possible to derive the nature or kind of component, think of _environments_. Is it serving an API or needs access to the K8S API Server?

## Goal

Propose a well-defined vocabular for components and introduce a naming pattern baed on that.

## Considerations

After evaluation of all the components in the Kyma repository, it turned out that they are either exposing an API or are dealing with the Kubernetes API Server or just proxying/aggregating a component as in ui-api-layer and api-server-proxy.
A separation of the components into the two main categories _API_ and _K8S_ sounds meaningful as it will imply already a lot of technical aspects like:
- Acccess to API Server required (Service Account, Role Binding, ..)
- Port exposure required (Istio routing)
- Public securly accessible (Authentication/Authorization)
- Runtime workload or management workload only (auto-scaling)
- ...

For the standard Business API specific components it sounds feasible to name them **service** as they are serving an API and usually require to prefix the API with the domain it is bounded to, like _application-connector-service_. Here, the _api-ui-layer_ should be counted in as well, as its main intention is to expose a public business API requiring a security model in place (even if it is accessing the API Server). Besides that, there is a special category of API which is implementing the _OpenServiceBrokerAPI_. As we have multiple of them, we should name them accordingly **broker**.

The Kubernetes API specific components are usually following one of the well-known kubernetes patterns like **controller** and **operator**. Furthermore, there is currently one component which is a **job** not fitting into any other pattern. The components proxying another component are not fitting well into one of the discussed categories as well, but for now they are proxying only K8S specific components, so we could categorize them the same way and open up a new sub-category **proxy**.

## Proposal

### Categories and Naming Pattern

In the follwoing, all proposed component categories are getting summarized.
The proposed naming patterns for the component folders are defined per sub-category and will contain the sub-category name itself as a suffix. The main category is not included as it is implied.

Kubernetes API specific:
- **controller** - is a [Kubernetes Controller](https://kubernetes.io/docs/concepts/extend-kubernetes/extend-cluster/) named by the primary kubernetes resource which it is controlling, like _api-controller_.
- **operator** - is a [Kubernetes Operator](https://coreos.com/operators/) named by the module it is operating, like _application-operator_.
- **job** - is a [Kubernetes Job](https://kubernetes.io/docs/tasks/job/), performing a task once or periodic, named by the task it is doing, like _istio-patch-job_
- **proxy** - is proxying an existing component usually introducing a security model for the proxied component, named by the component, like _api-server-proxy_.

Business API specific:
- **service** - is serving an HTTP/S-based API, usually exposed securely to the public, named by the domain and API it is serving, like _application-connector-service_.
- **broker** - is implementing the OpenServiceBroker API, named by its provider it is integrating into, like _azure-broker_.

### Concrete Action Items

- Have a readme in the _components_ folder explaining the naming convention for components
- Link the readme in the developer guide
- All folders will stay in the root components folder but will be renamed according to the following table:

|type| old folder | new folder | action required |
|------------|------------|-----------------|----|
|**-broker**|application-broker|application-broker| - |
| |helm-broker`*`|helm-broker| - |
|**-controller**|api-controller|api-controller| - |
| |binding-usage-controller|service-binding-usage-controller| yes |
| |connection-token-handler|request-token-controller| yes|
| |namespace-controller|namespace-controller|no|
| |idppreset|idppreset-controller|yes|
| |event-bus`**`/event-bus-sv|event-activation-controller|yes|
| |event-bus`**`/event-bus-push|subscription-controller|yes|
|**-operator**|application-operator|application-operator| - |
| |installer|kyma-operator| yes|
|**-service**|configurations-generator|iam-kubeconfig-service|yes|
| |connector-service|application-connector-service|yes|
| |application-registry| application-registration-service|yes|
| |event-service|application-event-service| yes|
| |application-proxy|application-proxy-service|yes|
| |event-bus`**`/event-bus-publish|event-publish-service|yes|
| |ui-api-layer|console-backend-service|yes|
|**-proxy**|apiserver-proxy|apiserver-proxy| - |
| |k8s-dashboard-proxy|k8s-dashboard-proxy| - |
|**-job**|istio-kyma-patch|istio-patch-job|yes|

`*` should be splitted into the helm-broker and tooling for bundle repositories outside of kyma repo

`**` should be splitted into the different components, the common parts not specific to eventing should be moved into the kyma common parts, otherwise they should be duplicated

## Addition

Going the extra mile we should also have a look at the `tools` folder. Currently it is not specified what a tool is. To get an answer let's define first a module and a component:

- `module` - A Kyma module is a helm chart installable by the installer located in the _kyma/resources_ folder. A module can be optional for a Kyma installation. An example for a module is `service-catalog`, a module consisting mainly of 3-party images but also levaring some kyma components like the _binding-usage-controller_. 
- `component` - A Kyma component is any pod/container/image which is deployed with a Kyma module to Kyma to provide its functionality. A component is made of sources located in the _kyma/components_ folder.

Any image not fitting into the definition is not providing functionality within Kyma and with that is a tool. For example the watch-pods container which is used for testing/monitoring. Without it Kyma is working as well as with it: it is not a component.

Having that definition in mind, there are tools listed in thte tools folder which are more a component and should be moved:

|type| old folder | new folder | action required | description |
|----|------------|------------|-----------------|-------------|
|**-job**|tools/etcd-backup|components/etcd-backup-job|yes| used in service catalog only as a cronjob |
| |tools/etcd-tls-setup|components/etcd-tls-setup-job|yes| used in service catalog as a job, misused in ark module (should be removed) and mentioned in helm-broker (should be removed) |
|**-configurer**|tools/alpine-net|components/alpine-net-configurer|yes| base image having net-utils installed, used only as init container for dependency checks|
| |tools/ark-plugins|components/ark-plugins-configurer|yes| init-container for configuring ark|
| |tools/static-users-generator|components/dex-static-user-configurer|yes| init-container for configuring dex with generated user credentials|

Furthermore, there are tools used in the release process or test infrastructure which should be moved into the _test-infra_ repository:

|old folder | new folder | action required | description |
|-----------|------------|-----------------|-------------|
|kyma:tools/changelog-generator|test-infra:changelog-generator|yes| used in release process only|
|kyma:tools/watch-pods|test-infra:watch-pods|yes| used for integration test execution |
|kyma:tools/stability-checker|test-infra:stability-checker|yes| used for continuous functional testing against a running cluster|

There are 3 tools left for now which should be moved somewhere else having the goal in mind to have no general purpose tools folder at all, proposals are welcome:

|old folder | new folder | action required | description |
|-----------|------------|-----------------|-------------|
|kyma:tools/docsbuilder|-|-|used to create docker images containing the documentation. It is a temporary solution and will be removed soon|
|kyma:gcp-broker-provider|?|?| script for provisioning a gcp broker into a namespace, main usage is as helm broker bundle, so maybe have it in the bundles repository?|
|kyma:kyma-installer|?|?|the ultimate installer, should be released tgether with the kyma modules/charts always, maybe should be closer to the current _resources_ folder?
