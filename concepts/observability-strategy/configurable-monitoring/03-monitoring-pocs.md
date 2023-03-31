# Proof of Concepts

## Scrape Scalability
Analyze how the scraping part based on the Prometheus pull model can scale out. A DaemonSet is a simple but effective way of scaling. However, it is not perfect and can lead to unequal distribution. A better approach will be to use a StatefulSet with sharding based on deterministic hashing of pods instead of just the node assignment, but that is much harder to achieve. See [results](./pocs/scrape-scalability/README.md)

## Processing Scalability and High Availability
Analyze how the main processing part (OTLP input to multiple outputs) can scale out, so that it can handle unlimited load, and the OTLP endpoint is available at any time.

## Pipeline Isolation / Buffer Management
Analyze how filters and outputs can run isolated from each other, so that in case of problems only one pipeline stops working. Clarify which guarantees about the metrics delivery can be given.

## Replace Kubernetes exporters with native otel receivers
Research whether there are native replacements available so that there is no need to have custom components deployed (like the node-exporter) or custom config.

Can we replace node-exporter metrics with [hostmetricsreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/hostmetricsreceiver)?
Can we replace kubelet metrics with [kubeletstatsreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/receiver/kubeletstatsreceiver)?
Can we replace kube-state-metrics with [k8sclusterreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/k8sclusterreceiver)?

Yes, an experiment showed that all basic functionality is covered. However, the receivers are following the Otel semantic conventions and the resulting metric names are not compatible with the prometheus based exporters. With that, existing Grafana Dashboards will not work anymore.

## Shared OpenTelemetry Collector vs. individual collectors per signal type 

Determine the pros and cons of having a single shared collector with different pipelines per signal type versus having an individual collector per signal type. See the [results](./pocs/shared-vs-per-sinal-type-collector/README.md).

