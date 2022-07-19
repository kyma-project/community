# Proof of Concepts

## Scrape Scalability
Analyze how the scraping part based on the Prometheus pull model can scale out. See [results](./pocs/scrape-scalability/README.md)

## Processing Scalability and High Availability
Analyze how the main processing part (OTLP input to multiple outputs) can scale out, so that it can handle unlimited load, and the OTLP endpoint is available at any time.

## Pipeline Isolation / Buffer Management
Analyze how filters and outputs can run isolated from each other, so that in case of problems only one pipeline stops working. Clarify which guarantees about the metrics delivery can be given.

## Replace Kubernetes exporters with native otel receivers
Research whether there are native replacements available so that there is no need to have custom components deployed (like the node-exporter) or custom config.

Can we replace node-exporter metrics with [hostmetricsreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/hostmetricsreceiver)?
Can we replace kubelet metrics with [kubeletstatsreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/receiver/kubeletstatsreceiver)?
Can we replace kube-state-metrics with [k8sclusterreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/k8sclusterreceiver)?
