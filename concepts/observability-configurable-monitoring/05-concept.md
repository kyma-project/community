# Concept

## Architecture Overview

![b](./assets/arch.drawio.svg)

### Prometheus Pull Support

**Workloads can be scraped with annotations only. A dedicated otel-collector agent with a hardcoded scrape config will take care and emit the metrics to the central otel-collector deployment.**

The scraping functionality that provides compatibility with the Prometheus-style of metrics collection will be realized by a dedicated otel-collector setup. This installation takes care of scraping workloads and pushing the scraped metrics to the central otel-collector via OTLP. Workloads can be annotated using specific annotations to enable scraping of the workload, which will be the only way for the end user to influence the scrape configuration. If users want something more custom, they can either use the push approach to the central OTLP endpoint, or run an otel-collector sidecar to scrape and push.
To assure that the scraping collector will scale with the targets to scrape, it will be deployed as a daemonset, and the targets are split on a per-node base across the agents. Hereby, the scrape config of the agent just uses the node information to achieve that without a central coordinator.

### OTLP Push Support

**Push support will be provided via OTLP only via the central deployment**

A central collector, which listens for OTLP metrics, will be part of the setup. This dedicated central collector runs independent from the pull-based daemonset as the criterias for scaling are different. The central deployment mainly will auto-scale on resource consumption related to the incoming OTLP traffic. The related service port can be used for pushing custom metrics of any kind. Also, that instance will be configurable via the telemetry-operators MetricPipeline resource. In future it might be used for other kind of signals like traces as well.

### Kubernetes and Kyma Integration

**Based on polling via hard scrape config or annotations**

Kubernetes infrastructure will be polled as part of the scrape agent configuration. If possible, potential replacements coming with the otel-collector will be used (like for scraping node metrics). Kyma components will be scraped via annotations like any other custom workload.

### Configurability

Telemetry Operator configures pipelines of a central collector, using sane defaults with focus on value for the user. User can easily select and deselect metrics by workload or Namespace, and so on.

### In-Cluster backend

The monitoring component brings config to push metrics using remote write to managed Prometheus instances. One each for system, Istio, and custom metrics. There will be no further tooling like the Prometheus operator or an Alertmanager available.
