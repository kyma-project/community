# Integration of the knative based function controller

## Introduction

One of the main features of kyma is its ability to run serverless functions. Currently this feature is implemented using the open source project [kubeless](https://github.com/kubeless/kubeless). With the focus shift towards knative will be replaced with a [knative](https://github.com/knative) based function runtime controller (called knative-functions from here on).

## Problem

As long as neither [knative-serving](https://github.com/knative/serving) / [knative-build](https://github.com/knative/build) nor the knative-functions itself provide a stable API it shall be an experimental feature to use the new function controller.

Switching the runtime should provide a similar user experience as with the currently available kubeless runtime. This includes CRUD operations for the functions from within the user interface as well as monitoring / alerting and  backup / restore.

## Goal

* Clearly show problems and provide possible solutions.
* Propose next steps for integration of the new controller runtime.
* Provide a migration path from kubeless to the new runtime

## Problems

### Multiple Runtimes

Currently kubeless is installed by default as a core component into every kyma installation. Technically it is easily possible to install the knative-functions in parallel with kubeless.

#### 'Solo' Installation

Installing only one of the runtimes requires changes in the current installation approach for kubeless. kubeless is currently not a separate module. This means its installation cannot be disabled at the moment. Moving kubeless to its own module is a major task. Due to the fact that the use of kubeless will be discontinued we advise against this solution.

##### Advantages

* no unused system ressources due to unused controllers
* cleaner setup (only one runtime available)

##### Disadvantages

* kubeless has to be converted into an optional module

#### Active runtime selection

Another approach to have only one active runtime could be a feature-flag that allows the administrator to select the active runtime during installation/upgrade of kyma.

##### Advantages

* installation easy to implement
* only one function implementation will be active per cluster

##### Disadvantages

* both runtimes use system ressources
* UI has to be able to cope with both runtimes
  
#### Both runtimes active

It is also possible to install both runtimes at the same time and support functions of both flavors simultaneously

##### Advantages

* installation easiest to implement
* users can compare both runtimes

##### Disadvantages

* both runtimes use system ressources
* UI has to be able to cope with both runtimes at the same time(probably hardest version to implement)

#### Proposal

install both runtimes into the same cluster and allow the user to use both runtimes at will.

### UI changes

Even though the kubeless function crd and the knative-functions crd look very similar some changes to the UI layer are required:

#### Builds

The new function runtime will create new container images based on the user supplied function code. Functions will now not only be in a serving state but also will report a 'building' state. This also enables us to have an old version of the function in a serving state and simultaneously have the runtime prepare a new version of the image.

#### Services

knative-serving allows scale to zero. This means a function can have zero running instances and still be fully operational as it will automatically be scaled up to as many instances as required if request are routed to the function.

#### Function Sizes

In the current implementation function sizes (S, M, L, XL) are configured and handled by the UI-Layer. This means the sizes are translated into their configured cpu / memory requests/limits. For the new controller this mapping is handled by the controller itself. It is also possible to add or remove additional sizes. The UI just has to retrieve the configured sizes from the controller and configure the function accordingly.

### Proposal

The UI for knative-functions should be implemented as a new micro frontend. This will allow us to enable the old and the new UI at the same time and keep the code clean. The feature flag will be used to enable the new UI.

### GraphQL

In order to allow changes to the function implementation without having to change the UI all calls to the API server have to be replaced by GraphQL calls. Currently these are not implemented, so an additional GraphQL layer has to be implemented that allows CRUD operations on function objects as well as their subobjects (services, routes, builds).

### Service Catalog

In order to support ServiceBindings on the new objects we need to create a UsageKind that supports knative-functions.

```yaml
apiVersion: servicecatalog.kyma-project.io/v1alpha1
kind: UsageKind
metadata:
  annotations:
  finalizers:
  - servicecatalog.kyma-project.io/usage-kind-protection
  name: ksvc
spec:
  displayName: KSVC
  labelsPath: spec.template.metadata.labels
  resource:
    group: serving.knative.dev
    kind: Service
    version: v1alpha1
```

The problem that needs to be solved here is that as soon as the service binding controller modifies the knative-service knative-serving creates a new revision and automatically starts a new pod with the updated configuration. The PodPreset also does this creates a new pod based on the old configuration plus the updated PodPreset. So until knative-serving scales down the service the pods are duplicated.

### Monitoring and Alerting

Currently existing dashboards must be adapted to support the new knative-functions. Existing alert rules have to be adapted as well.

## Migration

* integrate the knative function controller
* install both runtimes simultaneously (installation of knative-functions is triggered using a feature flag)
  * in this phase the user can run functions on kubeless with UI integration
  * knative-functions can be scheduled using the `kubectl`
* as soon as the UI for knative-functions is available it will be enabled by the feature flag
  * the user can now use the UI to schedule knative functions
  * kubeless functions can still be scheduled using the old UI
* all other parts (alerting, monitoring, upgrade) will be enabled as soon as they are available
* once knative-functions is production ready the upgrade job will migrate all existing kubeless function to knative-functions.

## Next Steps

* implement the required UI and GraphQL changes
* create knative-functions module for kyma installer
* implement an migration job

## Additional Thoughts
* it might be necessary to run multiple function controllers in the same cluster that reconcile the same CRD. (e.g. one that schedules the function locally, on that schedules it in a different cluster). The knative-function CRD should be implemented in a way to support this.
