# Cluster Operator proposal

## Motivation
Operator pattern is a framework that uses Kubernetes' Custom Resources as an API and introduces a control loop to manage applications and their components. The resources are reconciled periodically to ensure that all desired components are configured properly and running. This principle can be extended to cover also infrastructure components, this will introduce a lower maintenance effort and provide easy to use, declarative approach to infrastructure provisioning.

## Use cases for Cluster Operator

- Create/reconcile a Gardener Shoot cluster for Kyma Runtime 
- Create kubeconfigs for provisioned clusters, those should be limited by Kubernetes roles and rotated regularly. There should be also a possibility to revoke access to the cluster if needed.
- Manage hyperscaler subscriptions that are used by the Kubernetes Clusters
- Support bring your own cluster models (BYOC) - a cluster is ready and we get cluster-admin kubeconfig 
- Trigger reconciliation on a configuration or a cluster state change

## Proposed changes

### Cluster Operator handling cluster provisioning

To achieve a declarative and responsive way of interacting with the infrastructure provisioning a new Cluster Operator should be introduced. This component will use an event-based reconciliation of a new Cluster Custom Resource, which will serve as an API for Kubernetes Cluster provisioning.

The goal is to reduce time-based loops as those result in a delayed response from the components. The Cluster Operator would react upon any Cluster CR change (creation, update, deletion) to introduce changes to the infrastructure immediately.

> Note: Restoring resources accidentally removed by the user is not the responsibility of this reconciler. Removing a whole Shoot CR from the Gardener will result in the whole cluster being deleted. This is an issue that is not easy to mitigate as this component won't have information on what were the cluster contents. Perhaps some sort of cluster backup strategy should be introduced to handle this matter.

Lifecycle Manager, along with the Component Operators, should not be able to modify the cluster metadata, however, in some potential cases, it may be useful for those components to be able to read the Cluster CR status and metadata. This will provide a separation layer while leaving the possibility open for the Kyma Environment Broker and the operations team to make required changes directly.

For specific requests (such as additional network configuration or NAT Gateway toggles) the API of the Cluster Operator (Cluster CR) can be extended and additional ways of interacting with the Cluster Operator could be potentially added.

### Handling hyperscaler subscriptions

Separating user subscriptions is both a business and security requirement. Right now it is handled by the Hyperscaler Account Pool package in the Kyma Environment Broker, which is not a good place to have that. The proposal is to provide an operator that would handle the available hyperscaler subscription pool, deciding whether Gardener's Shoot should be provisioned with the already used credential set or whether a new one is needed. This introduces an abstraction layer which separates both Kyma Environment Broker and Cluster Operator from the subscription handling logic and provides a way to easily extend that system in the future.

### Issuing RBAC-based kubeconfigs

To provide an environment, where certain Operators are not able to leverage the cluster access, it is required to open up a possibility for the constrained kubeconfigs to be created. This ensures that Operators can work only with the resources they require access to and security-wise, potential malicious parties won't be able to use kubeconfigs issued for the Operators to execute cluster-wide attacks.

The solution is to introduce yet another operator which would be designed around creating Roles, RoleBindings and ServiceAccounts in the Kubernetes cluster. Those resources can be used to issue limited kubeconfigs, which will be rotated in a time-based manner. Creation of those resources in Kyma Runtimes will be done using the master kubeconfigs coming from the Gardener, which will not be stored anywhere in the Control Plane.

### Dependency management

It should be the responsibility of the Component Reconcilers to figure out if all dependencies exist and the reconciliation can be completed. This is true for all the infrastructure-related artefacts, such as the cluster state and kubeconfigs. If dependencies are not ready, the Component Reconciler should return the information that it is not ready and wait for the next trigger (time or event).

## High-level architecture

![](assets/infrastructure-reconciliation.svg)

The three new operators would need to be introduced:
- Cluster Operator
- Subscription Operator
- Kubeconfig Operator

Each of those components has a clear set of responsibilities that together handle the flow that creates a new, complete, Kyma-ready  Kubernetes cluster.

Along the operators, there are three new Custom Resources, that serve as an API for the components:
- Cluster CR
- SubscriptionRequest CR
- KubeconfigRequest CR

### Cluster Operator and the Cluster CR

The Cluster Operator's main responsibility is to manage the Cluster CRs, request subscriptions for the provisioning and create Shoot CRs in the Gardener cluster.

