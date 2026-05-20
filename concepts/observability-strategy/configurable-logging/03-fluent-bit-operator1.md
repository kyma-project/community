# Motivation
Kyma is using and wants to continue using Fluent Bit as a log collection solution. Fluent Bit is configured with one or more plain text files and does not support dynamic configuration.
It poses the following problems:

1. Static configuration must be provided upon Kyma installation and cannot be changed afterward.
2. It's impossible to separate the Kyma system configuration from the customer one.
3. Every time we change the system configuration in managed Kyma, we have to touch Kyma Environment Broker's code.

# Goal
Come up with a way to dynamically configure Fluent Bit (no manual pod restarting required). It should be possible to apply additional configuration snippets, which are not resettable by the Kyma upgrade process (reconciliation).

# Possible solutions
What immediately comes to mind is something similar to Prometheus Operator, which solves very similar problems. 
Prometheus, same as Fluent Bit, is configured with a plain text file (possibly, very long and complex). Configuring Prometheus by hand is not a trivial task.
Prometheus Operator simplifies the configuration by introducing a bunch of custom resources translatable in a Prometheus config.
It also makes it possible to dynamically update different parts of a Prometheus config without restarting any pods.

After some investigation, the following open source solutions have been found:

- Fluent Bit Operator from Kubesphere: https://github.com/kubesphere/fluentbit-operator
- Logging Operator from Banzai Cloud: https://github.com/banzaicloud/logging-operator

## Kubesphere Operator
Fluent Bit Operator defines the following custom resources:

* `FluentBit`: Defines Fluent Bit instances and its associated config. It requires kubesphere/fluent-bit for dynamic configuration.
* `FluentBitConfig`: Selects input/filter/output plugins and generates the final config into a Secret.
* `Input`: Defines input config sections.
* `Parser`: Defines parser config sections.
* `Filter`: Defines filter config sections.
* `Output`: Defines output config sections.

Each Input, Parser, Filter, Output represents a Fluent Bit config section, which are selected by FluentBitConfig using label selectors. The operator watches those objects, makes the final config data, and creates a Secret to store it, which will be mounted onto Fluent Bit instances owned by Fluent Bit.

