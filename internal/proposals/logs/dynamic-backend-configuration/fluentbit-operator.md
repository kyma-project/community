## Motivation
In [this document](https://github.com/chrkl/kyma-community/blob/2f4ec19bf9e3d8700aadffbd18d2b275888f380a/internal/proposals/logs/dynamic-backend-configuration/overview.md) we compared different solutions for a dynamic logging configuration setup. As a result, we decided to come up with a custom operator written and maintained by the `Huskies`. Since then the operator written by `kubesphere` has been aquired by `fluentbit`. Thus, the project now has more active contributors and got further developed. In this document we want to have a look at the changes/improvements since the last comparision, and want to see if it is worth switching to the (now) official fluentbit operator.

## Fluentbit Operator (former Kubesphere)
Fluent Bit Operator defines the following custom resources:
* `FluentBit`: Defines Fluent Bit instances and its associated config. It requires kubesphere/fluent-bit for dynamic configuration.
* `FluentBitConfig`: Selects input/filter/output plugins and generates the final config into a Secret.
* `Input`: Defines input config sections.
* `Parser`: Defines parser config sections.
* `Filter`: Defines filter config sections.
* `Output`: Defines output config sections.
Each Input, Parser, Filter, Output represents a Fluent Bit config section, which are selected by FluentBitConfig using label selectors. The operator watches those objects, makes the final config data, and creates a Secret to store it, which will be mounted onto Fluent Bit instances owned by Fluent Bit.

Note that the operator works with kubesphere/fluent-bit, a fork of fluent/fluent-bit. Due to [the known issue](https://github.com/fluent/fluent-bit/issues/365), the original Fluent Bit doesn't support dynamic configuration. To address that, kubesphere/fluent-bit incorporates a configuration reloader into the original. See kubesphere/fluent-bit documentation for more information.


# Feature comparison matrix
Feature | Kubesphere Operator | Banzai Cloud Operator
--- | --- | ---
License | [Apache 2.0](https://github.com/kubesphere/fluentbit-operator/blob/master/LICENSE)| [Apache 2.0](https://github.com/banzaicloud/logging-operator/blob/master/LICENSE)
Underlying technology | Fluent Bit | Fluent Bit for log collection and Fluentd for filtering and sending out to backends
Dynamic configuration | CRDs are directly translatable to Fluent Bit config sections in the most straightforward way. | CRDs provide a level of abstraction translatable to Fluent Bit/FluentD configurations (label matching, namespace matching vs cluster scope, secret injection).
Validation | CRD schema validation | CRD schema validation and Fluentd configuration checking
Config reloading | Custom Fluent Bit image with a bootstrapper that starts a child Fluent Bit process and restarts it if the config changes | Fluentd config reloading sidecar
Passing sensitive info as Secrets | Not implemented | Secrets can be used in `Output` definitions: https://banzaicloud.com/docs/one-eye/logging-operator/configuration/plugins/outputs/secret/
Debugging | Fluent Bit logs | Fluent Bit/Fluentd logs (for some reason Fluentd stores logs to a file)
Rollback | Not implemented | A config is applied if the checker run succeeds
Customization | Custom Fluent Bit parser plugins supported. Inputs, filters, and outputs are planned to be supported | No custom Fluentd plugins supported

# Conclusion

Kubesphere Operator is a very basic implementation and lacks some crucial features (secret injection, config check, bad config rollback, etc.).
We can not use it as-is, but it can be a good starting point if we decide to go with a custom logging operator.


