# Cluster Manager proposal

## Motivation
Operators use Kubernetes' Custom Resources as an API and introduce a control loop to manage third party applications and their components. The resources are reconciled periodically to ensure that all desired components are configured properly and running. This principle can be applied to cover also infrastructure components, this will introduce a lower maintenance effort and provides an easy to use, declarative approach to infrastructure provisioning.

## Use cases for Cluster Manager

- Create/reconcile a Gardener Shoot cluster for Kyma Runtime 
- Create kubeconfigs for provisioned clusters, those should be limited by Kubernetes roles and rotated regularly. There should be also a possibility to revoke access to the cluster if needed.
- Manage hyperscaler subscriptions that are used by the Kubernetes Clusters
- Support bring your own cluster models (BYOC) - a cluster is ready and we get cluster-admin kubeconfig 
- Trigger reconciliation on a configuration or a cluster state change

## Proposed changes

### Cluster Managers handling cluster provisioning

To achieve a declarative and responsive way of interacting with the infrastructure provisioning a pair of Cluster Managers should be introduced. This components will use an event-based reconciliation of a new Cluster Custom Resources, which will serve as an API for Kubernetes Cluster provisioning.

In order to provide a common Custom Resource for all clusters regardless of their licence type there will be a new pair of cluster CRs introduced:
- GardenerCluster CR (for managed clusters running on Gardener)
- ExternalCluster CR (for Bring-Your-Own-Cluster model)

And respectively - there will be two operators handling different processes for those CRs:
- Gardener Cluster Manager, which will handle Gardener Shoot provisioning and kubeconfig generation/rotation
- External Cluster Manager, which takes care of storing the kubeconfig uploaded by the customer safely and makes sure constrained kubeconfig is available for other operators

The goal is to reduce time-based loops as those result in a delayed response from the components. The Cluster Managers would react upon any Cluster CR change (creation, update, deletion) to introduce changes to the infrastructure immediately.

> Note: Restoring resources accidentally removed by the user is not the responsibility of this reconciler. Removing a whole Shoot CR from the Gardener will result in the whole cluster being deleted. This is an issue that is not easy to mitigate as this component won't have information on what were the cluster contents. Perhaps some sort of cluster backup strategy should be introduced to handle this matter.

Lifecycle Manager, along with the Component Operators, should not be able to modify the cluster metadata, however, in some potential cases, it may be useful for those components to be able to read the cluster CR status and metadata. This will provide a separation layer while leaving the possibility open for the Kyma Environment Broker and the operations team to make required changes directly.

For specific requests (such as additional network configuration or NAT Gateway toggles) the API of the Gardener Cluster Manager (GardenerCluster CR) can be extended and additional ways of interacting with the Gardener Cluster Manager could be potentially added.

### Handling hyperscaler subscriptions

Separating user subscriptions is both a business and security requirement. Right now it is handled by the Hyperscaler Account Pool package in the Kyma Environment Broker, which is not a good place to have that. The proposal is to provide an operator that would handle the available hyperscaler subscription pool, deciding whether Gardener's Shoot should be provisioned with the already used credential set or whether a new one is needed. This introduces an abstraction layer which separates both Kyma Environment Broker and Cluster Manager from the subscription handling logic and provides a way to easily extend that system in the future.

### Issuing RBAC-based kubeconfigs

To provide an environment, where certain Operators are not able to leverage the cluster access, it is required to open up a possibility for the constrained kubeconfigs to be created. This ensures that Operators can work only with the resources they require access to and security-wise, potential malicious parties won't be able to use kubeconfigs issued for the Operators to execute cluster-wide attacks.

The solution is to introduce a possibility in the Cluster Managers to issue limited kubeconfigs, which will be rotated in a time-based manner. Creation of those resources in Kyma Runtimes will be done using the master kubeconfigs coming from the Gardener (or Vault for BYOC), which will not be stored anywhere in the Control Plane.

### Dependency management

It should be the responsibility of the Component Reconcilers to figure out if all dependencies exist and the reconciliation can be completed. This is true for all the infrastructure-related artefacts, such as the cluster state and kubeconfigs. If dependencies are not ready, the Component Reconciler should return the information that it is not ready and wait for the next trigger (time or event).

## High-level architecture

![](assets/infrastructure-reconciliation.svg)

The three new operators would need to be introduced:
- Gardener Cluster Manager
- External Cluster Manager
- Subscription Manager

Each of those components has a clear set of responsibilities that together handle the flow that creates a new, complete, Kyma-ready  Kubernetes cluster.

Along the operators, there are four new Custom Resources, that serve as an API for the components:
- Gardener Cluster CR
- External Cluster CR
- SubscriptionRequest CR

### Cluster Managers and the Cluster CRs

The Cluster Managers' main responsibility is to manage the respective Cluster CRs, this is tied to both their provisioning and providing an access to the created K8S clusters.

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
		name: managed-cluster-admin-kubeconfig
		namespace: default
	  clusterRole:
		rules:
		- apiGroups: ["*"]
		  resources: ["*"]
		  verbs: ["*"]
	administrators: ["user.id@sap.com", "other.id@sap.com"]
```

This contains a minimal set of values that are required to provision a new Gardener cluster. The resource also contains the `administrators` and `subscriptionSelector` fields that are required for the Binding creation and requesting a hyperscaler subscription respectively.

And the second one for BYOC model:
```
apiVersion: operator.kyma-project.io/v1alpha1
kind: ExternalCluster
metadata:
	name: external-cluster
	namespace: default
spec:
	adminKubeconfig: LS0tV1...
	kubeconfigs:
	- secretRef: external-cluster-admin-kubeconfig
	  clusterRole:
		rules:
		- apiGroups: ["*"]
		  resources: ["*"]
		  verbs: ["*"]
```

This is the simpler one, as there is no cluster provisioning happening in this flow most of the parameters can be dropped.

### Issuing kubeconfigs

As stated before, the separation of kubeconfigs and limiting access is needed to have a healthy, secure environment. Cluster Managers will provide a possibility to create RBAC-based kubeconfigs in both BYOC and Gardener clusters. This will provide operators with the required access to perform any necessary actions.

The kubeconfigs for specific Kubernetes clusters will be fetched (both from Vault or Gardener) in the runtime and stored only in the components' memory for the limited time needed to issue another (constrained) kubeconfig for a given cluster.

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

The resource structure is really simple, the Subscription Manager only needs to know the provider and a proper selector for the subscription to be selected. It provides credentials in a form of the Gardener secret name, which can be simply used as a parameter in the Shoot CR.

## Cluster Provisioning Flow

The flow from the user's perspective stays intact, the part that changes is the list of actions needed by KEB to make sure the cluster is up and running.

First of all, after receiving the provisioning request, KEB should create a set of CRs for accordingly to the licence type of the Runtime:
- For BYOC - ExternalCluster CR
- For managed - GardenerCluster CR

KEB should deliver all the required parameters for the cluster provisioning, this translates to the minimal infrastructure config for Gardener clusters, and the role/kubeconfig details for both types.

Another thing that KEB would need to create is the Kyma CR for the Runtime, this will either require passing the proper secret reference so that kubeconfig can be accessed by the Lifecycle Manager.

After the CRs are created, the respective Cluster Manager will pick up the provisioning request and will create a new Shoot for the runtime (managed scenario only), and provide a requested kubeconfig.

At this point Lifecycle Manager can step in, it has all required resources to set up the Runtime:
- running cluster
- kubeconfig (passed via secret reference)
- Kyma CR