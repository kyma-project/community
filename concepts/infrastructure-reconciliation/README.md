# Cluster Manager proposal

## Motivation
Operators use Kubernetes' custom resources (CR) as an API and introduce a control loop to manage third-party applications and their components. The resources are reconciled periodically to ensure that all desired components are configured properly and running. This principle can be applied to cover also infrastructure components. This introduces a lower maintenance effort and provides an easy-to-use, declarative approach to infrastructure provisioning.

## Use cases for Cluster Manager

- Create/reconcile a Gardener Shoot cluster for Kyma Runtime 
- Create kubeconfigs for provisioned clusters, those should be limited by Kubernetes roles and rotated regularly. There should be also a possibility to revoke access to the cluster if needed.
- Manage hyperscaler subscriptions that are used by the Kubernetes clusters
- Support bring your own cluster models (BYOC) - a cluster is ready and we get cluster-admin kubeconfig 
- Trigger reconciliation on a configuration or a cluster state change

## Proposed changes

### Cluster Managers handling cluster provisioning

To achieve a declarative and responsive way of interacting with the infrastructure provisioning, a pair of Cluster Managers should be introduced. These components will use an event-based reconciliation of new cluster CRs, serving as an API for Kubernetes Cluster provisioning.

In order to provide a common CR for all clusters regardless of their license type there will be a new pair of cluster CRs introduced:
- GardenerCluster CR (for managed clusters running on Gardener)
- ExternalCluster CR (for Bring-Your-Own-Cluster model)

And respectively - there will be two operators handling different processes for those CRs:
- Gardener Cluster Manager, which handles Gardener Shoot provisioning and kubeconfig generation/rotation
- External Cluster Manager, which takes care of storing the kubeconfig uploaded by the customer safely and makes sure constrained kubeconfig is available for other operators

The goal is to reduce time-based loops as those result in a delayed response from the components. The Cluster Managers would react upon any Cluster CR change (creation, update, deletion) to introduce changes to the infrastructure immediately.

> Note: Restoring resources accidentally removed by the user is not the responsibility of this reconciler. Removing a whole Shoot CR from the Gardener will result in the whole cluster being deleted. This is an issue that is not easy to mitigate as this component won't have information on what were the cluster contents. Perhaps some sort of cluster backup strategy should be introduced to handle this matter.

Lifecycle Manager, along with the Component Operators, should not be able to modify the cluster metadata, however, in some potential cases, it may be useful for those components to be able to read the cluster CR status and metadata. This will provide a separation layer while leaving the possibility open for the Kyma Environment Broker and the operations team to make required changes directly.

For specific requests (such as additional network configuration or NAT Gateway toggles) the API of the Gardener Cluster Manager (GardenerCluster CR) can be extended and additional ways of interacting with the Gardener Cluster Manager could be potentially added.

### Handling hyperscaler subscriptions

Separating user subscriptions is both a business and security requirement. Right now it is handled by the Hyperscaler Account Pool package in the Kyma Environment Broker, which is not a good place to have that. The proposal is to provide an operator that would handle the available hyperscaler subscription pool, deciding whether Gardener's Shoot should be provisioned with the already used credential set or whether a new one is needed. This introduces an abstraction layer which separates both Kyma Environment Broker and Cluster Manager from the subscription handling logic and provides a way to easily extend that system in the future.

### Issuing RBAC-based kubeconfigs

To provide an environment where certain Operators cannot leverage the cluster access, it is required to open up a possibility for the constrained kubeconfigs to be created. This ensures that Operators can work only with the resources they need access to. Security-wise, potential malicious parties won't be able to use kubeconfigs issued for the Operators to execute cluster-wide attacks.

The solution is to introduce a possibility in the Cluster Managers to issue limited kubeconfigs, which will be rotated in a time-based manner. Creation of those resources in Kyma Runtimes will be done using the master kubeconfigs coming from the Gardener (or Vault for BYOC), which will not be stored anywhere in the Control Plane.

### Dependency management

It should be the responsibility of the Component Reconcilers to figure out if all dependencies exist and the reconciliation can be completed. This is true for all the infrastructure-related artefacts, such as the cluster state and kubeconfigs. If dependencies are not ready, the Component Reconciler should return the information that it is not ready and wait for the next trigger (time or event).

## High-level architecture

![](assets/infrastructure-reconciliation.svg)

The three new operators would need to be introduced:
- Gardener Cluster Manager
- External Cluster Manager
- Subscription Manager

Each of those components has a clear set of responsibilities that together handle the flow that creates a new, complete, Kyma-ready Kubernetes cluster.

Along the operators, there are four new CRs that serve as an API for the components:
- Gardener Cluster CR
- External Cluster CR
- SubscriptionRequest CR

### Cluster Managers and the Cluster CRs

The Cluster Managers' main responsibility is to manage the respective Cluster CRs, this is tied to both their provisioning and providing an access to the created Kubernetes clusters.

