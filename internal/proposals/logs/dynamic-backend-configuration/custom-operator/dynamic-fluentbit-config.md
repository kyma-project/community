# Dynamic Fluent Bit Configuration

This document proposes and evaluates a valid design for a custom operator. It is needed to enable a dynamic in-cluster logging configuration with `Fluent Bit`, as outlined in [Spike: Dynamic configuration of logging backend #11105](https://github.com/kyma-project/kyma/issues/11105).

## Criteria
- Customers can place fluent-bit config snippets as k8s resource in any namespace.
- The customers' configuration must not be reset during reconciliation, even though the central configuration might be overwritten at any time.
- A new configuration is picked up dynamically without any user interaction (for example, no need that users restart pods).
- Have basic validation to give early feedback on errors, which supports debugging a startup problem.
- The proposal should be extensible to solve the same issue for OpenTelemetry collector.
- There's a way to provide the auth details for a backend config in a secure way.

## Proposal

### Architecture

To build the operator, we use an SDK for a quick and easy start-up, among other reasons. We decided to use [`kubebuilder`](https://github.com/kubernetes-sigs/kubebuilder), an upstream project of `Kubernetes`. `kubebuilder` provides the required features for our operator with the smallest overhead. In addition, this `kubebuilder` is also used in other parts of Kyma, ensuring consistent design.

An alternative for `kubebuilder` is the [`Operator-SDK`](https://github.com/operator-framework/operator-sdk) from `Red Hat`. `Operator-SDK` has similar features to the `kubebuilder`, with some additions like the integration of [`operatorhub`](https://operatorhub.io/). We decided against it because such features aren't needed for our purpose, and because it isn't an upstream project of Kubernetes.

To have a simple API, one `CRD` for Fluent Bit configuration is created. This CRD has a field which holds the status of the CR, called `Status`, as well as a struct of the type `LoggingConfigurationSpec`, called `Spec`, holding a list of configuration sections, called `Sections`. Each `Section` determines the type of the configuration (`FILTER` or `OUTPUT`) and holds a list for configuration entries, called `Entries`. Each `Entry` has three values: The `Name`, determining the key for the value, the `Value`, and the `SecretValue`. Either the `Value` or the `SecretValue` should be used by the users for the `Name`: If a `Secret` must be stated, for example, for an API key, then a name, the namespace of the key, and the key itself must be stated. If no `Secret` is needed for a `Name`, then `Value` should be chosen.

Using this structure supports the full Fluent Bit syntax without needing to maintain new features of various plugins. Furthermore, the users get a clear overview of their sequence of applied filters and outputs. Using the Kyma documentation, we could also lead the users to think more in a way of pipelines, in such that they create one CR for each Fluent Bit pipeline.

<details>
<summary><b>Pipeline Overview</b> for User - Click to expand</summary>

![Thank you](images/fluentbit_CR_overview.svg)
</details> 

We're proposing the following constraints for the custom operator: 
- It doesn't support dynamic plugins that must be loaded into the Fluent Bit image. 
- Not all available plugins for filters and outputs can be configured by this operator (for example, `Lua`), because such plugins need certain files mounted into the container, which isn't possible. However, this can be mitigated: Users can configure an output plugin of their choice and process the logs with another log processor (for example, another instance of Fluent Bit). If users create a CR, the operator checks it against a list of unsupported plugins, and if there is a match, informs the users that they cannot create that CR.

![Fluent Bit Pipeline Architecture](images/fluentbit_dynamic_config.svg)

1. The logs are fetched using a Fluent Bit, which is created by this custom operator. This INPUT tags all logs with the following tag scheme: `kube.<namespace_name>.<pod_name>.<container_name>`. 
2. The logs are enriched with Kubernetes information using the `kubernetes` filter of Fluent Bit. 
3. The filter `rewrite_tag` is used to split the pipeline:
   - The original logs are forwarded to the Kyma Loki backend.
   - Users can use the new copy with another tag and configure new filters and outputs with the provided CRD.

Using this approach, we avoid having an unused INPUT or other overhead.
If users want to use more than one pipeline for the log processing, they can use the 'rewrite_tag' filter on their pipeline to create more pipelines. Alternatively, they can configure an output plugin to process them with another log processor, as mentioned before.

Additionally, when creating a `CR`, a [webhook](https://book.kubebuilder.io/cronjob-tutorial/webhook-implementation.html) validates the correctness of the Fluent Bit configuration based on the Fluent Bit `dry-run` feature.

To actually apply the changes of the users, the operator creates or adapts the ConfigMap for FluentBit and restarts Fluent Bit by deleting the pods of the Fluent Bit deployment.

To make sure that the configuration by the users won't be overwritten by reconciliation, the basic configuration (`kubernetes` filter, `rewrite_tag`, etc.) is written into a ConfigMap. This ConfigMap is embedded by an `@INCLUDE` statement in the chart of this operator.

<details>
  <summary><b>CustomResourceDefinition (go code) </b>- Click to expand</summary>

```go
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LoggingConfigurationSpec defines the desired state of LoggingConfiguration
type LoggingConfigurationSpec struct {
	Sections []Section `json:"sections,omitempty"`
}

type Section struct {
	Type    string  `json:"type,omitempty"`
	Entries []Entry `json:"entries,omitempty"`
}

type Entry struct {
	Name        string      `json:"name,omitempty"`
	Value       string      `json:"value,omitempty"`
	SecretValue SecretValue `json:"secretValue,omitempty"`
}

type SecretValue struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Key       string `json:"key,omitempty"`
}

// LoggingConfigurationStatus defines the observed state of LoggingConfiguration
type LoggingConfigurationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// LoggingConfiguration is the Schema for the loggingconfigurations API
type LoggingConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoggingConfigurationSpec   `json:"spec,omitempty"`
	Status LoggingConfigurationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LoggingConfigurationList contains a list of LoggingConfiguration
type LoggingConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoggingConfiguration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LoggingConfiguration{}, &LoggingConfigurationList{})
}
```
</details>


<details>
  <summary><b>CustomResourceDefinition (yaml file, created using kubebuilder SDK)</b> - Click to expand</summary>

```yaml
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.4.1
  creationTimestamp: null
  name: loggingconfigurations.telemetry.kyma-project.io
spec:
  group: telemetry.kyma-project.io
  names:
    kind: LoggingConfiguration
    listKind: LoggingConfigurationList
    plural: loggingconfigurations
    singular: loggingconfiguration
  scope: Namespaced
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: LoggingConfiguration is the Schema for the loggingconfigurations
          API
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
            description: LoggingConfigurationSpec defines the desired state of LoggingConfiguration
            properties:
              sections:
                items:
                  properties:
                    entries:
                      items:
                        properties:
                          name:
                            type: string
                          secretValue:
                            properties:
                              key:
                                type: string
                              name:
                                type: string
                              namespace:
                                type: string
                            type: object
                          value:
                            type: string
                        type: object
                      type: array
                    type:
                      type: string
                  type: object
                type: array
            type: object
          status:
            description: LoggingConfigurationStatus defines the observed state of
              LoggingConfiguration
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

To configure Fluent Bit, the user must create a new CR regarding to the CRD of this operator. Then, this operator will notice the new or changed CR, and will create or update a ConfigMap for Fluent Bit. Before the ConfigMap is applied, the operator uses the `dry-run` feature of Fluent Bit to validate the new configuration. If the check was successful, the new ConfigMap is applied and Fluent Bit is restarted.
