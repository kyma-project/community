# Comparison matrix

Feature | Kubesphere Operator | Banzai Cloud Operator
--- | --- | ---
License | [Apache 2.0](https://github.com/kubesphere/fluentbit-operator/blob/master/LICENSE)| [Apache 2.0](https://github.com/banzaicloud/logging-operator/blob/master/LICENSE)
Underlying technology | FluentBit | FluentBit for log collection and Fluentd for filtering and sending out to backends
Dynamic configuration | CRDs are directly translatable to FluentBit config sections: `Inputs`, `Filters`, `Outputs`: https://github.com/kubesphere/fluentbit-operator#overview. Individual pieces can be combined into processing pipelines via label matching | The `Logging` resource represents the logging system, and contains configurations for Fluentd and FluentBit. `Output` defines an output for a logging flow. `Flow` defines a logging flow with filters and outputs. There are also cluster-scoped counterparts: `ClusterOutput` and `ClusterFlow`
Validation | CRD schema validation | CRD schema validation and Fluentd configuration checking
Config reloading | Custom FluentBit image with a bootstrapper that starts a child FluentBit process and restarts it if the config changes | Fluentd config reloading sidecar
Passing sensitive info as Secrets | :x: | Secrets can be used in `Output` definitions: https://banzaicloud.com/docs/one-eye/logging-operator/configuration/plugins/outputs/secret/
Debugging | FluentBit logs | FluentBit/Fluentd logs (the latter does not produce any useful logs even with the log level set to `debug`)
Rollback | :x: | :x:
Label matching | |
