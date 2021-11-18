# Motivation

In [this document](./03-fluent-bit-operator1.md), we compared different solutions for a dynamic logging configuration setup. As a result, we decided to come up with a custom operator written and maintained by the Kyma team. Since then, the operator written by `kubesphere` has been contributed to `FluentBit`. Thus, the project now has more active contributors and has developed further. In this document, we want to have a look at the changes/improvements since the last comparison, and want to see if it is worth switching to the (now) official FluentBit operator.

# FluentBit Operator (former Kubesphere)
Fluent Bit Operator defines the following custom resources:
* `FluentBit`: Defines Fluent Bit instances and its associated config. It requires kubesphere/fluent-bit for dynamic configuration.
* `FluentBitConfig`: Selects input/filter/output plugins and generates the final config into a Secret.
* `Input`: Defines input config sections.
* `Parser`: Defines parser config sections.
* `Filter`: Defines filter config sections.
* `Output`: Defines output config sections.

Each Input, Parser, Filter, or Output represents a Fluent Bit config section selected by FluentBitConfig via label selectors. The operator watches those objects, constructs the final config, and finally creates a Secret to store the config. This secret will be mounted into the Fluent Bit DaemonSet. 
In addition, it supports dynamic reloading, which means it updates the configuration without rebooting the Fluent Bit pods.

## Features/Enhancement added since comparison (May 2021)

**Features:**

