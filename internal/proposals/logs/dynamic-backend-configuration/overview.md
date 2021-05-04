# Motivation
Kyma uses Fluent Bit as a log collection solution. Fluent Bit is configured with a plain text file (or multiple files) and does not support dynamic configuration.
It poses the following problems:

1. Static configuration should be provided upon Kyma installation and cannot be changed afterward.
2. It's impossible to separate the Kyma system configuration from the customer one.
3. Every time we change the system configuration in managed Kyma, we have to touch Kyma Environment Broker's code.

# Goal
Come up with a way to dynamically configure Fluent Bit (no manual pod restarting required). It should be possible to apply additional configuration snippets, which are not resettable by the Kyma upgrade process (reconciliation).

# Possible solutions
What immediately comes to mind is Prometheus Operator, which solves very similar problems. 
Prometheus, same as Fluent Bit, is configured with a plain text file (possibly, very long and complex). Configuring Prometheus by hand is not a trivial task.
Prometheus Operator simplifies the configuration by introducing a bunch of custom resources translatable in a Prometheus config.
It also makes it possible to dynamically update different parts of a Prometheus config without restarting any pods.

After some investigation, the following open source solutions have been found:

- Fluent Bit Operator from Kubesphere: https://github.com/kubesphere/fluentbit-operator
- Logging Operator from Banzai Cloud: https://github.com/banzaicloud/logging-operator

## Fluent Bit Operator from Kubesphere
Fluent Bit Operator defines following custom resources:
* `FluentBit`: Defines Fluent Bit instances and its associated config. (It requires to work with kubesphere/fluent-bit for dynamic configuration.)
* `FluentBitConfig`: Select input/filter/output plugins and generates the final config into a Secret.
* `Input`: Defines input config sections.
* `Parser`: Defines parser config sections.
* `Filter`: Defines filter config sections.
* `Output`: Defines output config sections.
Each Input, Parser, Filter, Output represents a Fluent Bit config section, which are selected by FluentBitConfig via label selectors. The operator watches those objects, make the final config data and creates a Secret for store, which will be mounted onto Fluent Bit instances owned by FluentBit.

Note that the operator works with kubesphere/fluent-bit, a fork of fluent/fluent-bit. Due to [the known issue](https://github.com/fluent/fluent-bit/issues/365), the original Fluent Bit doesn't support dynamic configuration. To address that, kubesphere/fluent-bit incorporates a configuration reloader into the original. See kubesphere/fluent-bit documentation for more information.

### Demo

Execute the following commands:

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

Inspect the logs of one of the Fluent Bit pods. It should print out raw logs collected by the container runtime, which look something like:
```bash
[17] kube.var.log.containers.fluent-bit-lcgh6_kubesphere-logging-system_fluent-bit-44cbdf54098e1cbb370215b18dca2a4f8f8660c08858756b4520d2c793c746a3.log: [1620154486.937270488, {"log"=>"{"log":"[3] kube.var.log.containers.ory-hydra-68f975fb7f-whfll_kyma-system_hydra-699cdccd54bf8589fc661d9011eb0f99f6fc27456e09b732efe333c4ba926cca.log: [1620154471.894581175, {\"log\"=\u003e\"{\"log\":\"time=\\\"2021-05-04T18:54:31Z\\\" level=info msg=\\\"completed handling request\\\" measure#hydra/admin: http://127.0.0.1:4444/.latency=160000 method=GET remote=\\\"127.0.0.1:59574\\\" request=/health/alive status=200 text_status=OK took=\\\"160µs\\\"\\n\",\"stream\":\"stderr\",\"time\":\"2021-05-04T18:54:31.563787749Z\"}\"}]\n","stream":"stdout","time":"2021-05-04T18:54:36.840279882Z"}"}]
```

Add a Kubernetes filter by executing the following command:
```bash
kubectl create -f https://raw.githubusercontent.com/skhalash/community/logging-backend-dynamic-configuration/internal/proposals/logs/dynamic-backend-configuration/fluent-bit-operator/kubernetes-filter.yaml
```

Inspect the logs of one of the Fluent Bit pods. Make sure that the logs are augmented with Kubernetes metadata.

## Logging Operator from Banzai Cloud