On top of that, there is also a requirement to provide out-of-the-box admin roles to some users, this is the only place where the Cluster Operator interacts directly with the created cluster.

Also, the exemplary Cluster CR resource looks as follows:
```
apiVersion: operator.kyma-project.io/v1alpha1
kind: Cluster
metadata:
    name: example-cluster
    namespace: cluster-operator
    labels:
        plan: "trial"
        globalaccountid: "abcd-0001"
        subaccountid: "abcd-0001"
        instanceid: "abcd-0001"
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
    administrators: ["user.id@sap.com", "other.id@sap.com"]
status:
    conditions:
        - type: ClusterProvisioned
          status: true
          lastTransitionTime: "timestamp"
          lastUpdateTime: "timestamp"
          reason: "reason"
          message: "cluster running"
```

This contains a minimal set of values that are required to provision a new Gardener cluster. The resource also contains the `administrators` and `subscriptionSelector` fields that are required for the Binding creation and requesting a hyperscaler subscription respectively.

### Subscription Operator and the SubscriptionRequest CR

This is a part that needs to be extracted from the Kyma Environment Broker, there are no significant changes in this area. The main change is introducing a new declarative API in form of the SubscriptionRequest CR:
```
apiVersion: control-plane.kyma-project.io/v1alpha1
kind: SubscriptionRequest
metadata:
    name: subscription-request
    namespace: cluster-operator
spec:
    provider: aws
    subscriptionSelector: "abcd-0001"
status:
    gardenerSecretName: subscription-secret
```

The resource structure is really simple, the Subscription Operator only needs to know the provider and a proper selector for the subscription to be selected. It provides credentials in a form of the Gardener secret name, which can be simply used as a parameter in the Shoot CR.

### Kubeconfig Operator and the KubeconfigRequest CR

As stated before, the separation of kubeconfigs and limiting access is needed to have a healthy, secure environment. Kubeconfig Operator is a component that will handle the kubeconfig requests for all the component operators (including Lifecycle Manager). To achieve that, access to the Gardener cluster is needed.

The kubeconfigs for specific Kubernetes clusters will be fetched in the runtime and stored only in the component's memory for the limited time needed to issue another (constrained) kubeconfig for a given cluster.

There is a possibility to drop the Kubeconfig Operator - Gardener connection, however, it will require storing and rotating another kubeconfig used solely for issuing another kubeconfigs, in the end, it does not differ from storing a master-kubeconfig (as it can create any Role and ServiceAccount the attacker would potentially require). We want to avoid such a situation - storing those configs is not a secure solution and we do need Gardener kubeconfig in the Control Plane anyways for the Cluster Operator and Shoot creation purposes. 

The other argument against dropping that connection is that the Kubeconfig Operator could potentially lose access to the cluster as well (if a failure/restart/network error occurs). Every time such an issue would pop up it will require manual action to restore access to the cluster.

The structure of the KubeconfigRequest is quite simple:
```
apiVersion: operator.kyma-project.io/v1alpha1
kind: KubeconfigRequest
metadata:
    name: example-cluster.compass-operator
    namespace: default
    annotations:
        operator.kyma-project.io/rotate: false
        operator.kyma-project.io/revoke: false
spec:
    cluster: example-cluster
    clusterRole:
        rules:
        - apiGroups: ["application-connector.kyma-project.io/v1alpha1"]
          resources: ["Application"]
          verbs: ["create", "get", "watch", "update", "delete"]
    role:
        namespace: "compass-system"
        rules:
        - apiGroups: ["*"]
          resources: ["Secret"]
          verbs: ["create", "get", "update", "delete"]
    expiration: 3600
status:
    expirationTimestamp: "2022-08-23T07:05:21Z"
    kubeconfig: "LS0tLS1..."
```

With that structure, the Kubeconfig Operator can create kubeconfigs with either (or both) namespaced and cluster-wide access to the specific resources. Since we want to rotate the kubeconfigs regularly the expiration time is also needed. 

After the kubeconfig is issued it will be encoded with the Base64 algorithm and placed in the resource `status` field (along its expiration timestamp).

The assumption is that Kubeconfig Operator will keep this Kubeconfig up-to-date and specific Component Operators will handle the logic of fetching a new kubeconfig every time the old one expires.

For security reasons there are two annotations added - one for rotating and the other one for revoking the kubeconfig, this allows easy access to the security measures in case one or all of the kubeconfigs are compromised.