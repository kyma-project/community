# Dynamic Fluent Bit Configuration

This document investigates a valid design for a custom operator, which is needed to enable a dynamic in-cluster logging configuration with `Fluent Bit`. The regarding issue can be found [here](https://github.com/kyma-project/kyma/issues/11105).

## Criterias
- Customers can place fluent-bit config snippets as k8s resource in a specific or in any namespace
- Have in mind that the central configuration will be overwritten at any time by reconcilation which should not reset customers configuration
- A new config gets picked up dynamically without any further interaction (like pod restart by the user)
- Have basic validation to give early feedback on errors -> Debugging a startup problem
- OpenTelemetry collector has the same problem to solve, proposal should be extensible
- There is a way to provide the auth details for a backend config in a secure way

## Proposal

### Architecture

To build the operator we will use an SDK for a quick and easy start-up, besides various other reasons. For this, we decided to use [`kubebuilder`](https://github.com/kubernetes-sigs/kubebuilder). `kubebuilder` is an upstream project of `Kubernetes`. It provides us with features we need for our operator with the smallest overhead. In addition, this SDK is also used in other parts of Kyma, and thus we have a more consistent picture.

An alternative for `kubebuilder` would be the [`Operator-SDK`](https://github.com/operator-framework/operator-sdk) from `Red Hat`. This has similar features to the kubebuilder, with some additions like the integration of [`operatorhub`](https://operatorhub.io/) - but such features are not needed for our purpose. Furthermore, this is not an upstream project of Kubernetes.

To have an simple API, one `CRD` for Fluent Bit configuration will be created. This includes a field to determine the type of the configuration (`FILTER` or `OUTPUT`). A list of maps for the Fluent Bit configuration code, and a status field to keep track of the status of the later created CR.
Using a list of maps for the actual configuration, enables the possibility to support the full Fluent Bit syntax without having the need to maintain new features of various plugins. Furthermore, this gives the user the ability to have a good overview of his/hers sequence of applied filters and outputs. Using the Kyma documentation, we could also lead the user to think more in a way of pipelines, in such that (s)he creates one CR for each Fluent Bit pipeline.

One constraint of this operator will be, that it will not support dynamic plugins which require to be loaded into the Fluent Bit image. Another constraint will be, that not all available plugins for filters and outputs can be configured by this operator (i.e. `Lua`), because such plugins need to have certain files mounted into the container, which is not possible. If a user still wants to use such, (s)he needs to configure an output plugin of his/her choice and process the logs with another log processor (i.e. another instance of Fluent Bit). The operator has a list of unsupported plugins, to achive that such plugins cannot be configured. Thus, every time the user creates a CR, the list is checked if it contains a wanted plugin, if it does the user will be informed that (s)he is not able to create such CR.

<details>
<summary><b>Pipeline Overview</b> for User</summary>

![Thank you](images/fluentbit_CR_overview.svg)
</details>  

The logs will be fetched using a Fluent Bit `INPUT` which will be created by this operator. This INPUT tags all logs with the following tag scheme: `kube.<namespace_name>.<pod_name>.<container_name>`. Afterwards, the logs will be enriched with Kubernetes informations using the `kubernetes` filter of Fluent Bit. Then, the filter `rewrite_tag` will be used to split the pipeline. The original logs are forwarded to the Kyma Loki backend, and the new copy with another tag can be used by the user by configuring new filters and outputs with the provided CRD.Using this approach, we avoid the situation of an unused INPUT or other overhead.

![Fluent Bit Pipeline Architecture](images/fluentbit_dynamic_config.svg)

If the user wants to use more than one pipeline for the log processing, (s)he can use the 'rewrite_tag' filter on their pipeline to create more pipelines. Or (s)he can configure an output plugin to process them with another log processor, as mentioned before.

Additionally, when creating a `CR` a [webhook](https://book.kubebuilder.io/cronjob-tutorial/webhook-implementation.html) will be used to validate the correctness of the Fluent Bit configuration based on the the Fluent Bit `dry-run` feature.

To actually apply the changes of the user, the operator will create/adapt the ConfigMap for FluentBit and restart Fluent Bit by deleting the pods of the Fluent Bit deployment.

<details>
  <summary><b>CRD Definition (go code) </b>- Click to expand</summary>

```go
// ConfigSectionSpec defines the desired state of ConfigSection
type ConfigSectionSpec struct {
	Type string `json:"type,omitempty"`
	Entries map[string]string `json:"entries,omitempty"`
}
// ConfigSectionStatus defines the observed state of ConfigSection
type ConfigSectionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// ConfigSection is the Schema for the configsections API
type ConfigSection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   ConfigSectionSpec   `json:"spec,omitempty"`
	Status ConfigSectionStatus `json:"status,omitempty"`
}
//+kubebuilder:object:root=true
// ConfigSectionList contains a list of ConfigSection
type ConfigSectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigSection `json:"items"`
}
```
</details>


<details>
  <summary><b>CRD Definition (yaml file, created using kubebuilder SDK)</b> - Click to expand</summary>

```yaml
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.4.1
  creationTimestamp: null
  name: configsections.logging.kyma-project.io
spec:
  group: logging.kyma-project.io
  names:
    kind: ConfigSection
    listKind: ConfigSectionList
    plural: configsections
    singular: configsection
  scope: Namespaced
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: ConfigSection is the Schema for the configsections API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ConfigSectionSpec defines the desired state of ConfigSection
            properties:
              entries:
                additionalProperties:
                  type: string
                type: object
              type:
                type: string
            type: object
          status:
            description: ConfigSectionStatus defines the observed state of ConfigSection
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []

```
</details>

 

### Workflow for the User

To configure Fluent Bit the user has to create a new CR regarding to the CRD of this operator. The operator then will notice the newly created or changed CR, and will create or update a ConfigMap for Fluent Bit. Before the ConfigMap gets applied, the operator uses the `dry-run` feature of Fluent Bit to validate the new configuration. If the check was successfull, the new ConfigMap will be applied and Fluent Bit will be restarted.
