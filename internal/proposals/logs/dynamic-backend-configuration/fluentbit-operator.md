# Motivation
In [this document](https://github.com/chrkl/kyma-community/blob/2f4ec19bf9e3d8700aadffbd18d2b275888f380a/internal/proposals/logs/dynamic-backend-configuration/overview.md) we compared different solutions for a dynamic logging configuration setup. As a result, we decided to come up with a custom operator written and maintained by the `Huskies`. Since then the operator written by `kubesphere` has been aquired by `FluentBit`. Thus, the project now has more active contributors and got further developed. In this document we want to have a look at the changes/improvements since the last comparision, and want to see if it is worth switching to the (now) official fluentbit operator.

# FluentBit Operator (former Kubesphere)
Fluent Bit Operator defines the following custom resources:
* `FluentBit`: Defines Fluent Bit instances and its associated config. It requires kubesphere/fluent-bit for dynamic configuration.
* `FluentBitConfig`: Selects input/filter/output plugins and generates the final config into a Secret.
* `Input`: Defines input config sections.
* `Parser`: Defines parser config sections.
* `Filter`: Defines filter config sections.
* `Output`: Defines output config sections.

Each Input, Parser, Filter, Output represents a Fluent Bit config section, which are selected by FluentBitConfig via label selectors. The operator watches those objects, constructs the final config, and finally creates a Secret to store the config. This secret will be mounted into the Fluent Bit DaemonSet. 
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
| Add [fluent-bit-watcher](https://github.com/fluent/fluentbit-operator/pull/62/files). | Kinda does what the basic-function of kubebuilder operator does (restart pods, cheks if FluentBit is running, etc.), but maybe this works more efficient |

**Enhancements:**

* Use hostpath instead of emptydir to store position db #72
* Improve fluent-bit-watcher synchronization mechanism #74
* Terminate the fluent-bit process in a more elegant way in fluent-bit-watcher #90
* fluent-bit-watcher: goroutine synchronization improvements. #74
* Add support for plugin alias property in input and output specs. #64.

# Feature comparison matrix
Feature | FluentBit Operator | Kyma Telemetry Operator
--- | --- | ---
License | [Apache 2.0](https://github.com/kubesphere/fluentbit-operator/blob/master/LICENSE)| [Apache 2.0](https://github.com/banzaicloud/logging-operator/blob/master/LICENSE)
Underlying technology | Fluent Bit, Operator built with Kubebuilder | Fluent Bit for log collection, Operator built with Kubebuilder
Dynamic configuration | CRDs are directly translatable to Fluent Bit config sections in the most straightforward way. | CRDs are simple structured and support FluentBit syntax, which are then cpoied
Validation | CRD schema validation | Uses FluentBit dryrun feature to validate configuration
Config reloading | Custom Fluent Bit image with a bootstrapper that starts a child Fluent Bit process and restarts it if the config changes | If config changes, configuration will be applied and pods restarted
Passing sensitive info as Secrets | Not implemented | SecretReferences can be passed. SecretReference is a pointer to a Kubernetes secret that should be provided as environment to Fluent Bit.
Rollback | Not implemented | Not implemented
Input, Filter, and Output support | Not all features of inputs, filters, and outputs are supported. But it seems since the acquision they are adding more and more. | All official inputs, filters, and outputs are supported. **But** we do not want that the user is able to create new inputs, creating inputs will be restricted (not implemented until now).
Customization | Custom Fluent Bit parser plugins supported. Inputs, filters, and outputs are planned to be supported | No custom Fluentd plugins supported, user needs to forward the logs using a custom output to another log processor.

# Conclusion

The FluentBit Operator is quite strict with what it supports, which means until now it is not fully capable of supporting all features and filters of FluentBit. This is due to the reasons that the inputs, filters, and outputs are explicit defined in yaml files. Means, when new features to a filter are added, then these features have to be added to the FluentBit Operator as well, which is a downside relating to time.
The Kyma Telemetry Operator is a very simple implementation of a FluentBit Operator, which supports all official configurations of the currently used FluentBit version in Kyma. Thus, we do not have to rely on the FluentBit Operator team to add new features. Same goes for wanted enhancements.
Furthermore, the FluentBit Operator lacks some crucial features (secret injection, config check, bad config rollback, etc.), on which no concrete plans are published.
