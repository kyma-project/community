# Proof of Concepts

## Scrape Scalability
Analyze how the scraping part on base of the prometheus pull model can scale-out. See [results](./pocs/scrape-scalability/README.md)

## Processing Scalability and High Availability
Analyze how the main processing part (OTLP input to muliple outputs) can scale-out so that it can handle unlimited load and the OTLP endpoint can be available at any time

## Pipeline Isolation / Buffer Management
Analyze how filters and outputs can run isolated by each other, so that on problems on one pipeline stops working. Figure out what guarantees about the delivery can be given.

## Replace k8s exporters with native otel receivers
Try pout if there a already native replacements available so that there is no need to have custom components deployed (like the node-exporter) or custom config is needed.

Can we replace node-exporter metrics with [hostmetricsreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/hostmetricsreceiver)
Can we replace kubelet metrics with [kubeletstatsreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/receiver/kubeletstatsreceiver)
Can we replace kube-state-metrics with [k8sclusterreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/k8sclusterreceiver)
