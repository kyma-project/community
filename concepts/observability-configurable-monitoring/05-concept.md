# Concept

## Architecture Overview

![b](./assets/arch.drawio.svg)

### Prometheus Pull Support

**Workloads can be scraped via annotations only. A dedicated otel-collector agent with a hard-coded scrape config will take care and emit the metrics to the central otel-collector deployment**

The scraping functionality providing compatibility with the prometheus-style of metrics collection will be realized by a dedicated otel-collector setup. This installation takes care of scraping workloads and pushing the scraped metrics to the central otel-collector via OTLP. Workload can be annotated using specific annotations to enable scraping of the workload. That will be the only way to influence the scrape configuration by the end-user. If the user wants something more custom, he can either use a push approach to the central OTLP endpoint or run an otel-collector sidecar to scrape and push.
To assure that the scraping collector will scale with the targets to scrape, it will be deployed as a daemonset and the targets get splitted on a per node base across the agents. Hereby, the scrape config of the agent just uses the node information to achieve that without a central coordinator.

### OTLP Push Support

**Push support will be provided via OTLP only via the central deployment**

A central collector will be part of the setup which listens for all OTLP metrics. This dedicated central collector runs independent from the pull-based daemonset as the criterias for scaling are different. The central deployment mainly will auto-scale on resource consumption related to the incoming OTLP traffic. The related service port can be used for pushing custom metrics of any kind. Also, that instance will be configurable via the telemetry-operators MetricPipeline resource. In future it might be used for other kind of signals like traces as well.

### K8S & Kyma Integration

**Based on polling via hard scrape config or annotations**

K8S infrastructure will be polled as part of the scrape agent configuration. Potential replacements coming with the ote-collector will be used if possible (like for scraping node metrics). Kyma components will be scraped via annotations like any other custom workload.

### Configurability

Telemetry Operator configures pipelines of central collector, using sane defaults with focus on value for the user. User can easily select/deselect metrics by workload/namespace and so on.

### In-Cluster backend

Monitoring component brings config to push metrics via remotewrite to managed prometheus instances. One for system, istio and custom metrics each. There will be no further tooling like the prometheus-operator ar an alertmanager available.