| New Feature | Comment | Implemented in Kyma Telemetry Operator
|---|---|---|
| Optional watch-namespaces for controller manager. [#117](https://github.com/fluent/fluentbit-operator/pull/117) | Very handy and useful to have.|  Not implemented and not planned. |
| Add support for Containerd and CRI-O regarding tail input plugin.| Nice to have. Not officially supported in tail input plugin.|  Not implemented and not planned. |
| Add support for inotify_watcher configuration of tail input plugin.| Allows to disbale inotify at runtime. inotify helps to keep track of the file changes under the directories.|   Not implemented and not planned. |
| Add support for [throttle filter plugin](https://docs.fluentbit.io/manual/pipeline/filters/throttle).|  | Implemented. |
| Add runtimeClassName support to FluentBit CRD. | (Not needed for our usecase, not quite sure) | Not implemented. Needs to be adapted in FluentBit Chart if wanted. |
| Support setting imagePullSecrets for both [operator](https://github.com/fluent/fluentbit-operator/pull/93/files) and [fluentbit](https://github.com/fluent/fluentbit-operator/pull/94/files) | Nice to have. | Not implemented.
| Add switch to input.tail.memBufLimit in helm chart |  | Implemented. |
| Add [fluent-bit-watcher](https://github.com/fluent/fluentbit-operator/pull/62/files). | Watches for new fluent-bit config changes. Restarts fluent-bit process instead of pod to apply config. This way the fluent-bit pod needn't be restarted to reload the new config. The fluent-bit config is reloaded in this way because there is no reload interface in fluent-bit itself. |

**Enhancements:**

* Use hostpath instead of emptydir to store position db
* Improve fluent-bit-watcher synchronization mechanism
* Terminate the fluent-bit process in a more elegant way in fluent-bit-watcher
* fluent-bit-watcher: goroutine synchronization improvements.
* Add support for plugin alias property in input and output specs.

# Feature comparison matrix
Feature | FluentBit Operator | Kyma Telemetry Operator
--- | --- | ---
License | [Apache 2.0](https://github.com/kubesphere/fluentbit-operator/blob/master/LICENSE)| [Apache 2.0](https://github.com/banzaicloud/logging-operator/blob/master/LICENSE)
Underlying technology | Fluent Bit, Operator built with Kubebuilder | Fluent Bit for log collection, Operator built with Kubebuilder
Dynamic configuration | CRDs are directly translatable to Fluent Bit config sections in the most straightforward way. | CRDs are simple structured and support FluentBit syntax, which are then cpoied
Validation | CRD schema validation (fast and preicse syntax validation, no semantic validation) | Uses FluentBit dryrun feature to validate configuration (more a semantic validation, more hard to translate validation error to the actual config line)
Config reloading | Custom Fluent Bit image with a bootstrapper that starts a child Fluent Bit process and restarts it if the config changes (intransparent to k8s)| If config changes, configuration will be applied and pods restarted (transparent to k8s)
Passing sensitive info as Secrets | Not implemented | SecretReferences can be passed. SecretReference is a pointer to a Kubernetes secret that should be provided as environment to Fluent Bit.
Rollback | Not implemented | Not implemented
Input, Filter, and Output support | Not all features of inputs, filters, and outputs are supported. But it seems since the acquision they are adding more and more. | All official inputs, filters, and outputs are supported. **But** we do not want that the user is able to create new inputs, creating inputs will be restricted (not implemented until now).
Customization | Custom Fluent Bit parser plugins supported. Inputs, filters, and outputs are planned to be supported | No custom Fluentd plugins supported, user needs to forward the logs using a custom output to another log processor.
| Daemonset Management | The Daemonset definition is managed by the operator, incl. potential migration logic | Daemonset will only be restarted but the actual definition is un-managed |

# Chances and Risks of Fluent Bit Operator

## Chances 
* Going with the Fluent-Bit operator will allow Kyma to **save investment** and maintenance. The logging configuration will be changeable at runtime without being forced to touch the fluent-bit system parts (main goal of the story). Furthermore it will maintain the Daemonset and might **reduce maintenance efforts on upgrades**.
* Furthermore, users will use **well known resource definitions** for the configuration with existing documentation (to be proven) and an existing community behind.
* The software will be **battle proven** and mostly likely will have a higher quality than a custom solution.

## Risks 
* On the other hand, using the fluent-bit operator as the direct interface to the user will **prevent Kyma to interfere with the configuration options** of the user. So if the Kyma use case is different or will develop in a different direction, it will not be possible to modify the interface (CRDs). That problem is exposed already by having the Daemonset and fluentbit service section configurable via the CRDs. That must be shielded from the user in order to have the actual operations of the component configurable by the reconciliation logic while the actual configuration should be doable outside of the reconciliation scope. **-> Strategic risk**
* The interface will be specific to the logging use case always (fluent-bit opens up to metrics nowadays, still it does not seem to be as multi-purpose as the otel-collector). As Kyma is planning to provide interface for configuration of tracing and monitoring as well, these configurations will always be different by design starting by the CRD APIgroup already. **-> Strategic risk**
* Another risk is the **indirection on top of fluent-bit** not being in control of the Kyma project. Assuming fluent-bit requires a patch which needs to be activated by some Daemonset config, then the time until that option will be available will get prolonged by additionally waiting to get the operator fixed and released. **-> Investment risk**
* The **strictness of the CRD definitions** might become another challenge, if there is some bug in the translation of the CRD into the actual configuration (a very specific field of a plugin configuration has a wrong validation rule), a fix must be created in the community first and released. Having a free-form plugin definition with a non strict-validation (but a dry-run) might be way more flexible from a dependency perspective. Furthermore all standard features are not supported yet. Means, when new features to a filter are added, then these features have to be added to the FluentBit Operator as well, which is a downside relating to time. **-> Investment risk**
* Furthermore, the FluentBit Operator **lacks important features** (secret injection, config check, bad config rollback, etc.), on which no concrete plans are published. Adding these features will require cooperation with the project and potential investments **-> Strategic risk**
* The operator is too early (no stable release even), it is not obvious if it will find adoption and the project will have an active community long-term **-> Strategic risk**
* An external component will bring maintenance efforts, the image must be hardened and often even duplicated to keep it clean from known vulnerabilities. Furthermore bugs must be fixed/workarounded on an external codebase. **-> Investment risk**
* Having a custom operator to keep the independence in the user-facing interface and additional have the advantage of the fluent-bit operator as internal mechanism will increase the dependencies heavy, and adds another component to maintain from all kind of aspects like for example security scanning.

## Chances custom operator
* The custom operator will be written as part of the Kyma team and will be in full control, especially its interface which can be designed so that it fits 100% into the usage scenarios. Furthermore, it does not have to be specific to the logging scenario, it can be extended to any other observability related usage and has the chance of providing an unified and consistent approach to the users.
* The Kyma Telemetry Operator is a very simple implementation of a FluentBit Operator, which supports all official configurations of the currently used FluentBit version in Kyma, as it relies on a free-form interface without specific syntax validation (**less indirection** which saves maintenance). Thus, Kyma does not have to rely on anyone to have new fluent-bit features usable.

## Risks custom operator
* If the custom operator develops in a similar direction as the fluent-bit operator, then Kyma could have used simply that operator and **could have saved investment**, seeing development and maintenance efforts. Potentially the fluent-bit operator will have a much bigger community and with that feature richness and stability in long-term. **-> Investment risk**
* Users might complain why it is so hard to configure fluent-bit while there is an accepted way out there available. So there must be a good differentiation so that users immediately see the value of the different approach. **-> Publicity/Strategic risk**

# Conclusion
Both operators are quite trivial from the functional aspects and code base (at least for now). Introducing the maintenance effort for a full external component for that relatively small functionality does not seem worthwhile. Furthermore, Kyma will lose the flexibility of changing the user interface as needed and cannot come up with a consistent solution across the other domains (tracing, monitoring).

For now, the decision is to go with the custom operator, but the decision should be re-evaluated periodically so as not to miss the point in time where the investment on our side should be stopped.

