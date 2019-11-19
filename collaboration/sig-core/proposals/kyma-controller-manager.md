# Consolidate Kyma controllers inside a single controller manager

Created on 2019-10-11 by Antoine Cotten (@antoineco).

## Status

DRAFT. Proposed on 2019-10-11.

## Motivation

Kyma is composed of multiple components that share a lot of architectural patterns, yet are written in isolation
following different engineering standards and using heterogeneous SDKs and libraries, before eventually being released
as _one product_.

Most of these components involve one or more Kubernetes controllers that synchronize API objects using identical
building blocks: watches, work queues, reconciliation loops, etc. Every controller instance maintains its own local
cache of several Kubernetes resources, which can result in a significant memory footprint in productive environments.
Besides, each of these controllers is operated in a similar fashion by nature: as a single instance or in an
active-passive setup.

These similarities represent an opportunity for us to consolidate Kyma controllers and run them under the umbrella of a
single _manager_, so that they can reuse common libraries, maintain a consistent level of compatibility with the
underlying Kubernetes layer, share memory efficiently, and avoid the fragmentation we are currently observing in the
project.

This mode of operation is directly inspired by the Kubernetes [controller-manager][k8s-cm].

## Goals

This initiative aims at improving three aspects:

**Maintainability**

Streamline the codebase regarding controllers to harmonize the different implementation aspects. (_"Common opinion"_)

**Footprint**

Reduce the resource consumption by sharing internal API object caches whenever possible. (_"Store once"_)

**Operability**

Improve operational aspects by reducing the number of controllers that have to be deployed separately. (_"Run once"_)

## Constraints

It must be possible to opt out of each individual controller inside the manager to keep Kyma modular. This could be
achieved using startup flags and/or a ConfigMap.

A controller ("operator") SDK should be agreed upon by all code owners. [kyma#4327][kyma4327] initiated
a discussion around that topic, but no conclusion has been reached at the time of writing. For the initial version of
this document, we assume that [`kubebuilder`][kb] will be our SDK of choice. A few reasons to believe this assumption is
safe to make:
- well maintained and documented by a dedicated group (SIG) within the Kubernetes project
- leverages `controller-runtime`, a collection of sensible libraries with everything a controller typically needs
- bootstaps projects with the manager pattern enabled right from the start
- includes tooling oriented towards future extensibility, with helpers to generate new APIs
- RBAC objects are part of the code generation

## Proposed solution

### Project structure

The structure of the controller manager is initialized from _scratch_, existing Kyma components get incorporated
gradually into it after suitable candidates have been identified.

A rough draft of what the structure of the Kyma controller manager could look like:

```
├── api
│   ├── applicationconnector
│   │   └── v1alpha1
│   │       ├── application_types.go
│   │       ├── <other types>…
│   │       ├── groupversion_info.go
│   │       └── zz_generated.deepcopy.go
│   ├── eventing
│   │   └── v1alpha1
│   │       ├── subscription_types.go
│   │       ├── <other types>…
│   │       ├── groupversion_info.go
│   │       └── zz_generated.deepcopy.go
│   ├── serverless
│   │   └── v1alpha1
│   │       ├── function_types.go
│   │       ├── <other types>…
│   │       ├── groupversion_info.go
│   │       └── zz_generated.deepcopy.go
│   └── <other API groups>…
│
├── controllers
│   ├── applicationconnector
│   │   ├── application_controller.go
│   │   ├── <other resource controllers>…
│   │   └── pkg
│   ├── eventing
│   │   ├── subscription_controller.go
│   │   ├── <other resource controllers>…
│   │   └── pkg
│   ├── serverless
│   │   ├── function_controller.go
│   │   ├── <other resource controllers>…
│   │   └── pkg
│   └── commons
│
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
└── Makefile
```

- `api` contains the various API definitions (custom types)
- `controllers/<group>` contains the high level interfaces and methods for the controllers of the given API group
- `controller/commons` contains packages that can be reused by controllers from multiple groups

> NOTE: although this project structure is heavily inspired by `kubebuilder`, it is currently not possible to scaffold
> multiple API groups using the `kubebuilder create api` helper. This setup is nevertheless [supported][kb-multigroup]
> if we decide to adopt that SDK.

### Entry point

The project entry point (`main`) is responsible for initializing common resources and starting all requested
controllers.

```go
var enableLeaderElection bool
flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
	"Enable leader election for controller manager. Ensures there is only one active controller manager.")
flag.Parse()

// Register custom types

var scheme = runtime.NewScheme()
applicationconnectorv1alpha1.AddToScheme(scheme)
eventingv1alpha1.AddToScheme(scheme)
serverlessv1alpha1.AddToScheme(scheme)

// Initialize global logger

ctrl.SetLogger(zap.Logger())

// Initialize controller manager
// * start shared informers
// * handle leader election
// * expose metrics endpoints

mgr := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
	Scheme:         scheme,
	LeaderElection: enableLeaderElection,
})

// Register individual controllers

&serverlesscontrollers.FunctionReconciler{
	Client: mgr.GetClient(),
	Log:    ctrl.Log.WithName("serverless").WithName("Function"),
}.SetupWithManager(mgr)

&eventingcontrollers.SubscriptionReconciler{
	Client: mgr.GetClient(),
	Log:    ctrl.Log.WithName("eventing").WithName("Subscription"),
}.SetupWithManager(mgr)

&applicationconnectorcontrollers.ApplicationReconciler{
	Client: mgr.GetClient(),
	Log:    ctrl.Log.WithName("applicationconnector").WithName("Application"),
}.SetupWithManager(mgr)

// Start controller manager

mgr.Start(ctrl.SetupSignalHandler())
```

> NOTE: again, types and methods used in this example are directly sourced from code scaffolded by `kubebuilder`.



[k8s-cm]: https://kubernetes.io/docs/concepts/overview/components/#kube-controller-manager
[kyma4327]: https://github.com/kyma-project/kyma/issues/4327
[kb]: https://github.com/kubernetes-sigs/kubebuilder
[kb-multigroup]: https://book.kubebuilder.io/migration/multi-group.html