Note that the operator works with kubesphere/fluent-bit, a fork of fluent/fluent-bit. Due to [the known issue](https://github.com/fluent/fluent-bit/issues/365), the original Fluent Bit doesn't support dynamic configuration. To address that, kubesphere/fluent-bit incorporates a configuration reloader into the original. See kubesphere/fluent-bit documentation for more information.

### Demo
1. Run the following commands:

```bash
// operator
kubectl create -f https://raw.githubusercontent.com/kubesphere/fluentbit-operator/master/manifests/setup/namespace-namespace.yaml
kubectl create -f https://raw.githubusercontent.com/kubesphere/fluentbit-operator/master/manifests/setup/fluentbit-operator-crd.yaml
kubectl create -f https://raw.githubusercontent.com/kubesphere/fluentbit-operator/master/manifests/setup/fluentbit-operator-clusterRole.yaml
kubectl create -f https://raw.githubusercontent.com/kubesphere/fluentbit-operator/master/manifests/setup/fluentbit-operator-serviceAccount.yaml
kubectl create -f https://raw.githubusercontent.com/kubesphere/fluentbit-operator/master/manifests/setup/fluentbit-operator-clusterRoleBinding.yaml
kubectl create -f https://raw.githubusercontent.com/kubesphere/fluentbit-operator/master/manifests/setup/fluentbit-operator-deployment.yaml

// fluent bit and fluent bit config
kubectl create -f https://raw.githubusercontent.com/kubesphere/fluentbit-operator/master/manifests/quick-start/fluentbitconfig-fluentBitConfig.yaml
kubectl create -f https://raw.githubusercontent.com/kubesphere/fluentbit-operator/master/manifests/quick-start/fluentbit-fluentBit.yaml

// tail input
kubectl create -f https://raw.githubusercontent.com/skhalash/community/logging-backend-dynamic-configuration/internal/proposals/logs/dynamic-backend-configuration/fluent-bit-operator/tail-input.yaml

// stdout output
kubectl create -f https://raw.githubusercontent.com/skhalash/community/logging-backend-dynamic-configuration/internal/proposals/logs/dynamic-backend-configuration/fluent-bit-operator/stdout-output.yaml
```

2. Inspect the logs of one of the Fluent Bit pods. It should print out raw logs collected by the container runtime, which look something like:

```bash
[17] kube.var.log.containers.fluent-bit-lcgh6_kubesphere-logging-system_fluent-bit-44cbdf54098e1cbb370215b18dca2a4f8f8660c08858756b4520d2c793c746a3.log: [1620154486.937270488, {"log"=>"{"log":"[3] kube.var.log.containers.ory-hydra-68f975fb7f-whfll_kyma-system_hydra-699cdccd54bf8589fc661d9011eb0f99f6fc27456e09b732efe333c4ba926cca.log: [1620154471.894581175, {\"log\"=\u003e\"{\"log\":\"time=\\\"2021-05-04T18:54:31Z\\\" level=info msg=\\\"completed handling request\\\" measure#hydra/admin: http://127.0.0.1:4444/.latency=160000 method=GET remote=\\\"127.0.0.1:59574\\\" request=/health/alive status=200 text_status=OK took=\\\"160µs\\\"\\n\",\"stream\":\"stderr\",\"time\":\"2021-05-04T18:54:31.563787749Z\"}\"}]\n","stream":"stdout","time":"2021-05-04T18:54:36.840279882Z"}"}]
```

3. Add a Kubernetes filter by executing the following command:

```bash
kubectl create -f https://raw.githubusercontent.com/skhalash/community/logging-backend-dynamic-configuration/internal/proposals/logs/dynamic-backend-configuration/fluent-bit-operator/kubernetes-filter.yaml
```

4. Inspect the logs of one of the Fluent Bit pods. Make sure that the logs are augmented with Kubernetes metadata.

## Banzai Cloud Operator
https://kube-logging.dev/docs/

Logging Operator automates the deployment and configuration of a Kubernetes logging pipeline. The operator deploys and configures a Fluent Bit daemonset on every node to collect container and application logs from the node file system. Fluent Bit queries the Kubernetes API and enriches the logs with metadata about the pods, and transfers both the logs and the metadata to Fluentd. Fluentd receives, filters, and transfer logs to multiple outputs. Your logs will always be transferred on authenticated and encrypted channels.

You can define outputs (destinations where you want to send your log messages, for example, Elasticsearch, or and Amazon S3 bucket), and flows that use filters and selectors to route log messages to the appropriate outputs. You can also define cluster-wide outputs and flows, for example, to use a centralized output that namespaced users cannot modify.

You can configure the Logging operator using the following Custom Resource Descriptions:

* `Logging`: Represents a logging system. Includes Fluentd and Fluent Bit configuration. Specifies the controlNamespace. Fluentd and Fluent Bit will be deployed in the controlNamespace
* `Output`: Defines an Output for a logging flow. This is a namespaced resource. See also `ClusterOutput`.
* `Flow`: Defines a logging flow with filters and outputs. You can specify selectors to filter logs by labels. Outputs can be output or `ClusterOutput`. This is a namespaced resource. See also `ClusterFlow`.
* `ClusterOutput`: Defines an output without namespace restriction. Only effective in controlNamespace.
* `ClusterFlow`: Defines a logging flow without namespace restriction.

### Configuration checks
When Fluentd aggregates logs, it shares the configurations of different log flows. To prevent a bad configuration from failing the Fluentd process, a configuration check validates these first. The new settings only go live after a successful check.

### Secret definition

Secrets can be used in Logging Operator Output definitions.

### Demo
https://kube-logging.dev/docs/examples/loki-nginx/

1. Run the following commands:

```bash
helm repo add grafana https://grafana.github.io/helm-charts
helm repo add loki https://grafana.github.io/loki/charts
helm repo add banzaicloud-stable https://kubernetes-charts.banzaicloud.com
helm repo update

kubectl create ns logging
kubectl label namespace logging istio-injection=disabled
helm upgrade --install --namespace logging loki loki/loki
helm upgrade --install --namespace logging grafana grafana/grafana \
 --set "datasources.datasources\\.yaml.apiVersion=1" \
 --set "datasources.datasources\\.yaml.datasources[0].name=Loki" \
 --set "datasources.datasources\\.yaml.datasources[0].type=loki" \
 --set "datasources.datasources\\.yaml.datasources[0].url=http://loki:3100" \
 --set "datasources.datasources\\.yaml.datasources[0].access=proxy"
helm upgrade --install --wait --namespace logging logging-operator banzaicloud-stable/logging-operator \
  --set "createCustomResource=false"
helm upgrade --install --wait --namespace logging logging-demo banzaicloud-stable/logging-demo \
  --set "loki.enabled=True"

```

2. Use the following command to retrieve the password of the Grafana admin user:

```bash
kubectl get secret --namespace logging grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
```

3. Enable port forwarding to the Grafana Service:

```bash
kubectl -n logging port-forward svc/grafana 3000:80
```
<!-- markdown-link-check-disable-next-line -->
4. Open the Grafana Dashboard: http://localhost:3000 and log in. Select Menu > Explore, select Data source > Loki, then select Log labels > namespace > logging. A list of logs should appear.

# Feature comparison matrix
Feature | Kubesphere Operator | Banzai Cloud Operator
--- | --- | ---
License | [Apache 2.0](https://github.com/kubesphere/fluentbit-operator/blob/master/LICENSE)| [Apache 2.0](https://github.com/banzaicloud/logging-operator/blob/master/LICENSE)
Underlying technology | Fluent Bit | Fluent Bit for log collection and Fluentd for filtering and sending out to backends
Dynamic configuration | CRDs are directly translatable to Fluent Bit config sections in the most straightforward way. | CRDs provide a level of abstraction translatable to Fluent Bit/FluentD configurations (label matching, namespace matching vs cluster scope, secret injection).
Validation | CRD schema validation | CRD schema validation and Fluentd configuration checking
Config reloading | Custom Fluent Bit image with a bootstrapper that starts a child Fluent Bit process and restarts it if the config changes | Fluentd config reloading sidecar
Passing sensitive info as Secrets | Not implemented | Secrets can be used in `Output` definitions
Debugging | Fluent Bit logs | Fluent Bit/Fluentd logs (for some reason Fluentd stores logs to a file)
Rollback | Not implemented | A config is applied if the checker run succeeds
Customization | Custom Fluent Bit parser plugins supported. Inputs, filters, and outputs are planned to be supported | No custom Fluentd plugins supported

# Conclusion

Kubesphere Operator is a very basic implementation and lacks some crucial features (secret injection, config check, bad config rollback, etc.).
We can not use it as-is, but it can be a good starting point if we decide to go with a custom logging operator.

Banzai Cloud Operator is feature-rich, but also based heavily on Fluentd. We don't have a lot of experience configuring and operating Fluentd. Furthermore, complicating the logging setup doesn't align well with the idea of making Kyma more simple and lightweight. We could reuse some of the ideas, but would also like to stick with plain Fluent Bit (without Fluentd).

Another solution would be to write a custom operator, which will manage Fluent Bit, but also implement all the required features. Even though there is an implementtion and maintainance cost, it seems to the most optimal solution so far.