There will be two resources, first for the Gardener-based, managed clusters:
```
apiVersion: operator.kyma-project.io/v1alpha1
kind: GardenerCluster
metadata:
	name: managed-cluster
	namespace: default
spec:
	core:
		purpose: "Production"
		region: "eu-central-1"
	extensions:
		shootNetworkingFilterDisabled: true
	oidcConfig:
		clientID: "asdf"
		groupsClaim: "asdf"
		issuerURL: "asdf"
		signingAlgs: ["adsf"]
		usernameClaim: "asdf"
		usernamePrefix: "asdf"
	infrastructure:
		machineType: "m5.2xlarge"
		autoScalerMin: 5
		autoScalerMax: 5
		provider: aws
		numberOfZones: 3
	subscriptionSelector: "globalaccount:abcd-0001" || "shared"
	kubeconfigs:
	- secretRef: 
		secretName: managed-cluster-admin-kubeconfig
		secretNamespace: default
	  clusterRole:
		rules:
		- apiGroups: ["*"]
		  resources: ["*"]
		  verbs: ["*"]
	administrators: ["user.id@sap.com", "other.id@sap.com"]
```

This contains a minimal set of values that are required to provision a new Gardener cluster. The resource also contains the `administrators` and `subscriptionSelector` fields that are required for the Binding creation and requesting a hyperscaler subscription respectively.

And the second one for the Bring Your Own Cluster (BYOC) model:
```
apiVersion: operator.kyma-project.io/v1alpha1
kind: ExternalCluster
metadata:
	name: external-cluster
	namespace: default
spec:
	adminKubeconfigRef:
		secretName: user-provided-kubeconfig
		secretNamespace: default
	kubeconfigs:
	- secretRef: 
		secretName: external-cluster-admin-kubeconfig
		secretNamespace: default
	  clusterRole:
		rules:
		- apiGroups: ["*"]
		  resources: ["*"]
		  verbs: ["*"]
```

This is the simpler one, as there is no cluster provisioning happening in this flow most of the parameters can be dropped.

In both resources' spec, the `kubeconfigs` sections describes all the kubeconfigs that should be issued for the use of KCP components. 

ExternalCluster CR also has a `adminKubeconfigRef` section that should point to the admin kubeconfig delivered by the user. This kubeconfig will be saved to the secure Vault and removed from the cluster.

### Issuing kubeconfigs

As stated before, the separation of kubeconfigs and limiting access is needed to have a healthy, secure environment. Cluster Managers will provide a possibility to create RBAC-based kubeconfigs in both BYOC and Gardener clusters. This will provide operators with the required access to perform any necessary actions.

The kubeconfigs for specific Kubernetes clusters will be fetched (both from Vault or Gardener) in the runtime and stored only in the components' memory for the limited time needed to issue another (constrained) kubeconfig for a given cluster. Those new kubeconfigs (used by the Lifecycle and Module Managers) will have a short expiration date and will be available in the KCP.

### Subscription Manager and the SubscriptionRequest CR

This is a part that needs to be extracted from the Kyma Environment Broker, there are no significant changes in this area. The main change is introducing a new declarative API in form of the SubscriptionRequest CR:
```
apiVersion: control-plane.kyma-project.io/v1alpha1
kind: SubscriptionRequest
metadata:
    name: subscription-request
    namespace: cluster-manager
spec:
    provider: aws
    subscriptionSelector: "abcd-0001"
status:
    gardenerSecretName: subscription-secret
```

The resource structure is really simple. Subscription Manager only needs to know the provider and a proper selector for the subscription to be selected. It provides credentials in the form of the Gardener's Secret name, which can be simply used as a parameter in the Shoot CR.

## Cluster Provisioning Flow

The flow from the user's perspective stays intact, the part that changes is the list of actions needed by KEB to make sure the cluster is up and running.

First of all, after receiving the provisioning request, KEB should create a set of CRs for accordingly to the licence type of the Runtime:
- For BYOC - ExternalCluster CR
- For managed - GardenerCluster CR

> In the BYOC flow, KEB should also create a secret with the kubeconfig uploaded by the user and provide a reference to that resource in the ExternalCluster CR

KEB should deliver all the required parameters for the cluster provisioning, this translates to the minimal infrastructure config for Gardener clusters, and the role/kubeconfig details for both types.

Another thing that KEB would need to create is the Kyma CR for the Runtime, this will either require passing the proper Secret reference so that kubeconfig can be accessed by the Lifecycle Manager.

After the CRs are created, the respective Cluster Manager will pick up the provisioning request and will create a new Shoot for the runtime (managed scenario only), and provide a requested kubeconfig.

At this point Lifecycle Manager can step in, it has all required resources to set up the Runtime:
- running cluster
- kubeconfig (passed via secret reference)
- Kyma CR

## Reacting to the kubeconfig rotation

Since Lifecycle Manager and Module Manager need to interact with the provisioned Kubernetes cluster directly, a mechanism for providing them with an up-to-date kubeconfig is needed.

There are few possible options here:
- Watching Kubernetes Secrets - Managers should watch the Secrets containing the kubeconfigs, with a big amount of Runtimes this may be resource intensive on the API server side.
- Implementing caching logic based on the timestamps for the kubeconfig expiration - Managers would determine themselves whether a new kubeconfig should be fetched or not, this one requires implementing custom caching mechanisms in the components.
- Simple retries - if the Manager is unable to connect to the cluster with it's current kubeconfig, it should try fetching a new one, it's not the prettiest and most performant option though.

Before making a final decision on that part, we should consider different options and compare them (at least) performance-wise.